package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	BookID string `json:"book_id"`
	Title  string `json:"tittle"`
	Author string `json:"author"`
	Decs   string `json:"decs"`
}

var BookDatas = []Book{}

func CreateBook(ctx *gin.Context) {
	var newBook Book

	err := ctx.ShouldBindJSON(&newBook)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newBook.BookID = fmt.Sprintf("c%d", len(BookDatas)+1)
	BookDatas = append(BookDatas, newBook)

	ctx.JSON(http.StatusCreated, gin.H{
		"book": newBook,
	})
}

func GetBooks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, BookDatas)
}

func GetBook(ctx *gin.Context) {
	bookID := ctx.Param("book_id")
	condition := false

	var result Book
	for i, v := range BookDatas {
		if bookID == v.BookID {
			condition = true
			result = BookDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Data not found",
			"error massage": fmt.Sprintf("book with id %v", bookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": result,
	})
}

func UpdateBook(ctx *gin.Context) {

	bookID := ctx.Param("book_id")
	condition := false
	var updateBook Book

	if err := ctx.ShouldBindJSON(&updateBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	for i, book := range BookDatas {
		if bookID == book.BookID {
			condition = true
			BookDatas[i] = updateBook
			ctx.JSON(http.StatusOK, updateBook)
			return
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "data not found",
			"error_message": fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book with id %v has been succesfully update", bookID),
	})
}

func DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("book_id")
	condition := false
	var bookIndex int

	for i, book := range BookDatas {
		if bookID == book.BookID {
			condition = true
			bookIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("car with id %v not found", bookID),
		})
	}

	copy(BookDatas[bookIndex:], BookDatas[bookIndex:])
	BookDatas[len(BookDatas)-1] = Book{}
	BookDatas = BookDatas[:len(BookDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"meesage": fmt.Sprintf("book with id %v has been successfully deleted", bookID),
	})
}
