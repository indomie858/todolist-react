package main

import (
	"database/request"

	"fmt"
    "log"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
)

type App struct {
    Router *mux.Router
}

// src: https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql
func (a *App) Initialize() {
    a.Router = mux.NewRouter().StrictSlash(true)
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
    log.Fatal(http.ListenAndServe(":10000", a.Router))
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.MarshalIndent(payload,  "", "    ")

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
    respondWithJSON(w, code, map[string]string{"error": message})
}

func (a *App) homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the HomePage!")
	//fmt.Printf("Endpoint Hit: homePage")
}

// Create a new user in the Firstore database wih the provided name
//
// Example:
// http://localhost:10000/create/user/{name}
// http://localhost:10000/create/user/sabra
func (a *App) createUser(w http.ResponseWriter, r *http.Request) {
    //fmt.Println("Endpoint Hit: createUser")

    // Read the variables passed
    vars := mux.Vars(r)
	uid := vars["uid"]
	name := vars["name"]

    // Create a new request for the app
    req := request.NewRequest("create", uid)

    // Get the payload params and display them to the terminal
	payload := r.URL.Query()

    /*fmt.Println("\nPAYLOAD PARAMATERS")
    for k, v := range payload {
        s := fmt.Sprintf("%v => %v", k, v)
        fmt.Printf("%v\n", s)
    }*/

    // Perform the requested action
    user, err := req.AddUser(name, payload)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }
    respondWithJSON(w, http.StatusOK, map[string]*request.UserJSON{"result": user})
}

// Create a new list in the Firstore database with the provided name & params
//
// Example:
// http://localhost:10000/create/{uid}/list/{name}?<params>
// http://localhost:10000/create/a3a1hWUx5geKB8qeR6fbk5LZZGI2/list/test_add_list?lock=false
//
// TODO: Add code to update users lists array after this
func (a *App) createList(w http.ResponseWriter, r *http.Request) {
    //fmt.Println("Endpoint Hit: createList")

    // Read the variables passed
    vars := mux.Vars(r)
	uid := vars["uid"]
	listname := vars["name"]
    //fmt.Printf("list_name: %v\n", listname)

    // Create a new request for the app
    req := request.NewRequest("create", uid)

    // Get the payload params and display them to the terminal
	payload := r.URL.Query()

    /*fmt.Println("\nPAYLOAD PARAMATERS")
    for k, v := range payload {
        s := fmt.Sprintf("%v => %v", k, v)
        fmt.Printf("\n%v\n", s)
    }*/

    list, err := req.AddList(listname, payload)
    // Perform the requested action
    if err != nil {
        respondWithError(w, http.StatusBadRequest, err.Error())
        return
    }
    respondWithJSON(w, http.StatusOK, map[string]*request.ListJSON{"result": list})
}

// Create a new task in the Firstore database
//
// Example:
// http://localhost:10000/create/{uid}/task/{name}?<params>
// http://localhost:10000/create/a3a1hWUx5geKB8qeR6fbk5LZZGI2/task/test_task_1
//
// TODO: Change to require list id so we can add task to user and list as well.
func (a *App) createTask(w http.ResponseWriter, r *http.Request) {
    //fmt.Println("Endpoint Hit: createTask")

    // Read the variables passed
    vars := mux.Vars(r)
	uid := vars["uid"]
    taskname := vars["name"]
    //fmt.Printf("task_name: %v", taskname)

    // Create a new request for the app
    req := request.NewRequest("create", uid)

    // Get the payload params and display them to the terminal
	payload := r.URL.Query()

    /*fmt.Println("\nPAYLOAD PARAMATERS")
    for k, v := range payload {
        s := fmt.Sprintf("%v => %v", k, v)
        fmt.Printf("\n%v\n", s)
    }*/

    // Perform the requested action
    task, err := req.AddTask(taskname, payload)
    if err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
        return
    }
    respondWithJSON(w, http.StatusOK, map[string]*request.TaskJSON{"result": task})
}

// Create a new sub task in the Firstore database with the provided name
//
// Example:
// http://localhost:10000/create/{uid}/subtask/{name}?<params>
//
func (a *App) createSubtask(w http.ResponseWriter, r *http.Request) {
    //fmt.Println("Endpoint Hit: createSubtask")

    // Read the variables passed
    vars := mux.Vars(r)
	uid := vars["uid"]
    taskname := vars["name"]
    //fmt.Printf("task_name: %v", taskname)

    // Create a new request for the app
    req := request.NewRequest("create", uid)

    // Get the payload params and display them to the terminal
	payload := r.URL.Query()

    /*fmt.Println("\nPAYLOAD PARAMATERS")
    for k, v := range payload {
        s := fmt.Sprintf("%v => %v", k, v)
        fmt.Printf("\n%v\n", s)
    }*/
    payload.Add("sub_task", "true")

    // Perform the requested action
    task, err := req.AddTask(taskname, payload)
    if err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
        return
    }
    respondWithJSON(w, http.StatusOK, map[string]*request.TaskJSON{"result": task})
}

// Remove a user from the Firstore database, specified by UID
//
// Example :
// http://localhost:10000/destroy/{uid}
// http://localhost:10000/destroy/MIUVfleqSkxAtzwNeW0W
//
func (a *App) destroyUser(w http.ResponseWriter, r *http.Request) {
    //fmt.Println("Endpoint Hit: destroyUser")

    // Read the variables passed
	vars := mux.Vars(r)
	uid := vars["uid"]

    // Create a new request for the app
    req := request.NewRequest("destroy", uid)

    // Perform the requested action
    if err := req.DestroyUser(); err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
        return
    }
    respondWithJSON(w, http.StatusOK, map[string]string{"result": "user successfully deleted"})
}

// Remove a list from the Firstore database, specified by list name
//
// Example :
// http://localhost:10000/destroy/{uid}/list/{name}
// http://localhost:10000/destroy/MIUVfleqSkxAtzwNeW0W/list/first_list
//
// TODO: Add code to delete all tasks and subtasks
func (a *App) destroyList(w http.ResponseWriter, r *http.Request) {
    //fmt.Println("Endpoint Hit: destroyList")

    // Read the variables passed
	vars := mux.Vars(r)
	uid := vars["uid"]
	name := vars["name"]

    // Create a new request for the app
    req := request.NewRequest("destroy", uid)

    // Perform the requested action
    if err := req.DestroyList(name); err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
        return
    }
    respondWithJSON(w, http.StatusOK, map[string]string{"result": "list successfully deleted"})
}

// Remove a task from the Firstore database, specified by task name
// AND parent_list id <-- TO DO
//
// Example :
// http://localhost:10000/destroy/{uid}/task/{name}
// http://localhost:10000/destroy/a3a1hWUx5geKB8qeR6fbk5LZZGI2/task/test_task_1
//
// TODO: Add code to delete all sub tasks + to filter by parent id
func (a *App) destroyTask(w http.ResponseWriter, r *http.Request) {
    //fmt.Println("Endpoint Hit: destroyTask")

    // Read the variables passed
	vars := mux.Vars(r)
	uid := vars["uid"]
	name := vars["name"]

    // Create a new request for the app
    req := request.NewRequest("destroy", uid)

    // Perform the requested action
    if err := req.DestroyTask(name); err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
        return
    }
    respondWithJSON(w, http.StatusOK, map[string]string{"result": "task successfully deleted"})
}

// Get a user from the Firstore database with the specified UID
//
// Example :
// http://localhost:10000/read/{uid}
// http://localhost:10000/read/a3a1hWUx5geKB8qeR6fbk5LZZGI2
//
func (a *App) getUser(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Endpoint Hit: getUser")

    // Read the variables passed
	vars := mux.Vars(r)
	uid := vars["uid"]

    // Create a new request for the app
    req := request.NewRequest("read", uid)

    // Perform the requested action
    user, err := req.GetUser()
	if err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
        return
    }
    respondWithJSON(w, http.StatusOK, map[string]request.UserJSON{"result": *user})
}

// Get a list from the Firstore database with the specified list name
// that has an owner with the provided UID
//
// Example :
// http://localhost:10000/read/{uid}/list/{name}
//
func (a *App) getList(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Endpoint Hit: getList")

    // Read the variables passed
	vars := mux.Vars(r)
	uid := vars["uid"]
	name := vars["name"]

    // Create a new request for the app
    req := request.NewRequest("read", uid)

    // Perform the requested action
    list, err := req.GetListByName(name)
	if err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
        return
    }
    respondWithJSON(w, http.StatusOK, map[string]*request.ListJSON{"result": list})
}

// Get ALL lists from the Firstore database with that has an owner with
// the provided UID
//
// Example :
// http://localhost:10000/read/{uid}/lists
// http://localhost:10000/read/a3a1hWUx5geKB8qeR6fbk5LZZGI2/lists
//
func (a *App) getLists(w http.ResponseWriter, r *http.Request) {
    //fmt.Println("Endpoint Hit: getLists")

    // Read the variables passed
    vars := mux.Vars(r)
    uid := vars["uid"]

    // Create a new request for the app
    req := request.NewRequest("read", uid)

    // Perform the requested action
    lists, err := req.GetLists()
    if err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
        return
    }
    respondWithJSON(w, http.StatusOK, map[string][]*request.ListJSON{"result": lists})
}

// Get ALL lists from the Firstore database that have been shared with
// the requesting user
//
// Example :
// http://localhost:10000/read/{uid}/lists
// http://localhost:10000/read/a3a1hWUx5geKB8qeR6fbk5LZZGI2/lists
//
func (a *App) getSharedLists(w http.ResponseWriter, r *http.Request) {
    //fmt.Println("Endpoint Hit: getLists")

    // Read the variables passed
    vars := mux.Vars(r)
    uid := vars["uid"]

    // Create a new request for the app
    req := request.NewRequest("read", uid)

    // Perform the requested action
    lists, err := req.GetSharedLists()
    if err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
        return
    }
    respondWithJSON(w, http.StatusOK, map[string][]*request.ListJSON{"result": lists})
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
func (a *App) getTask(w http.ResponseWriter, r *http.Request) {
    //fmt.Println("Endpoint Hit: getTask")

    // Read the variables passed
    vars := mux.Vars(r)
    uid := vars["uid"]
    name := vars["name"]

    // Create a new request for the app
    req := request.NewRequest("read", uid)

    // Perform the requested action

    // Return the task
    task, err := req.GetTaskByName(name)
    if err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
        return
    }
    respondWithJSON(w, http.StatusOK, map[string]*request.TaskJSON{"result": task})
}

// Returns all tasks in a given list that the user owns
//
// Example :
// http://localhost:10000/read/{uid}/tasks/{parent_id}
// http://localhost:10000/read/a3a1hWUx5geKB8qeR6fbk5LZZGI2/tasks/NIcoux7atd3A8Lv7guUO
//
func (a *App) getTasks(w http.ResponseWriter, r *http.Request) {
    //fmt.Println("Endpoint Hit: getTasks")

    // Read the variables passed
    vars := mux.Vars(r)
    uid := vars["uid"]
    parent := vars["parent_id"]

    // Create a new request for the app
    req := request.NewRequest("read", uid)

    // Perform the requested action
    tasks, err := req.GetTasks(parent)
    if err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
        return
    }
    respondWithJSON(w, http.StatusOK, map[string][]*request.TaskJSON{"result": tasks})
}

// Update a Firestore user data
//
// Example :
// http://localhost:10000/update/{uid}?<params>
// http://localhost:10000/update/MIUVfleqSkxAtzwNeW0W?lists=qqEkD06oFudIRrCVPAc5
//
func (a *App) updateUser(w http.ResponseWriter, r *http.Request) {
    //fmt.Println("Endpoint Hit: updateUser")

    // Read the variables passed
    vars := mux.Vars(r)
    uid := vars["uid"]

    // Get the payload params and display them to the terminal
    payload := r.URL.Query()

    /*fmt.Println("\nPAYLOAD PARAMATERS")
    for k, v := range payload {
        s := fmt.Sprintf("%v => %v", k, v)
        fmt.Printf("%v\n", s)
    }*/

    // Create a new request for the app
    req := request.NewRequest("update", uid)

    // Perform the requested action
    req.GetUser()
    user, err := req.UpdateUser(payload)
    if err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
        return
    }
    respondWithJSON(w, http.StatusOK, map[string]*request.UserJSON{"result": user})
}

// Update a Firestore list data
//
// Example :
// http://localhost:10000/update/{uid}/list/{name}?<params>
// http://localhost:10000/update/a3a1hWUx5geKB8qeR6fbk5LZZGI2/list/list1?list_name=list1updated&lock=false
//
func (a *App) updateList(w http.ResponseWriter, r *http.Request) {
    //fmt.Println("Endpoint Hit: updateList")

    // Read the variables passed
    vars := mux.Vars(r)
    uid := vars["uid"]
    listname := vars["name"]
    //fmt.Printf("listname: %v\n", listname)

    // Get the payload params and display them to the terminal
    payload := r.URL.Query()

    /*fmt.Println("\nPAYLOAD PARAMATERS")
    for k, v := range payload {
        s := fmt.Sprintf("%v => %v", k, v)
        fmt.Printf("%v\n", s)
    }*/

    // Create a new request for the app
    req := request.NewRequest("update", uid)

    // Perform the requested action
    list, err := req.UpdateList(listname, payload)
    if err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
        return
    }
    respondWithJSON(w, http.StatusOK, map[string]*request.ListJSON{"result": list})
}

// Update a Firestore task data
//
// Example :
// http://localhost:10000/update/{uid}/task/{name}/{parent_id}?<params>
//
func (a *App) updateTask(w http.ResponseWriter, r *http.Request) {
    //fmt.Println("Endpoint Hit: updateTask")

    // Read the variables passed
    vars := mux.Vars(r)
    uid := vars["uid"]
    taskname := vars["name"]
    //fmt.Printf("taskname: %v\n", taskname)

    // Get the payload params and display them to the terminal
    payload := r.URL.Query()

    /*fmt.Println("\nPAYLOAD PARAMATERS")
    for k, v := range payload {
        s := fmt.Sprintf("%v => %v", k, v)
        fmt.Printf("%v\n", s)
    }*/

    // Create a new request for the app
    req := request.NewRequest("update", uid)

    // Perform the requested action
    task, err := req.UpdateTask(taskname, payload)
    if err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
        return
    }
    respondWithJSON(w, http.StatusOK, map[string]*request.TaskJSON{"result": task})
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", a.homePage)
	a.Router.HandleFunc("/create/user/{name}", a.createUser).Methods("GET", "POST")
	a.Router.HandleFunc("/create/{uid}/list/{name}", a.createList).Methods("GET", "POST")
	a.Router.HandleFunc("/create/{uid}/task/{name}", a.createTask).Methods("GET", "POST")
	a.Router.HandleFunc("/create/{uid}/subtask/{name}", a.createSubtask).Methods("GET", "POST")

	a.Router.HandleFunc("/destroy/{uid}", a.destroyUser).Methods("DELETE")
	a.Router.HandleFunc("/destroy/{uid}/list/{name}", a.destroyList).Methods("DELETE")
	a.Router.HandleFunc("/destroy/{uid}/task/{name}", a.destroyTask).Methods("DELETE")

    a.Router.HandleFunc("/read/{uid}", a.getUser).Methods("GET", "POST")
    a.Router.HandleFunc("/read/{uid}/list/{name}", a.getList).Methods("GET", "POST")
    a.Router.HandleFunc("/read/{uid}/lists", a.getLists).Methods("GET", "POST")
    a.Router.HandleFunc("/read/{uid}/shared_lists", a.getLists).Methods("GET", "POST")

    // don't need a shared_tasks function cos tasks just checks by parent.
    // list needs a shared function cos lists are user bound
    a.Router.HandleFunc("/read/{uid}/task/{name}", a.getTask).Methods("GET", "POST")
    a.Router.HandleFunc("/read/{uid}/tasks/{parent_id}", a.getTasks).Methods("GET", "POST")

    a.Router.HandleFunc("/update/{uid}", a.updateUser).Methods("PUT")
	a.Router.HandleFunc("/update/{uid}/list/{name}", a.updateList).Methods("PUT")
	a.Router.HandleFunc("/update/{uid}/task/{name}/{parent_id}", a.updateTask).Methods("PUT")
}

func main() {
    a := App{}
    a.Initialize()
    a.Run(":10000")
}
