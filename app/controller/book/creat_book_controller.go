package book

import (
	"com.github.gin-common/common/resp"

	"com.github.gin-common/common/controllers"

	"com.github.gin-common/app/form"
	"com.github.gin-common/app/model"
	"com.github.gin-common/app/service"
	"github.com/gin-gonic/gin"
)

type CreateBookController struct {
	createBookForm *form.CreatBookForm
	bookService    service.BookService
}

func (controller *CreateBookController) Init(createBookForm *form.CreatBookForm, bookService service.BookService) {
	controller.createBookForm = createBookForm
	controller.bookService = bookService
}

//implement
func (controller *CreateBookController) createBook(context *gin.Context) (data *resp.Response, err error) {
	if e := context.ShouldBindJSON(controller.createBookForm); e != nil {
		err = e
		return
	}
	var book *model.Book
	book, err = controller.bookService.CreateBook(&model.Book{
		BookName: controller.createBookForm.BookName,
		Price:    controller.createBookForm.Price,
		UserID:   controller.createBookForm.UserID,
	})

	if err != nil {
		return
	}
	data = controllers.Success(gin.H{
		"user": book,
	})
	return
}

func (controller *CreateBookController) DoRequest(context *gin.Context) (data *resp.Response, err error) {
	return controller.createBook(context)
}

func (controller *CreateBookController) Name() string {
	return "create_book_controller"
}
