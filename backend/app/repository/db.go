package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/hoka-isdl/ISDL-BMS/backend/app/structure"
)

var db *sql.DB

func Opendb() {

	db_name, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/taskapp?")

	if err != nil {
		panic(err.Error())
	}

	db = db_name
}

func VerifybyName(name string, password string) bool {

	var user structure.User

	Opendb()
	defer db.Close()

	_ = db.QueryRow("SELECT name FROM User WHERE name = ? AND password = ?", name, password).Scan(&user.Name)

	if user.Name == "" {
		return false
	}

	return true
}

func Verify(name string, password string) bool {
	var user structure.User

	Opendb()
	defer db.Close()

	_ = db.QueryRow("SELECT name FROM User WHERE name = ? AND password = ?", name, password).Scan(&user.Name)
	if user.Name == "" {
		return false
	}
	return true
}

func ChangePass(name string, newpass string) {

	Opendb()
	defer db.Close()

	_ = db.QueryRow("UPDATE User SET password = ? WHERE name = ?", newpass, name)
}

func GetUserList() []string {

	var user_list []string
	var user structure.User

	Opendb()
	defer db.Close()

	rows, err := db.Query("SELECT name FROM User")

	if err != nil {
		log.Fatal(err.Error())
	}

	for rows.Next() {

		rows.Scan(&user.Name)

		user_list = append(user_list, user.Name)
	}

	return user_list
}

func CreateTask(task_name string, description string, pic string, deadline string) {

	Opendb()
	defer db.Close()

	insert, err := db.Prepare("INSERT INTO Task(task_name, description, pic, deadline, status) VALUES(?, ?, ?, ?, ?)")

	if err != nil {
		log.Fatal(err.Error())
	}

	insert.Exec(task_name, description, pic, deadline, "todo")
}

func GetTasks(user_name string) ([][]string, [][]string, [][]string) {

	var tasks structure.Task
	var todo_task_list [][]string
	var doing_task_list [][]string
	var done_task_list [][]string

	Opendb()
	defer db.Close()

	//todo
	rows_todo, err := db.Query("SELECT task_name, description, deadline FROM Task WHERE pic = ? AND status = ?", user_name, "todo")

	if err != nil {
		log.Fatal(err.Error())
	}

	for rows_todo.Next() {
		rows_todo.Scan(&tasks.Task_name, &tasks.Task_description, &tasks.Task_deadline)

		task_data := []string{tasks.Task_name, tasks.Task_description, tasks.Task_deadline}
		todo_task_list = append(todo_task_list, task_data)
	}

	//doing
	rows_doing, err := db.Query("SELECT task_name, description, deadline FROM Task WHERE pic = ? AND status = ?", user_name, "doing")

	if err != nil {
		log.Fatal(err.Error())
	}

	for rows_doing.Next() {
		rows_doing.Scan(&tasks.Task_name, &tasks.Task_description, &tasks.Task_deadline)

		task_data := []string{tasks.Task_name, tasks.Task_description, tasks.Task_deadline}
		doing_task_list = append(doing_task_list, task_data)
	}

	//done
	rows_done, err := db.Query("SELECT task_name, description, deadline FROM Task WHERE pic = ? AND status = ?", user_name, "done")

	if err != nil {
		log.Fatal(err.Error())
	}

	for rows_done.Next() {
		rows_done.Scan(&tasks.Task_name, &tasks.Task_description, &tasks.Task_deadline)

		task_data := []string{tasks.Task_name, tasks.Task_description, tasks.Task_deadline}
		done_task_list = append(done_task_list, task_data)
	}

	return todo_task_list, doing_task_list, done_task_list
}

func UpdateTaskStatus(user_name string, taskName string, taskDescription string, afterStatus string) {

	Opendb()
	defer db.Close()

	_ = db.QueryRow("UPDATE Task SET status = ? WHERE pic = ? AND task_name = ? AND description = ?", afterStatus, user_name, taskName, taskDescription)
}

func DeleteTask(user_name string, taskName string, taskDescription string) {

	Opendb()
	defer db.Close()

	_ = db.QueryRow("DELETE FROM Task WHERE pic = ? AND task_name = ? AND description = ?", user_name, taskName, taskDescription)
}