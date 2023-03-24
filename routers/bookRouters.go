package routers

import (
	"book_Gin/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/books", controllers.GetBooks)
	router.GET("/books/:book_id", controllers.GetBook)
	router.POST("/books", controllers.CreateBook)
	router.PUT("/books/:book_id", controllers.UpdateBook)
	router.DELETE("/books/:book_id", controllers.DeleteBook)

	return router
}
