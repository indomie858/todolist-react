package main

import (
	"database/request"

	"context"
	"fmt"
    "log"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

// Create a new user in the Firstore database wih the provided name
//
// Example:
// http://localhost:10000/create/user/{name}
// http://localhost:10000/create/user/sabra
func createUser(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: createUser")

    // Read the variables passed
    vars := mux.Vars(r)
	uid := vars["uid"]
	name := vars["name"]

    // Create a new request for the app
    req := newRequest("create", uid)

    // Get the payload params and display them to the terminal
	payload := r.URL.Query()

    fmt.Printf("PAYLOAD PARAMATERS\n")
    for k, v := range payload {
        s := fmt.Sprintf("%v => %v", k, v)
        fmt.Println("%v\n", s)
    }

    // Perform the requested action
    user, err := req.AddUser(name, payload)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }
    respondWithJSON(w, http.StatusOK, user)
}

// Create a new list in the Firstore database with the provided name & params
//
// Example:
// http://localhost:10000/create/{uid}/list/{name}?<params>
// http://localhost:10000/create/a3a1hWUx5geKB8qeR6fbk5LZZGI2/list/test_add_list?lock=false
//
// TODO: Add code to update users lists array after this
func createList(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: createList")

    // Read the variables passed
    vars := mux.Vars(r)
	uid := vars["uid"]
	listname := vars["name"]
    fmt.Printf("list_name: %v\n", listname)

    // Create a new request for the app
    req := newRequest("create", uid)

    // Get the payload params and display them to the terminal
	payload := r.URL.Query()

    fmt.Printf("\nPAYLOAD PARAMATERS\n")
    for k, v := range payload {
        s := fmt.Sprintf("%v => %v", k, v)
        fmt.Printf("\n%v\n", s)
    }

    list, err := req.AddList(listname, payload)
    // Perform the requested action
    if err != nil {
        respondWithError(w, http.StatusBadRequest, err.Error())
    }
    respondWithJSON(w, http.StatusOK, list)
}

// Create a new task in the Firstore database
//
// Example:
// http://localhost:10000/create/{uid}/task/{name}?<params>
// http://localhost:10000/create/a3a1hWUx5geKB8qeR6fbk5LZZGI2/task/test_task_1
//
// TODO: Change to require list id so we can add task to user and list as well.
func createTask(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: createTask")

    // Read the variables passed
    vars := mux.Vars(r)
	uid := vars["uid"]
	//listname := vars["list_name"]
    taskname := vars["name"]
    fmt.Printf("task_name: %v", taskname)

    // Create a new request for the app
    req := newRequest("create", uid)

    // Get the payload params and display them to the terminal
	payload := r.URL.Query()

    fmt.Printf("\nPAYLOAD PARAMATERS\n")
    for k, v := range payload {
        s := fmt.Sprintf("%v => %v", k, v)
        fmt.Printf("\n%v\n", s)
    }

    // Perform the requested action
    task, err := req.AddTask(taskname, payload)
    if err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
    }
    respondWithJSON(w, http.StatusOK, task)
}

// Create a new sub task in the Firstore database with the provided name
//
// Example:
// http://localhost:10000/create/{uid}/subtask/{name}
// 
func createSubtask(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: createSubtask")

    // Read the variables passed
    vars := mux.Vars(r)
	uid := vars["uid"]
	//listname := vars["list_name"]
    taskname := vars["name"]
    fmt.Printf("task_name: %v", taskname)

    // Create a new request for the app
    req := newRequest("create", uid)

    // Get the payload params and display them to the terminal
	payload := r.URL.Query()

    fmt.Printf("\nPAYLOAD PARAMATERS\n")
    for k, v := range payload {
        s := fmt.Sprintf("%v => %v", k, v)
        fmt.Printf("\n%v\n", s)
    }
    payload.Add("sub_task", "true")

    // Perform the requested action
    task, err := req.AddTask(taskname, payload)
    if err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
    }
    respondWithJSON(w, http.StatusOK, task)
}

// Remove a user from the Firstore database, specified by UID
//
// Example :
// http://localhost:10000/destroy/{uid}
// http://localhost:10000/destroy/MIUVfleqSkxAtzwNeW0W
//
func destroyUser(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: destroyUser")

    // Read the variables passed
	vars := mux.Vars(r)
	uid := vars["uid"]

    // Create a new request for the app
    req := newRequest("destroy", uid)

    // Perform the requested action
    if err := req.DestroyUser(); err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
    }
    respondWithJSON(w, http.StatusOK, "user successfully deleted")
}

// Remove a list from the Firstore database, specified by list name
//
// Example :
// http://localhost:10000/destroy/{uid}/list/{name}
// http://localhost:10000/destroy/MIUVfleqSkxAtzwNeW0W/list/first_list
//
// TODO: Add code to delete all tasks and subtasks
func destroyList(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: destroyList")

    // Read the variables passed
	vars := mux.Vars(r)
	uid := vars["uid"]
	name := vars["name"]

    // Create a new request for the app
    req := newRequest("destroy", uid)

    // Perform the requested action
	req.GetListByName(name)

    if err := req.DestroyList(); err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
    }
    respondWithJSON(w, http.StatusOK, "list successfully deleted")
}

// Remove a task from the Firstore database, specified by task name
// AND parent_list id <-- TO DO
//
// Example :
// http://localhost:10000/destroy/{uid}/task/{name}
// http://localhost:10000/destroy/a3a1hWUx5geKB8qeR6fbk5LZZGI2/task/test_task_1
//
// TODO: Add code to delete all sub tasks + to filter by parent id
func destroyTask(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: destroyTask")

    // Read the variables passed
	vars := mux.Vars(r)
	uid := vars["uid"]
	name := vars["name"]

    // Create a new request for the app
    req := newRequest("destroy", uid)

    // Perform the requested action
	req.GetTaskByName(name)
    if err := req.DestroyTask(); err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
    }
    respondWithJSON(w, http.StatusOK, "task successfully deleted")
}

// Get a user from the Firstore database with the specified UID
//
// Example :
// http://localhost:10000/read/{uid}
// http://localhost:10000/read/a3a1hWUx5geKB8qeR6fbk5LZZGI2
//
func getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getUser")

    // Read the variables passed
	vars := mux.Vars(r)
	uid := vars["uid"]

    // Create a new request for the app
    req := newRequest("read", uid)

    // Perform the requested action
    user, err := req.GetUser()
	if err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
    }
    respondWithJSON(w, http.StatusOK, user)
}

// Get a list from the Firstore database with the specified list name
// that has an owner with the provided UID
//
// Example :
// http://localhost:10000/read/{uid}/list/{name}
//
func getList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getList")

    // Read the variables passed
	vars := mux.Vars(r)
	uid := vars["uid"]
	name := vars["name"]

    // Create a new request for the app
    req := newRequest("read", uid)

    // Perform the requested action
    list, err := req.GetListByName(name)
	if err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
    }
    respondWithJSON(w, http.StatusOK, list)
}

// Get ALL lists from the Firstore database with that has an owner with
// the provided UID
//
// Example :
// http://localhost:10000/read/{uid}/lists
// http://localhost:10000/read/a3a1hWUx5geKB8qeR6fbk5LZZGI2/lists
//
func getLists(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: getLists")

    // Read the variables passed
    vars := mux.Vars(r)
    uid := vars["uid"]

    // Create a new request for the app
    req := newRequest("read", uid)

    // Perform the requested action
    lists, err := req.GetLists()
    if err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
    }
    respondWithJSON(w, http.StatusOK, lists)
}

// Get a task from the Firstore database with the specified task name
// that has an owner with the provided UID
// AND has the same parent_id as the one provided <--- TO DO
// in case user names a bunch of tasks the same thing just in diff. lists
//
// Example :
// http://localhost:10000/read/{uid}/task/{name}
// http://localhost:10000/read/a3a1hWUx5geKB8qeR6fbk5LZZGI2/task/task1
//
func getTask(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: getTask")

    // Read the variables passed
    vars := mux.Vars(r)
    uid := vars["uid"]
    name := vars["name"]

    // Create a new request for the app
    req := newRequest("read", uid)

    // Perform the requested action

    // Return the task
    task, err := req.GetTaskByName(name)
    if err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
    }
    respondWithJSON(w, http.StatusOK, task)
}

// Get ALL tasks from the Firestore database with the provided UID
// AND has the same parent_id as the one provided <--- TO DO
//
// Example :
// http://localhost:10000/read/{uid}/tasks/{parent_id}
// http://localhost:10000/read/a3a1hWUx5geKB8qeR6fbk5LZZGI2/tasks/NIcoux7atd3A8Lv7guUO
//
func getTasks(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: getTasks")

    // Read the variables passed
    vars := mux.Vars(r)
    uid := vars["uid"]
    parent := vars["parent_id"]

    // Create a new request for the app
    req := newRequest("read", uid)

    // Perform the requested action
    tasks, err := req.GetTasks(parent)
    if err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
    }
    respondWithJSON(w, http.StatusOK, tasks)
}

// Update a Firestore user data
//
// Example :
// http://localhost:10000/update/{uid}?<params>
// http://localhost:10000/update/MIUVfleqSkxAtzwNeW0W?lists=qqEkD06oFudIRrCVPAc5
//
func updateUser(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: updateUser")

    // Read the variables passed
    vars := mux.Vars(r)
    uid := vars["uid"]

    // Get the payload params and display them to the terminal
    payload := r.URL.Query()

    fmt.Printf("\nPAYLOAD PARAMATERS\n")
    for k, v := range payload {
        s := fmt.Sprintf("%v => %v", k, v)
        fmt.Printf("%v\n", s)
    }

    // Create a new request for the app
    req := newRequest("update", uid)

    // Perform the requested action
    req.GetUser()
    user, err := req.UpdateUser(payload)
    if err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
    }
    respondWithJSON(w, http.StatusOK, user)
}

// Update a Firestore list data
//
// Example :
// http://localhost:10000/update/{uid}/list/{list}?<params>
// http://localhost:10000/update/a3a1hWUx5geKB8qeR6fbk5LZZGI2/list/list1?list_name=list1updated&lock=false
//
func updateList(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: updateList")

    // Read the variables passed
    vars := mux.Vars(r)
    uid := vars["uid"]
    listname := vars["list"]
    fmt.Printf("listname: %v\n", listname)

    // Get the payload params and display them to the terminal
    payload := r.URL.Query()

    fmt.Printf("\nPAYLOAD PARAMATERS\n")
    for k, v := range payload {
        s := fmt.Sprintf("%v => %v", k, v)
        fmt.Printf("%v\n", s)
    }

    // Create a new request for the app
    req := newRequest("update", uid)

    // Perform the requested action
	req.GetListByName(listname)
    list, err := req.UpdateList(payload)
    if err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
    }
    respondWithJSON(w, http.StatusOK, list)
}

// Update a Firestore task data
//
// Example :
// http://localhost:10000/update/{uid}/task/{task}?<params>
//
func updateTask(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: updateTask")

    // Read the variables passed
    vars := mux.Vars(r)
    uid := vars["uid"]
    taskname := vars["task"]
    fmt.Printf("taskname: %v\n", taskname)

    // Get the payload params and display them to the terminal
    payload := r.URL.Query()

    fmt.Printf("\nPAYLOAD PARAMATERS\n")
    for k, v := range payload {
        s := fmt.Sprintf("%v => %v", k, v)
        fmt.Printf("%v\n", s)
    }

    // Create a new request for the app
    req := newRequest("update", uid)

    // Perform the requested action
	req.GetTaskByName(taskname)
    task, err := req.UpdateTask(payload)
    if err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
    }
    respondWithJSON(w, http.StatusOK, task)
}

// src: https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.MarshalIndent(payload,  "", "    ")

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
    respondWithJSON(w, code, map[string]string{"error": message})
}

func handleRequests() {
    router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)

	router.HandleFunc("/create/user/{name}", createUser).Methods("GET", "POST")
	router.HandleFunc("/create/{uid}/list/{name}", createList).Methods("GET", "POST")
	router.HandleFunc("/create/{uid}/task/{name}", createTask).Methods("GET", "POST")
	router.HandleFunc("/create/{uid}/subtask/{name}", createSubtask).Methods("GET", "POST")

	router.HandleFunc("/destroy/{uid}", destroyUser).Methods("DELETE")
	router.HandleFunc("/destroy/{uid}/list/{name}", destroyList).Methods("DELETE")
	router.HandleFunc("/destroy/{uid}/task/{name}", destroyTask).Methods("DELETE")

	router.HandleFunc("/read/{uid}", getUser).Methods("GET")

    router.HandleFunc("/read/{uid}/list/{name}", getList).Methods("GET")
    router.HandleFunc("/read/{uid}/lists", getLists).Methods("GET")

    router.HandleFunc("/read/{uid}/task/{name}", getTask).Methods("GET")
    router.HandleFunc("/read/{uid}/tasks/{parent_id}", getTasks).Methods("GET")

	router.HandleFunc("/update/{uid}", updateUser).Methods("PUT")
	router.HandleFunc("/update/{uid}/list/{list}", updateList).Methods("PUT")
	router.HandleFunc("/update/{uid}/task/{task}", updateTask).Methods("PUT")

	log.Fatal(http.ListenAndServe(":10000", router))
}

func newRequest(rtype, uid string) *request.Request {
    var req request.Request

    req.Type = rtype
	req.UserId = uid
	req.Ctx = context.Background()
	req.GetClient()

    return &req
}

func main() {
	handleRequests()
}
