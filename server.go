package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "modernc.org/sqlite"
)

type Agent struct {
	Id             int
	Name           string
	LastTimeOnline string
}

type Task struct {
	Id               int
	AgentName        string
	TaskCreationTime string
	Text             string
	Result           string
	ResultTime       string
}

func dbConnect() *sql.DB {
	db, err := sql.Open("sqlite", "agents.db")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

}
func tasksHandler(w http.ResponseWriter, r *http.Request) {

}
func createTaskHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", indexHandler)                 //просмотр таблицы агентов, создание нового агента
	http.HandleFunc("/tasks", tasksHandler)            //просмотр таблицы задач, обновление задачи
	http.HandleFunc("/create_task", createTaskHandler) //просмотр формы создания задачи, создание новой задачи
	fmt.Println("server is starting...")
	http.ListenAndServe(":8080", nil)
}
