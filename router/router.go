package router

import (
	"gin-restful-mysql/apis"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", apis.IndexApi)
	router.GET("/person", apis.AddPersonApi)
	router.GET("/persons", apis.SelectAllPerson)
	router.POST("/addBook25", apis.Add_25_history_book)
	router.POST("/addbookdetail", apis.Add_25_history_book_detail)

	return router
}
