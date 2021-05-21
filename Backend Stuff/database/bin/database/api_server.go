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
    Request *request.Request
}

// src: https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql
func (a *App) Initialize() {
    a.Router = mux.NewRouter().StrictSlash(true)
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
    log.Printf("Running on port %s", addr)
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

type Result struct {
    User *request.UserJSON
    List *request.ListJSON
    Lists []*request.ListJSON
    Task *request.TaskJSON
    Tasks []*request.TaskJSON
    AllTasks [][]*request.TaskJSON
}

// Create a new user in the Firstore database wih the provided name
//
// Example:
// http://localhost:10000/create/user/{uid}
func (a *App) createUser(w http.ResponseWriter, r *http.Request) {
    //fmt.Println("Endpoint Hit: createUser")

    // Read the variables passed
    vars := mux.Vars(r)
	uid := vars["uid"]

    // Create a new request for the app
    a.Request = request.NewRequest("create", uid)

    // Get the payload params and display them to the terminal
	payload := r.URL.Query()

    /*fmt.Println("\nPAYLOAD PARAMATERS")
    for k, v := range payload {
        s := fmt.Sprintf("%v => %v", k, v)
        fmt.Printf("%v\n", s)
    }*/

    // Perform the requested action
    user, err := a.Request.AddUser(uid, payload)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    lists, err := a.Request.GetLists()
    if err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
        return
    }

    var tasks [][]*request.TaskJSON
    for _, list := range lists {
        t, _ := a.Request.GetTasks(list.Id)
        tasks = append(tasks, t)
    }

    var res Result
    res.User = user
    res.Lists = lists
    res.AllTasks = tasks

    respondWithJSON(w, http.StatusOK, map[string]*Result{"result": &res})
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
    a.Request = request.NewRequest("create", uid)

    // Get the payload params and display them to the terminal
	payload := r.URL.Query()

    /*fmt.Println("\nPAYLOAD PARAMATERS")
    for k, v := range payload {
        s := fmt.Sprintf("%v => %v", k, v)
        fmt.Printf("\n%v\n", s)
    }*/

    list, err := a.Request.AddList(listname, payload)
    // Perform the requested action
    if err != nil {
        respondWithError(w, http.StatusBadRequest, err.Error())
        return
    }
    tasks, _ := a.Request.GetTasks(list.Id)

    var res Result
    res.List = list
    res.Tasks = tasks

    respondWithJSON(w, http.StatusOK, map[string]*Result{"result": &res})
}

// Create a new task in the Firstore database
//
// Example:
// http://localhost:10000/create/{uid}/task/{name}/parent/{pid}?<params>
// http://localhost:10000/create/a3a1hWUx5geKB8qeR6fbk5LZZGI2/task/test_task_1/parent/asdjfasdlfja;ldfj
//
// TODO: Change to require list id so we can add task to user and list as well.
func (a *App) createTask(w http.ResponseWriter, r *http.Request) {
    //fmt.Println("Endpoint Hit: createTask")

    // Read the variables passed
    vars := mux.Vars(r)
	uid := vars["uid"]
    pid := vars["pid"]
    taskname := vars["name"]
    //fmt.Printf("task_name: %v", taskname)

    // Create a new request for the app
    a.Request = request.NewRequest("create", uid)

    // Get the payload params and display them to the terminal
	payload := r.URL.Query()

    /*fmt.Println("\nPAYLOAD PARAMATERS")
    for k, v := range payload {
        s := fmt.Sprintf("%v => %v", k, v)
        fmt.Printf("\n%v\n", s)
    }*/

    // Perform the requested action
    task, err := a.Request.AddTask(taskname, pid, payload)
    if err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
        return
    }

    var res Result
    res.Task = task

    respondWithJSON(w, http.StatusOK, map[string]*Result{"result": &res})
}

// Create a new sub task in the Firstore database with the provided name
//
// Example:
// http://localhost:10000/create/{uid}/subtask/{name}/parent/{pid}?<params>
//
func (a *App) createSubtask(w http.ResponseWriter, r *http.Request) {
    //fmt.Println("Endpoint Hit: createSubtask")

    // Read the variables passed
    vars := mux.Vars(r)
	uid := vars["uid"]
    pid := vars["pid"]
    taskname := vars["name"]
    //fmt.Printf("task_name: %v", taskname)

    // Create a new request for the app
    a.Request = request.NewRequest("create", uid)

    // Get the payload params and display them to the terminal
	payload := r.URL.Query()

    /*fmt.Println("\nPAYLOAD PARAMATERS")
    for k, v := range payload {
        s := fmt.Sprintf("%v => %v", k, v)
        fmt.Printf("\n%v\n", s)
    }*/
    payload.Add("sub_task", "true")

    // Perform the requested action
    task, err := a.Request.UpdateTaskSubtasks(pid, taskname)
    if err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
        return
    }

    var res Result
    res.Task = task

    respondWithJSON(w, http.StatusOK, map[string]*Result{"result": &res})
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
    a.Request = request.NewRequest("destroy", uid)

    // Perform the requested action
    if err := a.Request.DestroyUser(); err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
        return
    }
    respondWithJSON(w, http.StatusOK, map[string]string{"result": "user successfully deleted"})
}

// Remove a list from the Firstore database, specified by list name
//
// Example :
// http://localhost:10000/destroy/{uid}/list/{id}
// http://localhost:10000/destroy/MIUVfleqSkxAtzwNeW0W/list/first_list
//
// TODO: Add code to delete all tasks and subtasks
func (a *App) destroyList(w http.ResponseWriter, r *http.Request) {
    //fmt.Println("Endpoint Hit: destroyList")

    // Read the variables passed
	vars := mux.Vars(r)
	uid := vars["uid"]
	id := vars["id"]

    // Create a new request for the app
    a.Request = request.NewRequest("destroy", uid)

    // Perform the requested action
    if err := a.Request.DestroyListById(id); err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
        return
    }
    respondWithJSON(w, http.StatusOK, map[string]string{"result": "list successfully deleted"})
}

// Remove a task from the Firstore database, specified by task name
// AND parent_list id <-- TO DO
//
// Example :
// http://localhost:10000/destroy/{uid}/task/{id}
// http://localhost:10000/destroy/f9oXnGYUlUADNIDambFG/task/test_task_1/parent/hsHYrOZeeAAuIAOSWaLk/
//
// TODO: Add code to delete all sub tasks + to filter by parent id
func (a *App) destroyTask(w http.ResponseWriter, r *http.Request) {
    //fmt.Println("Endpoint Hit: destroyTask")

    // Read the variables passed
	vars := mux.Vars(r)
	uid := vars["uid"]
    id := vars["id"]

    // Create a new request for the app
    a.Request = request.NewRequest("destroy", uid)

    // Perform the requested action
    if err := a.Request.DestroyTaskById(id); err != nil {
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
    a.Request = request.NewRequest("read", uid)

    // Perform the requested action
    user, err := a.Request.GetUser()
	if err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
        return
    }

    lists, err := a.Request.GetLists()
    if err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
        return
    }

    var tasks [][]*request.TaskJSON
    for _, list := range lists {
        t, _ := a.Request.GetTasks(list.Id)
        tasks = append(tasks, t)
    }

    var res Result
    res.User = user
    res.Lists = lists
    res.AllTasks = tasks

    respondWithJSON(w, http.StatusOK, map[string]*Result{"result": &res})
}

// Get all the tasks with reminders
//
// Example :
// http://localhost:10000/readusers
// http://localhost:10000/read/taskreminders
func (a *App) getAllUsers(w http.ResponseWriter, r *http.Request) {
    //fmt.Println("Endpoint Hit: updateUser")

    // Read the variables passed
    //vars := mux.Vars(r)
    //uid := vars["uid"]
    //now := vars["now"]

    // Create a new request for the app
    a.Request = request.NewRequest("read", "read_all_users")

    users, err := a.Request.GetAllUsers()
    if err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
        return
    }

    respondWithJSON(w, http.StatusOK, map[string][]*request.UserJSON{"users": users})

}

// Get a list from the Firstore database with the specified list name
// that has an owner with the provided UID
//
// Example :
// http://localhost:10000/read/{uid}/list/{id}
//
func (a *App) getList(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Endpoint Hit: getList")

    // Read the variables passed
	vars := mux.Vars(r)
	uid := vars["uid"]
	id := vars["id"]

    // Create a new request for the app
    a.Request = request.NewRequest("read", uid)

    // Perform the requested action
    list, err := a.Request.GetListByID(id)
	if err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
        return
    }

    tasks, _ := a.Request.GetTasks(id)

    var res Result
    res.List = list
    res.Tasks = tasks

    respondWithJSON(w, http.StatusOK, map[string]*Result{"result": &res})
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
    a.Request = request.NewRequest("read", uid)

    // Perform the requested action
    lists, err := a.Request.GetLists()
    if err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
        return
    }

    var tasks [][]*request.TaskJSON
    for _, list := range lists {
        t, _ := a.Request.GetTasks(list.Id)
        tasks = append(tasks, t)
    }

    var res Result
    res.Lists = lists
    res.AllTasks = tasks

    respondWithJSON(w, http.StatusOK, map[string]*Result{"result": &res})
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
    a.Request = request.NewRequest("read", uid)

    // Perform the requested action
    lists, err := a.Request.GetSharedLists()
    if err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
        return
    }
    var tasks [][]*request.TaskJSON
    for _, list := range lists {
        t, _ := a.Request.GetTasks(list.Id)
        tasks = append(tasks, t)
    }

    var res Result
    res.Lists = lists
    res.AllTasks = tasks

    respondWithJSON(w, http.StatusOK, map[string]*Result{"result": &res})
}

// Get a task from the Firstore database with the specified task name
// that has an owner with the provided UID
// AND has the same parent_id as the one provided <--- TO DO
// in case user names a bunch of tasks the same thing just in diff. lists
//
// Example :
// http://localhost:10000/read/{uid}/task/{id}
// http://localhost:10000/read/a3a1hWUx5geKB8qeR6fbk5LZZGI2/task/task1/list/thgvhhcyresjc
//
func (a *App) getTask(w http.ResponseWriter, r *http.Request) {
    //fmt.Println("Endpoint Hit: getTask")

    // Read the variables passed
    vars := mux.Vars(r)
    uid := vars["uid"]
    id := vars["id"]

    // Create a new request for the app
    a.Request = request.NewRequest("read", uid)

    // Perform the requested action

    // Return the task
    task, err := a.Request.GetTaskByID(id)
    if err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
        return
    }

    var res Result
    res.Task = task

    respondWithJSON(w, http.StatusOK, map[string]*Result{"result": &res})
}

// Returns all tasks in a given list that the user owns
//
// Example :
// http://localhost:10000/read/{uid}/tasks/{pid}
// http://localhost:10000/read/a3a1hWUx5geKB8qeR6fbk5LZZGI2/tasks/NIcoux7atd3A8Lv7guUO
//
func (a *App) getTasks(w http.ResponseWriter, r *http.Request) {
    //fmt.Println("Endpoint Hit: getTasks")

    // Read the variables passed
    vars := mux.Vars(r)
    uid := vars["uid"]
    parent := vars["pid"]

    // Create a new request for the app
    a.Request = request.NewRequest("read", uid)

    // Perform the requested action
    tasks, err := a.Request.GetTasks(parent)
    if err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
        return
    }

    var res Result
    res.Tasks = tasks

    respondWithJSON(w, http.StatusOK, map[string]*Result{"result": &res})
}

// Get all the tasks with reminders
//
// Example :
// http://localhost:10000/readtaskreminders
// http://localhost:10000/read/taskreminders
func (a *App) getRemindTasks(w http.ResponseWriter, r *http.Request) {
    //fmt.Println("Endpoint Hit: updateUser")

    // Read the variables passed
    //vars := mux.Vars(r)
    //uid := vars["uid"]
    //now := vars["now"]

    // Create a new request for the app
    a.Request = request.NewRequest("read", "reminder_read")

    tasks, err := a.Request.GetRemindTasks()
    if err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
        return
    }

    var res Result
    res.Tasks = tasks

		//Tasks []*request.TaskJSON

    respondWithJSON(w, http.StatusOK, &res.Tasks)
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
    a.Request = request.NewRequest("update", uid)

    // Perform the requested action
    a.Request.GetUser()
    user, err := a.Request.UpdateUser(payload)
    if err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
        return
    }

    lists, err := a.Request.GetLists()
    if err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
        return
    }

    var tasks [][]*request.TaskJSON
    for _, list := range lists {
        t, _ := a.Request.GetTasks(list.Id)
        tasks = append(tasks, t)
    }

    var res Result
    res.User = user
    res.Lists = lists
    res.AllTasks = tasks

    respondWithJSON(w, http.StatusOK, map[string]*Result{"result": &res})
}

// Update a Firestore list data
//
// Example :
// http://localhost:10000/update/{uid}/list/{id}?<params>
// http://localhost:10000/update/a3a1hWUx5geKB8qeR6fbk5LZZGI2/list/isolated_test_list?list_name=updated_isolated_list&lock=false&shared=true
//
func (a *App) updateList(w http.ResponseWriter, r *http.Request) {
    //fmt.Println("Endpoint Hit: updateList")

    // Read the variables passed
    vars := mux.Vars(r)
    uid := vars["uid"]
    id := vars["id"]
    //fmt.Printf("listname: %v\n", listname)

    // Get the payload params and display them to the terminal
    payload := r.URL.Query()

    /*fmt.Println("\nPAYLOAD PARAMATERS")
    for k, v := range payload {
        s := fmt.Sprintf("%v => %v", k, v)
        fmt.Printf("%v\n", s)
    }*/

    // Create a new request for the app
    a.Request = request.NewRequest("update", uid)

    // Perform the requested action
    list, err := a.Request.UpdateList(id, payload)
    if err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
        return
    }

    tasks, _ := a.Request.GetTasks(id)

    var res Result
    res.List = list
    res.Tasks = tasks

    respondWithJSON(w, http.StatusOK, map[string]*Result{"result": &res})
}

// Update a Firestore task data
//
// Example :
// http://localhost:10000/update/{uid}/task/{id}?<params>
//
func (a *App) updateTask(w http.ResponseWriter, r *http.Request) {
    //fmt.Println("Endpoint Hit: updateTask")

    // Read the variables passed
    vars := mux.Vars(r)
    uid := vars["uid"]
    id := vars["id"]
    //fmt.Printf("taskname: %v\n", taskname)

    // Get the payload params and display them to the terminal
    payload := r.URL.Query()

    /*fmt.Println("\nPAYLOAD PARAMATERS")
    for k, v := range payload {
        s := fmt.Sprintf("%v => %v", k, v)
        fmt.Printf("%v\n", s)
    }*/

    // Create a new request for the app
    a.Request = request.NewRequest("update", uid)

    // Perform the requested action
    task, err := a.Request.UpdateTask(id, payload)
    if err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
        return
    }

    var res Result
    res.Task = task

    respondWithJSON(w, http.StatusOK, map[string]*Result{"result": &res})
}

func (a *App) initializeRoutes() {
    a.Router.HandleFunc("/", a.homePage)

    // Create functions
    a.Router.HandleFunc("/create/user/{uid}", a.createUser).Methods("GET", "POST")
    a.Router.HandleFunc("/create/{uid}/list/{name}", a.createList).Methods("GET", "POST")
    a.Router.HandleFunc("/create/{uid}/task/{name}/parent/{pid}", a.createTask).Methods("GET", "POST")
    a.Router.HandleFunc("/create/{uid}/subtask/{name}/parent/{pid}", a.createSubtask).Methods("GET", "POST")

    // Destroy functions
    // We can use one for destroying both tasks and subtasks due to requiring the parent id
    a.Router.HandleFunc("/destroy/{uid}", a.destroyUser).Methods("GET","DELETE")
    a.Router.HandleFunc("/destroy/{uid}/list/{id}", a.destroyList).Methods("GET", "DELETE")
    a.Router.HandleFunc("/destroy/{uid}/task/{id}", a.destroyTask).Methods("GET", "DELETE")

    // Read functions
    // Only one for tasks & subtaks, as we get both using just the parent id, not user id
    a.Router.HandleFunc("/read/{uid}", a.getUser).Methods("GET", "POST")
    a.Router.HandleFunc("/read/{uid}/list/{id}", a.getList).Methods("GET", "POST")
    a.Router.HandleFunc("/read/{uid}/lists", a.getLists).Methods("GET", "POST")
    a.Router.HandleFunc("/read/{uid}/shared_lists", a.getLists).Methods("GET", "POST")
    a.Router.HandleFunc("/read/{uid}/task/{id}", a.getTask).Methods("GET", "POST")
    a.Router.HandleFunc("/read/{uid}/tasks/{pid}", a.getTasks).Methods("GET", "POST")
    a.Router.HandleFunc("/readtaskreminders", a.getRemindTasks).Methods("GET", "POST")
    a.Router.HandleFunc("/readusers", a.getAllUsers).Methods("GET", "POST")

    // Update functions
    a.Router.HandleFunc("/update/{uid}", a.updateUser).Methods("GET", "PUT", "POST")
    a.Router.HandleFunc("/update/{uid}/list/{id}", a.updateList).Methods("GET", "PUT", "POST")
    a.Router.HandleFunc("/update/{uid}/task/{id}", a.updateTask).Methods("GET", "PUT", "POST")
}

func main() {
    a := App{}
    a.Initialize()
    a.Run(":10000")
}
