package service

type TokenString string

type AuthService interface {
	// 登录
	Login(username string, password string) (token TokenString, err error)
	// 登出
	Logout(sessionID string) error
}
