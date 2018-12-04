package databases

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var SqlDB *sql.DB

func Database() {
	var err error
	SqlDB, err = sql.Open("mysql","root:@tcp(127.0.0.1:3306)/beego?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}
	if err := SqlDB.Ping(); err!=nil {
		log.Fatalln(err)
	}
}