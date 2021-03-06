package bookctx

type Repository interface {
	Create(book *Book)
	FindByISBN(isbn string)
	Delete(book *Book)
}
