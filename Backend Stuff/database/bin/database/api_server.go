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

}

func createList(w http.ResponseWriter, r *http.Request) {

}

func createTask(w http.ResponseWriter, r *http.Request) {

}

func createSubtask(w http.ResponseWriter, r *http.Request) {

}

func destroyUser(w http.ResponseWriter, r *http.Request) {

}

func destroyList(w http.ResponseWriter, r *http.Request) {

}

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

func getTask(w http.ResponseWriter, r *http.Request) {

}

func editUser(w http.ResponseWriter, r *http.Request) {

}

func editList(w http.ResponseWriter, r *http.Request) {

}

func editTask(w http.ResponseWriter, r *http.Request) {

}

func updateUser(w http.ResponseWriter, r *http.Request) {

}

func updateList(w http.ResponseWriter, r *http.Request) {
   fmt.Println("Endpoint Hit: updateList")

   vars := mux.Vars(r)
   uid := vars["uid"]

   var list []*request.ListJSON
   t := []byte(vars["list"])
   json.Unmarshal(t, &list)

   listname := list[0].Name
   fmt.Printf(listname)

   var req request.Request
   req.Type = "update"
   req.UserId = uid
   req.Ctx = context.Background()
   req.GetClient()
	req.GetListByName(listname)

   //req.UpdateList(fields)

}

func updateTask(w http.ResponseWriter, r *http.Request) {

}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	router.HandleFunc("/create/user/{name}", createUser)
   router.HandleFunc("/create/list/{uid}/{name}", createList)
   router.HandleFunc("/create/task/{uid}/{name}", createTask)
   router.HandleFunc("/create/task/{uid}/{name}/", createSubtask)

   router.HandleFunc("/destroy/{uid}", destroyUser)
   router.HandleFunc("/destroy/list/{lists}", destroyList)
   router.HandleFunc("/destroy/task/{tasks}", destroyTask)

   router.HandleFunc("/read/{uid}", getUser)
   router.HandleFunc("/read/list/{uid}/{name}", getList)
   router.HandleFunc("/read/task/{uid}/{name}", getTask)

   router.HandleFunc("/update/{uid}", editUser)
   router.HandleFunc("/update/list/{uid}/{list}", editList)
   router.HandleFunc("/update/task/{uid}/{task}", editTask)

	log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
	handleRequests()
}
