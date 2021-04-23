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

func createUser(w http.ResponseWriter, r *http.Request) {
   fmt.Println("Endpoint Hit: createUser")
   vars := mux.Vars(r)

	uid := vars["uid"]
	name := vars["name"]

   fmt.Printf("New user's name: %v\n\n", name)

	payload := r.URL.Query()

   fmt.Printf("PAYLOAD PARAMATERS\n")
   for k, v := range payload {
      s := fmt.Sprintf("%v => %v", k, v)
      fmt.Printf("%v\n", s)
   }

	fmt.Fprintf(w, "\n")

	var req request.Request
	req.Type = "create"
	req.UserId = uid
	req.Ctx = context.Background()
	req.GetClient()

   req.AddUser(name, payload)
   req.GetUser()

   jsonUser, _ := json.MarshalIndent(req.User,  "", "    ")
	fmt.Fprintf(w, "%v", string(jsonUser[:]))
}

// TO DO : Add code to update users lists array after this
func createList(w http.ResponseWriter, r *http.Request) {
   fmt.Println("Endpoint Hit: createList")
   vars := mux.Vars(r)

	uid := vars["uid"]
	listname := vars["name"]

   fmt.Printf("list_name: %v\n", listname)

	payload := r.URL.Query()

   fmt.Printf("\nPAYLOAD PARAMATERS\n")
   for k, v := range payload {
      s := fmt.Sprintf("%v => %v", k, v)
      fmt.Printf("\n%v\n", s)
   }


   var req request.Request
   req.Type = "create"
   req.UserId = uid
   req.Ctx = context.Background()
   req.GetClient()

   req.AddList(listname, payload)
   req.GetListByName(listname)

   jsonList, _ := json.MarshalIndent(req.List,  "", "    ")
	fmt.Fprintf(w, "%v", string(jsonList[:]))
	fmt.Println(req.List)
}

// TO DO : Change to require list id so we can add task to user and list as well.
func createTask(w http.ResponseWriter, r *http.Request) {
   fmt.Println("Endpoint Hit: createList")
   vars := mux.Vars(r)

	uid := vars["uid"]
	//listname := vars["list_name"]
   taskname := vars["name"]

   //fmt.Printf("list_name: %v\ntask_name: %v", listname, taskname)

	payload := r.URL.Query()

   fmt.Printf("\nPAYLOAD PARAMATERS\n")
   for k, v := range payload {
      s := fmt.Sprintf("%v => %v", k, v)
      fmt.Printf("\n%v\n", s)
   }

   var req request.Request
   req.Type = "create"
   req.UserId = uid
   req.Ctx = context.Background()
   req.GetClient()

   req.AddTask(taskname, payload)
   req.GetTaskByName(taskname)

   jsonList, _ := json.MarshalIndent(req.Task,  "", "    ")
	fmt.Fprintf(w, "%v", string(jsonList[:]))
	fmt.Println(req.Task)
}

// TO DO:
func createSubtask(w http.ResponseWriter, r *http.Request) {

}

func destroyUser(w http.ResponseWriter, r *http.Request) {
   fmt.Println("Endpoint Hit: destroyUser")

	vars := mux.Vars(r)
	uid := vars["uid"]

	var req request.Request
	req.Type = "destroy"
	req.UserId = uid
	req.Ctx = context.Background()
	req.GetClient()

   err := req.DestroyUser()
   if err != nil {
      fmt.Fprintf(w, "err deleting user: %v", err)
      fmt.Printf("ERR deleting user: %v", err)
   }
   fmt.Printf("user successfully deleted")
	fmt.Fprintf(w, "user successfully deleted")

}

// TO DO : Add code to delete all tasks and subtasks
func destroyList(w http.ResponseWriter, r *http.Request) {
   fmt.Println("Endpoint Hit: destroyList")

	vars := mux.Vars(r)
	uid := vars["uid"]
	name := vars["name"]

	var req request.Request
	req.Type = "destroy"
	req.UserId = uid
	req.Ctx = context.Background()
	req.GetClient()

	req.GetListByName(name)
   err := req.DestroyList()
   if err != nil {
      fmt.Fprintf(w, "err deleting list: %v", err)
      log.Printf("ERR deleting list: %v", err)
   }

   fmt.Printf("list successfully deleted")
	fmt.Fprintf(w, "list successfully deleted")
}

// TO DO : Add code to delete all sub tasks
func destroyTask(w http.ResponseWriter, r *http.Request) {
   fmt.Println("Endpoint Hit: destroyTask")

	vars := mux.Vars(r)
	uid := vars["uid"]
	name := vars["name"]

	var req request.Request
	req.Type = "destroy"
	req.UserId = uid
	req.Ctx = context.Background()
	req.GetClient()

	req.GetTaskByName(name)
   err := req.DestroyTask()
   if err != nil {
      fmt.Fprintf(w, "err deleting task: %v", err)
      log.Printf("ERR deleting task: %v", err)
   }

   fmt.Printf("task successfully deleted")
	fmt.Fprintf(w, "task successfully deleted")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getUser")

	vars := mux.Vars(r)
	uid := vars["uid"]

	var req request.Request
	req.Type = "read"
	req.UserId = uid
	req.Ctx = context.Background()
	req.GetClient()

	req.GetUser()
	jsonUser, _ := json.MarshalIndent(req.User, "", "    ")
	fmt.Fprintf(w, "%v", string(jsonUser[:]))
	fmt.Println(req.User)
}

func getList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getList")

	vars := mux.Vars(r)
	uid := vars["uid"]
	name := vars["name"]

	var req request.Request
	req.Type = "read"
	req.UserId = uid
	req.Ctx = context.Background()

   req.GetClient()
	req.GetListByName(name)
	jsonList, _ := json.MarshalIndent(req.List, "", "    ")
	fmt.Fprintf(w, "%v", string(jsonList[:]))
	fmt.Println(req.List)
}

func getLists(w http.ResponseWriter, r *http.Request) {
   fmt.Println("Endpoint Hit: getLists")

   vars := mux.Vars(r)
   uid := vars["uid"]

   var req request.Request
   req.Type = "read"
   req.UserId = uid
   req.Ctx = context.Background()
   req.GetClient()

   lists := req.GetLists()
   jsonLists, _ := json.MarshalIndent(lists, "", "    ")
   fmt.Fprintf(w, "%v", string(jsonLists[:]))
}

func getTask(w http.ResponseWriter, r *http.Request) {
   fmt.Println("Endpoint Hit: getTask")

   vars := mux.Vars(r)
   uid := vars["uid"]
   name := vars["name"]

   var req request.Request
   req.Type = "read"
   req.UserId = uid
   req.Ctx = context.Background()

   req.GetClient()
   req.GetTaskByName(name)
   jsonList, _ := json.MarshalIndent(req.Task, "", "    ")
   fmt.Fprintf(w, "%v", string(jsonList[:]))
   fmt.Println(req.List)
}

func getTasks(w http.ResponseWriter, r *http.Request) {
   fmt.Println("Endpoint Hit: getTasks")

   vars := mux.Vars(r)
   uid := vars["uid"]
   parent := vars["parent_id"]

   var req request.Request
   req.Type = "read"
   req.UserId = uid
   req.Ctx = context.Background()
   req.GetClient()

   tasks := req.GetTasks(parent)
   jsonLists, _ := json.MarshalIndent(tasks, "", "    ")
   fmt.Fprintf(w, "%v", string(jsonLists[:]))
}

func updateUser(w http.ResponseWriter, r *http.Request) {
   fmt.Println("Endpoint Hit: updateUser")

   vars := mux.Vars(r)
   uid := vars["uid"]

   payload := r.URL.Query()

   fmt.Printf("\nPAYLOAD PARAMATERS\n")
   for k, v := range payload {
      s := fmt.Sprintf("%v => %v", k, v)
      fmt.Printf("%v\n", s)
   }

   var req request.Request
   req.Type = "update"
   req.UserId = uid
   req.Ctx = context.Background()
   req.GetClient()

   req.GetUser()
   req.UpdateUser(payload)
   req.GetUser()
	jsonUser, _ := json.MarshalIndent(req.User, "", "    ")
	fmt.Fprintf(w, "%v", string(jsonUser[:]))
	fmt.Println(req.User)
}

func updateList(w http.ResponseWriter, r *http.Request) {
   fmt.Println("Endpoint Hit: updateList")

   vars := mux.Vars(r)
   uid := vars["uid"]
   listname := vars["list"]
   fmt.Printf("listname: %v\n", listname)

   payload := r.URL.Query()

   fmt.Printf("\nPAYLOAD PARAMATERS\n")
   for k, v := range payload {
      s := fmt.Sprintf("%v => %v", k, v)
      fmt.Printf("%v\n", s)
   }

   var req request.Request
   req.Type = "update"
   req.UserId = uid
   req.Ctx = context.Background()
   req.GetClient()

	req.GetListByName(listname)
   req.UpdateList(payload)
   req.GetListByID()
   jsonList, _ := json.MarshalIndent(req.List,  "", "    ")

	fmt.Fprintf(w, "%v", string(jsonList[:]))
	fmt.Println(req.List)
}

// TO DO:
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
