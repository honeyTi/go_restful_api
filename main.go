package main

import (
	db "gin-restful-mysql/databases"
	router2 "gin-restful-mysql/router"
)

func main() {
	db.Database()
	defer db.SqlDB.Close()
	router := router2.InitRouter()
	router.Run(":8090")
}
