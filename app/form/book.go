package form

type CreatBookForm struct {
	BookName string  `binding:"required" json:"book_name"`
	UserID   uint    `binding:"required" json:"user_id"`
	Price    float64 `binding:"required" json:"price"`
}
