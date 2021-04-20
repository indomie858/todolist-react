package main

import (
	"database/request"
	"database/user"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func printResult(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Println("Endpoint Hit: userData")

	// user.User k;
	u, _ := findUser(key)
	jsonUser, _ := json.Marshal(u)
	fmt.Fprintf(w, "%v", string(jsonUser[:]))

	fmt.Println(u)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/userData/{id}", printResult)
	log.Fatal(http.ListenAndServe(":10000", router))
}

func findUser(uid string) (user.User, error) {
	action := "edit"
	item := "list"
	description := map[string]interface{}{
		"name":   "list1",
		"field":  "lock",
		"newval": false,
	}
	r, err := request.DataBaseRequest(uid, action, item, description)
	return *r.User, err

}

func main() {
	handleRequests()
}
