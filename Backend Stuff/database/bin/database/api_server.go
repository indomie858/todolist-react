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
   fmt.Println("New user's name: %v\n", name)

   // Get the payload params and display them to the terminal
	payload := r.URL.Query()

   fmt.Printf("PAYLOAD PARAMATERS\n")
   for k, v := range payload {
      s := fmt.Sprintf("%v => %v", k, v)
      fmt.Println("%v\n", s)
   }

   // Create a new request
	var req request.Request
	req.Type = "create"
	req.UserId = uid
	req.Ctx = context.Background()
	req.GetClient()

   // Perform the requested action
   req.AddUser(name, payload)

   // Return the new user
   req.GetUser()
   jsonUser, _ := json.MarshalIndent(req.User,  "", "    ")
	fmt.Fprintf(w, "%v", string(jsonUser[:]))
}

// Create a new list in the Firstore database with the provided name & params
//
// Example:
// http://localhost:10000/create/{uid}/list/{name}?<params>
// http://localhost:10000/create/a3a1hWUx5geKB8qeR6fbk5LZZGI2/list/test_add_list?lock=false
//
// TO DO : Add code to update users lists array after this
func createList(w http.ResponseWriter, r *http.Request) {
   fmt.Println("Endpoint Hit: createList")

   // Read the variables passed
   vars := mux.Vars(r)
	uid := vars["uid"]
	listname := vars["name"]
   fmt.Printf("list_name: %v\n", listname)

   // Get the payload params and display them to the terminal
	payload := r.URL.Query()

   fmt.Printf("\nPAYLOAD PARAMATERS\n")
   for k, v := range payload {
      s := fmt.Sprintf("%v => %v", k, v)
      fmt.Printf("\n%v\n", s)
   }

   // Create a new request
   var req request.Request
   req.Type = "create"
   req.UserId = uid
   req.Ctx = context.Background()
   req.GetClient()

   // Perform the requested action
   req.AddList(listname, payload)

   // Return the new list
   req.GetListByName(listname)
   jsonList, _ := json.MarshalIndent(req.List,  "", "    ")
	fmt.Fprintf(w, "%v", string(jsonList[:]))
	fmt.Println(req.List)
}

// Create a new task in the Firstore database
//
// Example:
// http://localhost:10000/create/{uid}/task/{name}?<params>
// http://localhost:10000/create/a3a1hWUx5geKB8qeR6fbk5LZZGI2/task/test_task_1
//
// TO DO : Change to require list id so we can add task to user and list as well.
func createTask(w http.ResponseWriter, r *http.Request) {
   fmt.Println("Endpoint Hit: createList")

   // Read the variables passed
   vars := mux.Vars(r)
	uid := vars["uid"]
	//listname := vars["list_name"]
   taskname := vars["name"]
   fmt.Printf("task_name: %v", taskname)

   // Get the payload params and display them to the terminal
	payload := r.URL.Query()

   fmt.Printf("\nPAYLOAD PARAMATERS\n")
   for k, v := range payload {
      s := fmt.Sprintf("%v => %v", k, v)
      fmt.Printf("\n%v\n", s)
   }

   // Create a new request
   var req request.Request
   req.Type = "create"
   req.UserId = uid
   req.Ctx = context.Background()
   req.GetClient()

   // Perform the requested action
   req.AddTask(taskname, payload)

   // Return the new task
   req.GetTaskByName(taskname)
   jsonList, _ := json.MarshalIndent(req.Task,  "", "    ")
	fmt.Fprintf(w, "%v", string(jsonList[:]))
	fmt.Println(req.Task)
}

// Create a new sub task in the Firstore database with the provided name
//
// Example:
// http://localhost:10000/create/{uid}/subtask/{name}
//
// TO DO: ALL
func createSubtask(w http.ResponseWriter, r *http.Request) {

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

   // Create a new request
	var req request.Request
	req.Type = "destroy"
	req.UserId = uid
	req.Ctx = context.Background()
	req.GetClient()

   // Perform the requested action
   err := req.DestroyUser()

   // Return the result of the delete
   if err != nil {
      fmt.Fprintf(w, "err deleting user: %v", err)
      fmt.Printf("ERR deleting user: %v", err)
   }
   fmt.Printf("user successfully deleted")
	fmt.Fprintf(w, "user successfully deleted")

}

// Remove a list from the Firstore database, specified by list name
//
// Example :
// http://localhost:10000/destroy/{uid}/list/{name}
// http://localhost:10000/destroy/MIUVfleqSkxAtzwNeW0W/list/first_list
//
// TO DO : Add code to delete all tasks and subtasks
func destroyList(w http.ResponseWriter, r *http.Request) {
   fmt.Println("Endpoint Hit: destroyList")

   // Read the variables passed
	vars := mux.Vars(r)
	uid := vars["uid"]
	name := vars["name"]

   // Create a new request
	var req request.Request
	req.Type = "destroy"
	req.UserId = uid
	req.Ctx = context.Background()
	req.GetClient()

   // Perform the requested action
	req.GetListByName(name)
   err := req.DestroyList()

   // Return the result of the delete
   if err != nil {
      fmt.Fprintf(w, "err deleting list: %v", err)
      log.Printf("ERR deleting list: %v", err)
   }

   fmt.Printf("list successfully deleted")
	fmt.Fprintf(w, "list successfully deleted")
}

// Remove a task from the Firstore database, specified by task name
// AND parent_list id <-- TO DO
//
// Example :
// http://localhost:10000/destroy/{uid}/task/{name}
// http://localhost:10000/destroy/a3a1hWUx5geKB8qeR6fbk5LZZGI2/task/test_task_1
//
// TO DO : Add code to delete all sub tasks + to filter by parent id
func destroyTask(w http.ResponseWriter, r *http.Request) {
   fmt.Println("Endpoint Hit: destroyTask")

   // Read the variables passed
	vars := mux.Vars(r)
	uid := vars["uid"]
	name := vars["name"]

   // Create a new request
	var req request.Request
	req.Type = "destroy"
	req.UserId = uid
	req.Ctx = context.Background()
	req.GetClient()

   // Perform the requested action
	req.GetTaskByName(name)
   err := req.DestroyTask()

   // Return the result of the delete
   if err != nil {
      fmt.Fprintf(w, "err deleting task: %v", err)
      log.Printf("ERR deleting task: %v", err)
   }

   fmt.Printf("task successfully deleted")
	fmt.Fprintf(w, "task successfully deleted")
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

   // Create a new request
	var req request.Request
	req.Type = "read"
	req.UserId = uid
	req.Ctx = context.Background()
	req.GetClient()

   // Perform the requested action
	req.GetUser()

   // Return the user
	jsonUser, _ := json.MarshalIndent(req.User, "", "    ")
	fmt.Fprintf(w, "%v", string(jsonUser[:]))
	fmt.Println(req.User)
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

   // Create a new request
	var req request.Request
	req.Type = "read"
	req.UserId = uid
	req.Ctx = context.Background()

   // Perform the requested action
   req.GetClient()

   // Return the list
	req.GetListByName(name)
	jsonList, _ := json.MarshalIndent(req.List, "", "    ")
	fmt.Fprintf(w, "%v", string(jsonList[:]))
	fmt.Println(req.List)
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

   // Create a new request
   var req request.Request
   req.Type = "read"
   req.UserId = uid
   req.Ctx = context.Background()
   req.GetClient()

   // Perform the requested action
   lists := req.GetLists()

   // Return the lists
   jsonLists, _ := json.MarshalIndent(lists, "", "    ")
   fmt.Fprintf(w, "%v", string(jsonLists[:]))
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

   // Create a new request
   var req request.Request
   req.Type = "read"
   req.UserId = uid
   req.Ctx = context.Background()

   // Perform the requested action
   req.GetClient()

   // Return the task
   req.GetTaskByName(name)
   jsonList, _ := json.MarshalIndent(req.Task, "", "    ")
   fmt.Fprintf(w, "%v", string(jsonList[:]))
   fmt.Println(req.List)
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

   // Create a new request
   var req request.Request
   req.Type = "read"
   req.UserId = uid
   req.Ctx = context.Background()
   req.GetClient()

   // Perform the requested action
   tasks := req.GetTasks(parent)

   // Return the task
   jsonLists, _ := json.MarshalIndent(tasks, "", "    ")
   fmt.Fprintf(w, "%v", string(jsonLists[:]))
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

   // Create a new request
   var req request.Request
   req.Type = "update"
   req.UserId = uid
   req.Ctx = context.Background()
   req.GetClient()

   // Perform the requested action
   req.GetUser()
   req.UpdateUser(payload)

   // Return the updated user
   req.GetUser()
	jsonUser, _ := json.MarshalIndent(req.User, "", "    ")
	fmt.Fprintf(w, "%v", string(jsonUser[:]))
	fmt.Println(req.User)
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

   // Create a new request
   var req request.Request
   req.Type = "update"
   req.UserId = uid
   req.Ctx = context.Background()
   req.GetClient()

   // Perform the requested action
	req.GetListByName(listname)
   req.UpdateList(payload)

   // Return the updated list
   req.GetListByID()
   jsonList, _ := json.MarshalIndent(req.List,  "", "    ")
	fmt.Fprintf(w, "%v", string(jsonList[:]))
	fmt.Println(req.List)
}

// Update a Firestore task data
//
// Example :
// http://localhost:10000/update/{uid}/task/{task}?<params>
//
// TO DO: ALL 
func updateTask(w http.ResponseWriter, r *http.Request) {

}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)

	router.HandleFunc("/create/user/{name}", createUser).Methods("GET", "POST")
	router.HandleFunc("/create/{uid}/list/{name}", createList).Methods("GET", "POST")
	router.HandleFunc("/create/{uid}/task/{name}", createTask).Methods("GET", "POST")
	router.HandleFunc("/create/{uid}/subtask/{name}", createSubtask).Methods("GET", "POST")

	router.HandleFunc("/destroy/{uid}", destroyUser).Methods("GET", "POST")
	router.HandleFunc("/destroy/{uid}/list/{name}", destroyList).Methods("GET", "POST")
	router.HandleFunc("/destroy/{uid}/task/{name}", destroyTask).Methods("GET", "POST")

	router.HandleFunc("/read/{uid}", getUser).Methods("GET", "POST")

   router.HandleFunc("/read/{uid}/list/{name}", getList).Methods("GET", "POST")
   router.HandleFunc("/read/{uid}/lists", getLists).Methods("GET", "POST")

   router.HandleFunc("/read/{uid}/task/{name}", getTask).Methods("GET", "POST")
   router.HandleFunc("/read/{uid}/tasks/{parent_id}", getTasks).Methods("GET", "POST")

	router.HandleFunc("/update/{uid}", updateUser).Methods("GET", "POST")
	router.HandleFunc("/update/{uid}/list/{list}", updateList).Methods("GET", "POST")
	router.HandleFunc("/update/{uid}/task/{task}", updateTask).Methods("GET", "POST")

	log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
	handleRequests()
}
