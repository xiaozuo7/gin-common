package impl

import (
	"com.github.gin-common/app/exception"
	"com.github.gin-common/app/model"
	"com.github.gin-common/common/exceptions"
	"gorm.io/gorm"
)

type BookServiceImpl struct {
	session *gorm.DB
}

func (service *BookServiceImpl) Init(session *gorm.DB) {
	service.session = session
}

func (service *BookServiceImpl) CreateBook(book *model.Book) (*model.Book, error) {
	// 创建书籍
	result := service.session.Create(book)
	if result.Error != nil {
		return nil, exceptions.GetDefinedErrors(exception.BookCreateFailed)
	}

	return book, nil

}
