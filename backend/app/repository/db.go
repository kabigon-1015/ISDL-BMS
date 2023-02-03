package repository

import (
	"database/sql"
	"log"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	// "github.com/hoka-isdl/ISDL-BMS/backend/app/structure"
)

var db *sql.DB

func Opendb() {

	db_name, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/bms_db?")

	if err != nil {
		panic(err.Error())
	}

	db = db_name
}
func CreateTask() {
	fmt.Println(1)
	Opendb()
	defer db.Close()

	insert, err := db.Prepare("INSERT INTO Books(id,title,title_kana,ISBN,author,author_kana,publisher,item_capthion,image_url) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)")

	if err != nil {
		log.Fatal(err.Error())
	}
	// insert.Exec(2,"isbn","岡")
	insert.Exec(2,"実践力を身につけるPythonの教科書","ジッセンリョクヲミニツケルパイソンノキョウカショ",9784839960247,"クジラ飛行机","クジラヒコウヅクエ","マイナビ出版","基本文法から始めてアプリ開発までしっかり解説","https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/0247/9784839960247.jpg?_ex=200x200")
}

func Researchbook(isbn string) string {

	// Opendb()
	// defer db.Close()

	// _ = db.QueryRow("SELECT name FROM User WHERE name = ? AND password = ?", name, password).Scan(&user.Name)

	// if user.Name == "" {
	// 	return false
	// }

	return isbn
}
