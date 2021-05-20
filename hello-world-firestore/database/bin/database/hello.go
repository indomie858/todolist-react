package main

import (
	"todolist-react/hello-world-firestore/database/request"

	"fmt"
    "log"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
)

// A struct is Go datatype
//
// App is our structure to hold the API router
// and our request object
type App struct {
    Router *mux.Router
    Request *request.Request
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

// src: https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql
func (a *App) Initialize() {
    a.Router = mux.NewRouter().StrictSlash(true)
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
    fmt.Printf("Running on port %s", addr)
    log.Fatal(http.ListenAndServe(":5000", a.Router))
}

func (a *App) homePage(w http.ResponseWriter, r *http.Request) {
    //fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Printf("Endpoint Hit: homePage")
}

// Create a new user in the Firstore database wih the provided name
//
// Example:
// http://localhost:5000/api/newUser/{firstname}/lastname/{lastname}
// http://localhost:5000/api/newUser/sabra/lastname/bilodeau
//
func (a *App) createUser(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: createUser")

    // Read the variables passed
    vars := mux.Vars(r)
	uid := vars["uid"]
	firstname := vars["firstname"]
    lastname := vars["lastname"]

    // Create a new request for the app
    a.Request = request.NewRequest("create", uid)

    /* The code below is only necessary if we are getting extra information
    // from the payload - things that are after a ? in the URL, which we aren't
    // in this paticular example, but we'll leave it just in case.
    // i.e.
    // http://localhost:5000/api/newUser/sabra/lastname/bilodeau?major=computer_science */

    // Get the payload params and display them to the terminal
	payload := r.URL.Query()

    /*fmt.Println("\nPAYLOAD PARAMATERS")
    for k, v := range payload {
        s := fmt.Sprintf("%v => %v", k, v)
        fmt.Printf("%v\n", s)
    }*/

    // Perform the requested action
    user, err := a.Request.AddUser(firstname, lastname, payload)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    respondWithJSON(w, http.StatusOK, map[string]interface{}{"user": user})
}

// Get a user from the Firstore database with the specified UID
//
// Example :
// http://localhost:5000/api/readUser/{uid}
// http://localhost:5000/api/readUser/a3a1hWUx5geKB8qeR6fbk5LZZGI2
//
func (a *App) getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getUser")

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

    respondWithJSON(w, http.StatusOK, map[string]*request.UserJSON{"user": user})
}

// Update a Firestore user data
//
// Example :
// http://localhost:5000/api/updateUser/{uid}/major/{major}
// http://localhost:5000/api/updateUser/MIUVfleqSkxAtzwNeW0W/major/computer_science
//
func (a *App) updateUser(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: updateUser")

    // Read the variables passed
    vars := mux.Vars(r)
    uid := vars["uid"]
    major := vars["major"]

    /* The code below is only necessary if we are getting extra information
    // from the payload - things that are after a ? in the URL, which we aren't
    // in this paticular example, but we'll leave it just in case.
    // i.e.
    // http://localhost:5000/api/newUser/sabra/lastname/bilodeau?major=computer_science */

    // Get the payload params and display them to the terminal
    /*payload := r.URL.Query()

    fmt.Println("\nPAYLOAD PARAMATERS")
    for k, v := range payload {
        s := fmt.Sprintf("%v => %v", k, v)
        fmt.Printf("%v\n", s)
    }*/

    // Create a new request for the app
    a.Request = request.NewRequest("update", uid)

    // Perform the requested action
    a.Request.GetUser()                         // First we need to get the user
    user, err := a.Request.UpdateUser(major)  // Then we can update the user
    if err != nil {
        respondWithError(w, http.StatusBadRequest,  err.Error())
        return
    }

    respondWithJSON(w, http.StatusOK, map[string]*request.UserJSON{"user": user})
}


// Remove a user from the Firstore database, specified by UID
//
// Example :
// http://localhost:10000/destroy/{uid}
// http://localhost:10000/destroy/MIUVfleqSkxAtzwNeW0W
//
func (a *App) destroyUser(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: destroyUser")

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

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", a.homePage)

    // Create function, has two variables in the route - firstname and lastname
	a.Router.HandleFunc("/api/newUser/{firstname}/lastname/{lastname}", a.createUser).Methods("GET", "POST")

    // Destroy function - has one variable, the user id of the user to destroy
	a.Router.HandleFunc("/api/destroyUser/{uid}", a.destroyUser).Methods("GET","DELETE")

    // Read function - has one variable, the user id of the user to read
    a.Router.HandleFunc("/api/readUser/{uid}", a.getUser).Methods("GET", "POST")

    // Update function - has two variables, the user id and what their major is
    a.Router.HandleFunc("/api/updateUser/{uid}/major/{major}", a.updateUser).Methods("GET", "PUT", "POST")
}

func main() {
    a := App{}
    a.Initialize()
    a.Run(":5000")
}
