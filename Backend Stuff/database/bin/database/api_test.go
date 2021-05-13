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

/*
    http://localhost:10000/create/user/sabra
    {"id":"hLVgv6QshC9NBhX9lisl","name":"sabra","lists":["flLMiIuYN2rRjNwweTJi"]}

    http://localhost:10000/read/hLVgv6QshC9NBhX9lisl
    {
        "id": "hLVgv6QshC9NBhX9lisl",
        "name": "sabra",
        "lists": [
            "flLMiIuYN2rRjNwweTJi"
        ]
    }

    TESTS TO WRITE:\
    CREATE
    ~~TestCreateUser(t *testing.T)~~
    ~~TestCreateList(t *testing.T)~~
    ~~TestCreateListWithPayload(t *testing.T)~~
    ~~TestCreateTaskWithPaylod(t *testing.T)~~
    ~~TestCreateSubTask(t *testing.T)~~

    DESTROY
    ~~TestDestroyUser(t *testing.T)~~
    TestDestroyList(t *testing.T)
    TestDestroyTask(t *testing.T)

    READ
    ~~TestGetUser(t *testing.T)~~
    ~~TestGetList(t *testing.T)~~
    TestGetLists(t *testing.T)
    TestGetTask(t *testing.T)
    TestGetTasks(t *testing.T)

    UPDATE
    TestEditUser(t *testing.T)
    TestEditList(t *testing.T)
    TestEditTask(t *testing.T)
*/

var a App
var testuid, testlid1, testlid2, testtid, teststid string

// src: https://github.com/TomFern/go-mux-api/blob/master/main_test.go
func TestMain(m *testing.M) {
    a = App{}
    a.Initialize()

    code := m.Run()
    os.Exit(code)
}

func TestCreateUser(t *testing.T) {
    req, _ := http.NewRequest("POST", "/create/user/testing_user_1", nil)
    //fmt.Printf("req: %v", req)

    response := executeRequest(req)
    checkResponseCode(t, http.StatusOK, response.Code)

    //fmt.Printf("%v",response)
    var m map[string]request.UserJSON
    json.Unmarshal(response.Body.Bytes(), &m)

    user := m["result"]
    testuid = user.Id
    if user.Name != "testing_user_1" {
        t.Errorf("Expected the name to be set to 'testing_user_1'. Got '%v' instead.", user.Name)
    }
}

func TestCreateList(t *testing.T) {
    url := fmt.Sprintf("/create/%s/list/test_list_1", testuid)
    req, _ := http.NewRequest("POST", url, nil)

    response := executeRequest(req)
    checkResponseCode(t, http.StatusOK, response.Code)

    var m map[string]request.ListJSON
    json.Unmarshal(response.Body.Bytes(), &m)

    list := m["result"]
    testlid1 = list.Id
    if list.Name != "test_list_1" {
        t.Errorf("Expected the name to be set to 'test_list_1'. Got '%v' instead.", list.Name)
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

func TestCreateListWithPayload(t *testing.T) {
    url := fmt.Sprintf("/create/%s/list/test_list_2?lock=false&shared=false", testuid)
    req, _ := http.NewRequest("POST", url, nil)

    response := executeRequest(req)
    checkResponseCode(t, http.StatusOK, response.Code)

    var m map[string]request.ListJSON
    json.Unmarshal(response.Body.Bytes(), &m)

    list := m["result"]
    testlid2 = list.Id
    if list.Name != "test_list_2" {
        t.Errorf("Expected the name to be set to 'test_list_2'. Got '%v' instead.", list.Name)
    }

    if list.Lock {
        t.Errorf("Expected lock field to be 'false'. Got '%v' instead.", list.Lock)
    }

    if list.Shared {
        t.Errorf("Expected shared field to be 'false'. Got '%v' instead.", list.Shared)
    }
}

func TestCreateTaskWithPaylod(t *testing.T) {
    url := fmt.Sprintf("/create/%s/task/test_task_1?parent_id=%s&sub_task=false&lock=false", testuid, testlid1)
    req, _ := http.NewRequest("POST", url, nil)

    response := executeRequest(req)
    checkResponseCode(t, http.StatusOK, response.Code)

    var m map[string]request.TaskJSON
    json.Unmarshal(response.Body.Bytes(), &m)

    task := m["result"]
    testtid = task.Id
    if task.Name != "test_task_1" {
        t.Errorf("Expected the name to be set to 'test_task_1'. Got '%v' instead.", task.Name)
    }

    if task.Parent != testlid1 {
        t.Errorf("Expected parent field to be set to '%s'. Got '%v' instead.", testlid1, task.Parent)
    }
    if task.Lock {
        t.Errorf("Expected lock field to be set to 'false'. Got '%v' instead.", task.Lock)
    }
    if task.Subtask {
        t.Errorf("Expected subtask field to be set to 'false'. Got '%v' instead.", task.Subtask)
    }
}

func TestCreateSubTask(t *testing.T) {
    url := fmt.Sprintf("/create/%s/task/sub_task_1?parent_id=%s&sub_task=true", testuid, testtid)
    req, _ := http.NewRequest("POST", url, nil)

    response := executeRequest(req)
    checkResponseCode(t, http.StatusOK, response.Code)

    var m map[string]request.TaskJSON
    json.Unmarshal(response.Body.Bytes(), &m)

    task := m["result"]
    teststid = task.Id
    if task.Name != "sub_task_1" {
        t.Errorf("Expected the name to be set to 'sub_task_1'. Got '%v' instead.", task.Name)
    }

    if task.Parent != testtid {
        t.Errorf("Expected parent field to be set to '%s'. Got '%v' instead.", testtid, task.Parent)
    }

    if !task.Subtask {
        t.Errorf("Expected subtask field to be set to 'true'. Got '%v' instead.", task.Subtask)
    }
}

func TestGetUser(t *testing.T) {
    url := fmt.Sprintf("/read/%s", testuid)
    req, _ := http.NewRequest("GET", url, nil)

    response := executeRequest(req)
    checkResponseCode(t, http.StatusOK, response.Code)

    var m map[string]request.UserJSON
    json.Unmarshal(response.Body.Bytes(), &m)

    user := m["result"]
    if user.Name != "testing_user_1" {
        t.Errorf("Expected the name to be set to 'testing_user_1'. Got '%v' instead.", user.Name)
    }
}

func TestGetList(t *testing.T) {
    url := fmt.Sprintf("/read/%s/list/test_list_1", testuid)
    req, _ := http.NewRequest("GET", url, nil)

    response := executeRequest(req)
    checkResponseCode(t, http.StatusOK, response.Code)

    var m map[string]request.ListJSON
    json.Unmarshal(response.Body.Bytes(), &m)

    list := m["result"]
    if list.Name != "test_list_1" {
        t.Errorf("Expected the name to be set to 'test_list_1'. Got '%v' instead.", list.Name)
    }
}

func TestGetLists(t *testing.T) {
    url := fmt.Sprintf("/read/%s/lists", testuid)
    req, _ := http.NewRequest("GET", url, nil)

    response := executeRequest(req)
    checkResponseCode(t, http.StatusOK, response.Code)

    var m map[string][]request.ListJSON
    json.Unmarshal(response.Body.Bytes(), &m)

    lists := m["result"]
    if len(lists) == 0 {
        t.Errorf("Expected data to be in result. Got '%v' instead.", lists)
    }
}


func TestDestroyUser(t *testing.T) {
    // delete the user
    url := fmt.Sprintf("/destroy/%s", testuid)
    req, _ := http.NewRequest("DELETE", url, nil)

    response := executeRequest(req)
    checkResponseCode(t, http.StatusOK, response.Code)

}