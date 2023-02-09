package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hoka-isdl/ISDL-BMS/backend/app/structure"
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

	insert, err := db.Prepare("INSERT INTO Books(id,title,title_kana,ISBN,author,author_kana,publisher,item_caption,image_url) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)")

	if err != nil {
		log.Fatal(err.Error())
	}
	// insert.Exec(2,"isbn","岡")
	insert.Exec(2, "実践力を身につけるPythonの教科書", "ジッセンリョクヲミニツケルパイソンノキョウカショ", 9784839960247, "クジラ飛行机", "クジラヒコウヅクエ", "マイナビ出版", "基本文法から始めてアプリ開発までしっかり解説", "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/0247/9784839960247.jpg?_ex=200x200")
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

func GetBookinfo(isbn string) (string, string, string) {
	var book structure.Books

	Opendb()
	defer db.Close()

	rows_title, err := db.Query("SELECT title, author, publisher FROM Books WHERE isbn = ?", isbn)

	if err != nil {
		log.Fatal(err.Error())
	}

	for rows_title.Next() {
		rows_title.Scan(&book.Book_title, &book.Book_author, &book.Book_publisher)
	}

	return book.Book_title, book.Book_author, book.Book_publisher
}

func Returnbook(id string, isbn string) {
	var book_id string
	var search_result string

	Opendb()
	defer db.Close()

	rows_bookid, err := db.Query("SELECT id FROM Books WHERE isbn = ?", isbn)

	if err != nil {
		log.Fatal(err.Error())
	}

	for rows_bookid.Next() {
		rows_bookid.Scan(&book_id)
	}
	fmt.Print(book_id)

	rows_histid, err := db.Query("SELECT id FROM Rent_hist WHERE bookid = ? AND returned=?", book_id, "1")

	if err != nil {
		log.Fatal(err.Error())
	}

	for rows_histid.Next() {
		rows_histid.Scan(&search_result)
	}
	fmt.Print(search_result)

	if rows_histid != nil {
		_ = db.QueryRow("UPDATE Rent_hist SET returned = ? WHERE id = ? AND renterid = ?", "0", search_result, id)
	}
}

func GetAllBookData() [][]string {
	var all_book_data [][]string
	var book structure.Books

	Opendb()
	defer db.Close()

	rows_all, err := db.Query("SELECT title, author, publisher, id from Books")

	if err != nil {
		log.Fatal(err.Error())
	}
	for rows_all.Next() {
		rows_all.Scan(&book.Book_title, &book.Book_author, &book.Book_publisher, &book.Book_id)

		book_data := []string{book.Book_title, book.Book_author, book.Book_publisher, book.Book_id}
		all_book_data = append(all_book_data, book_data)
	}

	return all_book_data
}

func GetRenterInfo(book_id string) string {
	var hist structure.Rent_hist
	var name string

	Opendb()
	defer db.Close()

	rows_renterid, err := db.Query("SELECT renterid FROM Rent_hist WHERE bookid = ? AND returned=?", book_id, "1")

	if err != nil {
		log.Fatal(err.Error())
	}

	for rows_renterid.Next() {
		rows_renterid.Scan(&hist.Hist_renterid)
	}

	name = "null"

	rows_name, err := db.Query("SELECT name FROM Students where id=?", hist.Hist_renterid)

	if err != nil {
		log.Fatal(err.Error())
	}

	for rows_name.Next() {
		rows_name.Scan(&name)
	}

	return name
}
