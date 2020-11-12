package service

import (
	"com.github.gin-common/app/model"
)

type BookService interface {
	// 创建用户
	CreateBook(book *model.Book) (*model.Book, error)

}
