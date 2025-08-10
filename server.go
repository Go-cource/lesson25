package main

import (
	"fmt"
	"net/http"
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
