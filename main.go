package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type task struct {
	ID      int    `json:ID`
	Name    string `json:Name`
	Content string `json:Content`
}

type allTasks []task

var tasks = allTasks{
	{
		ID:      1,
		Name:    "Task One",
		Content: "Some Text",
	},
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome To My API")
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask task
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert Valid Data")
		return
	}

	json.Unmarshal(body, &newTask)

	newTask.ID = len(tasks) + 1

	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newTask)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "Insert Valid ID")
		return
	}

	for _, task := range tasks {
		if task.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "This ID Not Exists")
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "Insert Valid ID")
		return
	}
	var newTask task
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert Valid Data")
		return
	}
	json.Unmarshal(body, &newTask)

	for i, item := range tasks {
		if item.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			newTask.ID = id
			tasks = append(tasks, newTask)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(newTask)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "This ID Not Exists")
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "Insert Valid ID")
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, "Task Deleted")
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "This ID Not Exists")
}

func main() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", index)
	r.HandleFunc("/tasks", getTasks)
	r.HandleFunc("/create", createTask).Methods("POST")
	r.HandleFunc("/task/{id}", getTask)
	r.HandleFunc("/delete/{id}", deleteTask).Methods("DELETE")
	r.HandleFunc("/update/{id}", updateTask).Methods("PUT")

	log.Fatal(http.ListenAndServe(":3000", r))
}
