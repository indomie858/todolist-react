package main

import (
    "database/request"

    "os"
    "fmt"
    "testing"
    "net/http"
    "encoding/json"
    "net/http/httptest"
)

var a App
var testuid string
const (
    FN = "Sabra"
    LN = "Bilodeau"
    MJ = "computer_science"
)
// src: https://github.com/TomFern/go-mux-api/blob/master/main_test.go
func TestMain(m *testing.M) {
    a = App{}
    a.Initialize()

    code := m.Run()
    os.Exit(code)
}

func TestCreateUser(t *testing.T) {
    url := fmt.Sprintf("/api/newUser/%s/lastname/%s", FN, LN)
    req, _ := http.NewRequest("POST", url, nil)
    fmt.Printf("New User Request: %v\n", req)

    response := executeRequest(req)
    checkResponseCode(t, http.StatusOK, response.Code)
    fmt.Printf("New User Response: %v\n",response)

    var m map[string]interface{}
    json.Unmarshal(response.Body.Bytes(), &m)

    user := m["user"]
    testuid = user.Id
    if user.FirstName != FN {
        t.Errorf("Expected the first name to be set to '%s'. Got '%v' instead.", FN, user.FirstName)
    }

    if user.LastName != LN {
        t.Errorf("Expected the last name to be set to '%s'. Got '%v' instead.", LN, user.LastName)
    }
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
    rr := httptest.NewRecorder()
    a.Router.ServeHTTP(rr, req)

    return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
    if expected != actual {
        t.Errorf("Expected response code %d. Got %d\n", expected, actual)
    }
}

func TestGetUser(t *testing.T) {
    url := fmt.Sprintf("/api/readUser/%s", testuid)
    req, _ := http.NewRequest("GET", url, nil)
    fmt.Printf("Get User Request: %v\n", req)

    response := executeRequest(req)
    checkResponseCode(t, http.StatusOK, response.Code)
    fmt.Printf("Get User Response: %v\n",response)

    var m map[string]interface{}
    json.Unmarshal(response.Body.Bytes(), &m)

    user := m["user"]
    if user.FirstName != FN {
        t.Errorf("Expected the first name to be set to '%s'. Got '%v' instead.", FN, user.FirstName)
    }

    if user.LastName != LN {
        t.Errorf("Expected the last name to be set to '%s'. Got '%v' instead.", LN, user.LastName)
    }

    if user.Major != nil {
        t.Errorf("Expected the major not to be set. Got '%v' instead.", user.Major)
    }
}

func TestUpdateUser(t *testing.T) {
    url := fmt.Sprintf("api/updateUser/%s/major/%s", testuid, mj)
    req, _ := http.NewRequest("GET", url, nil)
    fmt.Printf("Update Task Request: %v\n", req)

    response := executeRequest(req)
    checkResponseCode(t, http.StatusOK, response.Code)
    fmt.Printf("Update Task Response: %v\n",response)

    var m map[string]interface{}
    json.Unmarshal(response.Body.Bytes(), &m)

    user := m["user"]
    if user.FirstName != FN {
        t.Errorf("Expected the first name to be set to '%s'. Got '%v' instead.", FN, user.FirstName)
    }

    if user.LastName != LN {
        t.Errorf("Expected the last name to be set to '%s'. Got '%v' instead.", LN, user.LastName)
    }

    if user.Major != MJ {
        t.Errorf("Expected the major to be set to '%s'. Got '%v' instead.", MJ, user.Major)
    }
}

func TestGetUser(t *testing.T) {
    url := fmt.Sprintf("/api/readUser/%s", testuid)
    req, _ := http.NewRequest("GET", url, nil)
    fmt.Printf("Get User Request: %v\n", req)

    response := executeRequest(req)
    checkResponseCode(t, http.StatusOK, response.Code)
    fmt.Printf("Get User Response: %v\n",response)

    var m map[string]interface{}
    json.Unmarshal(response.Body.Bytes(), &m)

    user := m["user"]
    if user.FirstName != FN {
        t.Errorf("Expected the first name to be set to '%s'. Got '%v' instead.", FN, user.FirstName)
    }

    if user.LastName != LN {
        t.Errorf("Expected the last name to be set to '%s'. Got '%v' instead.", LN, user.LastName)
    }

    if user.Major != MJ {
        t.Errorf("Expected the major to be set to '%s'. Got '%v' instead.", MJ, user.Name)
    }
}

func TestDestroyUser(t *testing.T) {
    // delete the user
    url := fmt.Sprintf("/destroy/%s", testuid)
    req, _ := http.NewRequest("DELETE", url, nil)
    fmt.Printf("Destroy User Request: %v\n", req)

    response := executeRequest(req)
    checkResponseCode(t, http.StatusOK, response.Code)
    fmt.Printf("Destroy User Response: %v\n",response)
}
