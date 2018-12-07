package model

import (
	db "gin-restful-mysql/databases"
	"log"
)

type Person struct {
	Id        int    `json:"id" form:"id"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
}

type Book_history_25 struct {
	Id          int    `json:"id" form:"id"`
	BookName    string `json:"book_name" form:"book_name"`
	BookIndex   string `json:"book_index" form:"book_index"`
	BookContent string `json:"book_content" form:"book_content"`
	BookKey    string `json:"book_key" form:"book_key"`
}

type Book_history_25_detail struct {
	Id         int    `json:"id" form:"id"`
	BookName   string `json:"book_name" form:"book_name"`
	BookAuthor string `json:"book_author" form:"book_author"`
	BookDesc   string `json:"book_desc" form:"book_desc"`
	BookKey    string `json:"book_key" form:"book_key"`
}

func (b *Book_history_25_detail) AddHistoryBook25Detail() (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO history_book_25_detail (book_name, book_author, book_desc, book_key) VALUES (?, ?, ?, ?)", b.BookName, b.BookAuthor, b.BookDesc, b.BookKey)
	if err != nil {
		log.Fatalln(err)
		return
	}
	id, err = rs.LastInsertId()
	return
}

func (b *Book_history_25) AddBookHistory25() (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO history_book_25 (book_name, book_index, book_content, book_key) VALUES (?, ?, ?, ?)", b.BookName, b.BookIndex, b.BookContent, b.BookKey)
	if err != nil {
		log.Fatalln(err)
		return
	}
	id, err = rs.LastInsertId()
	return
}

func (p *Person) AddPerson() (id int64, err error) {

	rs, err := db.SqlDB.Exec("INSERT INTO person(first_name, last_name) VALUES (?, ?)", p.FirstName, p.LastName)
	if err != nil {
		log.Fatalln(err)
		return
	}
	id, err = rs.LastInsertId()
	return
}

func (p *Person) GetPersons() (persons []Person, err error) {

	persons = make([]Person, 0)
	rows, err := db.SqlDB.Query("SELECT id, first_name, last_name FROM person")
	defer rows.Close()
	if err != nil {
		return
	}

	for rows.Next() {
		var person Person
		rows.Scan(&person.Id, &person.FirstName, &person.LastName)
		persons = append(persons, person)
	}

	if err = rows.Err(); err != nil {
		return
	}
	return
}

func (b *Book_history_25_detail) GetBookList() (books []Book_history_25_detail, err error) {
	books = make([]Book_history_25_detail, 0)
	rows, err := db.SqlDB.Query("SELECT id, book_name, book_author, book_desc, book_key FROM history_book_25_detail")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var book Book_history_25_detail
		rows.Scan(&book.Id, &book.BookName, &book.BookAuthor, &book.BookDesc, &book.BookKey)
		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		return
	}
	return
}

// 查询书籍章回目录
func (b *Book_history_25) GetBookIndex() (books []Book_history_25, err error) {
	books = make([]Book_history_25, 0)
	rows, err := db.SqlDB.Query("SELECT * FROM history_book_25 WHERE book_key=? and book_key=?", b.BookKey, b.BookKey)
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var book Book_history_25
		rows.Scan(&book.Id, &book.BookName, &book.BookIndex, &book.BookContent, &book.BookKey)
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}
