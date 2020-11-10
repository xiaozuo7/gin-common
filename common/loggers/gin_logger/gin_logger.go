package gin_logger

import (
	"fmt"
	"io"
	"time"

	"go.uber.org/zap/zapcore"

	"com.github.gin-common/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var Log *zap.Logger

func Logger() gin.HandlerFunc {
	return LoggerWithConfig(LoggerConfig{})
}

func LoggerWithFormatter(f gin.LogFormatter) gin.HandlerFunc {
	return LoggerWithConfig(LoggerConfig{
		LoggerConfig: gin.LoggerConfig{
			Formatter: f,
		},
	})
}

type LoggerConfig struct {
	gin.LoggerConfig
	LogLevel zapcore.Level
	Options  []zap.Option
}

func LoggerWithWriter(out io.Writer, logLevel zapcore.Level, options []zap.Option, notLogged ...string) gin.HandlerFunc {
	return LoggerWithConfig(LoggerConfig{
		LogLevel: logLevel,
		LoggerConfig: gin.LoggerConfig{
			Output:    out,
			SkipPaths: notLogged,
		},
		Options: options,
	})
}

func LoggerWithConfig(conf LoggerConfig) gin.HandlerFunc {
	formatter := conf.Formatter
	if formatter == nil {
		formatter = defaultLogFormatter
	}

	out := conf.Output
	if out == nil {
		out = gin.DefaultWriter
	}

	notLogged := conf.SkipPaths

	var skip map[string]struct{}

	if length := len(notLogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range notLogged {
			skip[path] = struct{}{}
		}
	}
	Log = util.GetLogger(conf.LogLevel, out, conf.Options...)

	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		c.Next()

		if _, ok := skip[path]; !ok {
			param := gin.LogFormatterParams{
				Request: c.Request,
				Keys:    c.Keys,
			}
			// Stop timer
			param.TimeStamp = time.Now()
			param.Latency = param.TimeStamp.Sub(start)

			param.ClientIP = c.ClientIP()
			param.Method = c.Request.Method
			param.StatusCode = c.Writer.Status()
			param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()

			param.BodySize = c.Writer.Size()

			if raw != "" {
				path = path + "?" + raw
			}

			param.Path = path
			Log.Info(formatter(param), zap.String("errMsg", param.ErrorMessage), zap.String("logType", "request"))
		}
	}
}

func defaultLogFormatter(param gin.LogFormatterParams) string {
	return fmt.Sprintf("[Request]:[%s] %s %s %s %d",
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
	)
}
