package apis

import (
	"fmt"
	"gin-restful-mysql/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 测试服务器连接情况
func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "it works")
}

// 增加用户api
func AddPersonApi(c *gin.Context) {
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")

	p := model.Person{FirstName: firstName, LastName: lastName}

	ra, err := (&p).AddPerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("insert successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

// 查询所有用户api
func SelectAllPerson(c *gin.Context) {
	p := model.Person{}
	ra, err := (&p).GetPersons()
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":     "successful",
		"persons": ra,
	})
}

// 增加书内容
func Add_25_history_book(c *gin.Context) {
	bookName := c.Request.FormValue("book_name")
	bookIndex := c.Request.FormValue("book_index")
	bookContent := c.Request.FormValue("book_content")
	bookKey := c.Request.FormValue("book_key")
	p := model.Book_history_25{BookName: bookName, BookIndex: bookIndex, BookContent: bookContent, BookKey:bookKey}
	_, err := (&p).AddBookHistory25()
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "add book successful",
	})
}

// 增加书籍描述信息
func Add_25_history_book_detail(c *gin.Context) {
	bookName := c.Request.FormValue("book_name")
	bookAuthor := c.Request.FormValue("book_author")
	bookDesc := c.Request.FormValue("book_desc")
	bookKey := c.Request.FormValue("book_key")
	p := model.Book_history_25_detail{BookName: bookName, BookAuthor: bookAuthor, BookDesc: bookDesc, BookKey:bookKey}
	_, err := (&p).AddHistoryBook25Detail()
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "add book successful",
	})
}

// 获取书籍
func Select_book_desc(c *gin.Context) {
	b := model.Book_history_25_detail{}
	ra, err := (&b).GetBookList()
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":    "successful",
		"result": ra,
	})
}

// 获取书籍内容
func Select_book_Contents(c *gin.Context) {
	bookKey := c.Query("book_key")
	bookName := c.Query("book_name")
	b := model.Book_history_25{BookKey: bookKey, BookName:bookName}
	ra, err := (&b).GetBookIndex()
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":    "successful",
		"result": ra,
	})
}
