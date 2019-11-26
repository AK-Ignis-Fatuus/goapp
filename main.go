package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "akram1911"
	DB_NAME     = "postgres"
        DB_HOST     = "184.172.229.212"
        DB_PORT     = "31611"
)

type Task struct {
	TaskID   string `json:"taskID"`
	TaskName string `json:"taskName"`
        ToDo     string `json:"Todo"`
}

type JsonResponse struct {
	Type    string `json:"type"`
	Data    []Task `json:"data"`
	Message string `json:"message"`
}

func main() {
	router := mux.NewRouter()

	// Get all tasks
	router.HandleFunc("/tasks/", GetTasks).Methods("GET")

	// Create a task
	router.HandleFunc("/tasks/create", CreateTask).Methods("POST")

	// Delete a task by id
	router.HandleFunc("/tasks/{taskid}", DeleteTask).Methods("DELETE")

	// Delete all tasks
	router.HandleFunc("/tasks/deleteall", DeleteTasks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

// Get all tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	printMessage("Getting tasks...")

	// Get all tasks from tasks table that don't have taskID = "1"
	rows, err := db.Query("SELECT * FROM tasks where taskID <> $1", "1")

	checkErr(err)
	var tasks []Task
	// var response []JsonResponse
	// Foreach task
	for rows.Next() {
		var id int
		var taskID string
		var taskName string
                var Todo string

		err = rows.Scan(&id, &taskID, &taskName, &Todo)

		checkErr(err)

		tasks = append(tasks, Task{TaskID: taskID, TaskName: taskName, ToDo: Todo})
	}

	var response = JsonResponse{Type: "success", Data: tasks}

	json.NewEncoder(w).Encode(response)
}

// Create task
func CreateTask(w http.ResponseWriter, r *http.Request) {
	taskID := r.FormValue("taskID")
	taskName := r.FormValue("taskName")
	Todo := r.FormValue("Todo")

	var response = JsonResponse{}

	if taskID == "" || taskName == "" {
		response = JsonResponse{Type: "error", Message: "You are missing taskID, taskName or Todo parameter."}
	} else {
		db := setupDB()

		printMessage("Inserting task into DB")

		fmt.Println("Inserting new task with ID: " + taskID + " name: " + taskName + " and todo " + Todo)

		//var lastInsertID int
		//err := db.QueryRow("INSERT INTO tasks(taskID, taskName) VALUES($1, $2) returning id;", taskID, taskName).Scan(&lastInsertID)
                _, err := db.Exec("INSERT INTO tasks (taskid, taskname, todo) VALUES($1, $2, $3)", taskID, taskName, Todo)
		checkErr(err)

		response = JsonResponse{Type: "success", Message: "Task has been inserted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

// Delete task
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	taskID := params["taskid"]

	var response = JsonResponse{}

	if taskID == "" {
		response = JsonResponse{Type: "error", Message: "You are missing taskID parameter."}
	} else {
		db := setupDB()

		printMessage("Deleting task from DB")

		_, err := db.Exec("DELETE FROM tasks where taskID = $1", taskID)
		checkErr(err)

		response = JsonResponse{Type: "success", Message: "Task has been deleted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

// Delete all tasks
func DeleteTasks(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	printMessage("Deleting all tasks...")

	_, err := db.Exec("DELETE FROM tasks")
	checkErr(err)

	printMessage("All tasks have been deleted successfully!")

	var response = JsonResponse{Type: "success", Message: "All tasks have been deleted successfully!"}

	json.NewEncoder(w).Encode(response)
}

// DB set up
func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	checkErr(err)

	return db
}

// Function for handling messages
func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

// Function for handling errors
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
