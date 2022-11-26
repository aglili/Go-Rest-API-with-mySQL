package models

import (
	"github.com/jinzhu/gorm"
	"github.com/cecilcodespython/bookstore-crud-api/pkg/config"
)

var db *gorm.DB

type Book struct{
	gorm.Model
	Title string `gorm:""json:"title"`
	Author string `json:"author"`
	Publisher string `json:"publisher"`
	Isbn string `json:"isbn"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
} 

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}
	
func GetBookbyId(Id int64)  (*Book,*gorm.DB){
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) Book{
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book

}