package main

import (
	"database/request"

	"fmt"
	"log"
	"net/http"
   "context"

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

   fmt.Fprintf(w,"\n")

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

// TO DO:
func createTask(w http.ResponseWriter, r *http.Request) {

}

// TO DO:
func createSubtask(w http.ResponseWriter, r *http.Request) {

}

// TO DO:
func destroyUser(w http.ResponseWriter, r *http.Request) {

}

// TO DO:
func destroyList(w http.ResponseWriter, r *http.Request) {

}

// TO DO:
func destroyTask(w http.ResponseWriter, r *http.Request) {

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
   jsonUser, _ := json.MarshalIndent(req.User,  "", "    ")
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
   jsonList, _ := json.MarshalIndent(req.List,  "", "    ")
	fmt.Fprintf(w, "%v", string(jsonList[:]))
	fmt.Println(req.List)
}

// TO DO:
func getTask(w http.ResponseWriter, r *http.Request) {

}

// TO DO:
func updateUser(w http.ResponseWriter, r *http.Request) {

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
	router.HandleFunc("/create/user/{name}", createUser)
   router.HandleFunc("/create/{uid}/list/{name}", createList).Methods("GET", "POST")
   router.HandleFunc("/create/{uid}/task/{name}", createTask)
   router.HandleFunc("/create/{uid}/subtask/{name}", createSubtask)

   router.HandleFunc("/destroy/{uid}", destroyUser)
   router.HandleFunc("/destroy/list/{lists}", destroyList)
   router.HandleFunc("/destroy/task/{tasks}", destroyTask)

   router.HandleFunc("/read/{uid}", getUser)
   router.HandleFunc("/read/{uid}/list/{name}", getList)
   router.HandleFunc("/read/{uid}/task/{name}", getTask)

   router.HandleFunc("/update/{uid}", updateUser)
   router.HandleFunc("/update/{uid}/list/{list}", updateList).Methods("GET", "POST")

   router.HandleFunc("/update/{uid}/task/{task}", updateTask)

	log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
	handleRequests()
}
