package main

import (

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
    //fmt.Printf("Create User Request: %v\n", req)

    response := executeRequest(req)
    checkResponseCode(t, http.StatusOK, response.Code)
    fmt.Printf("Create User Response: %v\n",response)

    var m map[string]*Result
    json.Unmarshal(response.Body.Bytes(), &m)

    result := m["result"]
    user := result.User
    testuid = user.Id
    if user.Name != "testing_user_1" {
        t.Errorf("Expected the name to be set to 'testing_user_1'. Got '%v' instead.", user.Name)
    }
}

func TestCreateList(t *testing.T) {
    url := fmt.Sprintf("/create/%s/list/test_list_1", testuid)
    req, _ := http.NewRequest("POST", url, nil)
    //fmt.Printf("Create List Request: %v\n", req)

    response := executeRequest(req)
    checkResponseCode(t, http.StatusOK, response.Code)
    fmt.Printf("Create List Response: %v\n",response)

    var m map[string]*Result
    json.Unmarshal(response.Body.Bytes(), &m)

    result := m["result"]

    list := result.List
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
    url := fmt.Sprintf("/create/%s/list/test_list_2?lock=true&shared=false", testuid)
    req, _ := http.NewRequest("POST", url, nil)
    //fmt.Printf("Create List with Payload Request: %v\n", req)

    response := executeRequest(req)
    checkResponseCode(t, http.StatusOK, response.Code)
    fmt.Printf("Create List with Payload Response: %v\n",response)

    var m map[string]Result
    json.Unmarshal(response.Body.Bytes(), &m)

    result := m["result"]
    list := result.List
    testlid2 = list.Id
    if list.Name != "test_list_2" {
        t.Errorf("Expected the name to be set to 'test_list_2'. Got '%v' instead.", list.Name)
    }

    if !list.Lock {
        t.Errorf("Expected lock field to be 'true'. Got '%v' instead.", list.Lock)
    }

    if list.Shared {
        t.Errorf("Expected shared field to be 'false'. Got '%v' instead.", list.Shared)
    }
}

func TestCreateTaskWithPaylod(t *testing.T) {
    date_due := "01/02/2006 3:04:05 PM"
    url := fmt.Sprintf("/create/%s/task/test_task_1/parent/%s?sub_task=false&lock=false&date_due=%s", testuid, testlid1, date_due)
    req, _ := http.NewRequest("POST", url, nil)
    //fmt.Printf("Create Task with Payload Request: %v\n", req)

    response := executeRequest(req)
    checkResponseCode(t, http.StatusOK, response.Code)
    fmt.Printf("Create Task with Payload Response: %v\n",response)

    var m map[string]Result
    json.Unmarshal(response.Body.Bytes(), &m)

    result := m["result"]
    task := result.Task
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
}

func TestCreateSubTask(t *testing.T) {
    url := fmt.Sprintf("/create/%s/subtask/sub_task_1/parent/%s", testuid, testtid)
    req, _ := http.NewRequest("POST", url, nil)
    //fmt.Printf("Create Subtask Request: %v\n", req)

    response := executeRequest(req)
    checkResponseCode(t, http.StatusOK, response.Code)
    fmt.Printf("Create Subtask Response: %v\n",response)

    var m map[string]Result
    json.Unmarshal(response.Body.Bytes(), &m)

    result := m["result"]
    task := result.Task
    teststid = task.Id
    if task.Subtasks[0] != "sub_task_1" {
        t.Errorf("Expected the name to be set to 'sub_task_1'. Got '%v' instead.", task.Name)
    }
}

func TestGetUser(t *testing.T) {
    url := fmt.Sprintf("/read/%s", testuid)
    req, _ := http.NewRequest("GET", url, nil)
    //fmt.Printf("Get User Request: %v\n", req)

    response := executeRequest(req)
    checkResponseCode(t, http.StatusOK, response.Code)
    fmt.Printf("Get User Response: %v\n",response)

    var m map[string]Result
    json.Unmarshal(response.Body.Bytes(), &m)

    result := m["result"]
    user := result.User
    if user.Name != "testing_user_1" {
        t.Errorf("Expected the name to be set to 'testing_user_1'. Got '%v' instead.", user.Name)
    }
}

func TestGetList(t *testing.T) {
    url := fmt.Sprintf("/read/%s/list/%s", testuid, testlid1)
    req, _ := http.NewRequest("GET", url, nil)
    //fmt.Printf("Get List Request: %v\n", req)

    response := executeRequest(req)
    checkResponseCode(t, http.StatusOK, response.Code)
    //fmt.Printf("Get List Response: %v\n",response)

    var m map[string]Result
    json.Unmarshal(response.Body.Bytes(), &m)

    result := m["result"]
    list := result.List
    if list.Name != "test_list_1" {
        t.Errorf("Expected the name to be set to 'test_list_1'. Got '%v' instead.", list.Name)
    }
}

func TestGetLists(t *testing.T) {
    url := fmt.Sprintf("/read/%s/lists", testuid)
    req, _ := http.NewRequest("GET", url, nil)
    //fmt.Printf("Get Lists Request: %v\n", req)

    response := executeRequest(req)
    checkResponseCode(t, http.StatusOK, response.Code)
    fmt.Printf("Get Lists Response: %v\n",response)

    var m map[string]Result
    json.Unmarshal(response.Body.Bytes(), &m)

    result := m["result"]
    lists := result.Lists
    if len(lists) == 0 {
        t.Errorf("Expected data to be in result. Got '%v' instead.", lists)
    }
}

func TestGetTask(t *testing.T) {
    url := fmt.Sprintf("/read/%s/task/%s", testuid, testtid)
    req, _ := http.NewRequest("GET", url, nil)
    //fmt.Printf("Get Task Request: %v\n", req)

    response := executeRequest(req)
    checkResponseCode(t, http.StatusOK, response.Code)
    fmt.Printf("Get Task Response: %v\n",response)

    var m map[string]Result
    json.Unmarshal(response.Body.Bytes(), &m)

    result := m["result"]
    task := result.Task
    if task.Name != "test_task_1" {
        t.Errorf("Expected the name to be set to 'test_task_1'. Got '%v' instead.", task.Name)
    }

    if task.Parent != testlid1 {
        t.Errorf("Expected the parent_id to be set to '%s'. Got '%v' instead.", testlid1, task.Parent)
    }
}

func TestUpdateTask(t *testing.T) {
    url := fmt.Sprintf("/update/%s/task/%s?done=true&discord=true", testuid, testtid)
    req, _ := http.NewRequest("GET", url, nil)
    //fmt.Printf("Update Task Request: %v\n", req)

    response := executeRequest(req)
    checkResponseCode(t, http.StatusOK, response.Code)
    fmt.Printf("Update Task Response: %v\n",response)

    var m map[string]Result
    json.Unmarshal(response.Body.Bytes(), &m)

    result := m["result"]
    task := result.Task
    if task.Name != "test_task_1" {
        t.Errorf("Expected the name to be set to 'test_task_1'. Got '%v' instead.", task.Name)
    }

    if task.Parent != testlid1 {
        t.Errorf("Expected the parent_id to be set to '%s'. Got '%v' instead.", testlid1, task.Parent)
    }

    if !task.Done {
        t.Errorf("Expected done to be set to true. Got '%v' instead.", task.Done)
    }
}

func TestUpdateUser(t *testing.T) {
    url := fmt.Sprintf("/update/%s?discord_reminder=true", testuid)

    req, _ := http.NewRequest("GET", url, nil)
    fmt.Printf("Update Task Request: %v\n", req)

    response := executeRequest(req)
    checkResponseCode(t, http.StatusOK, response.Code)
    fmt.Printf("Update Task Response: %v\n",response)

    var m map[string]Result
    json.Unmarshal(response.Body.Bytes(), &m)

    result := m["result"]
    user := result.User
    if !user.DiscordReminder {
        t.Errorf("Expected discord_reminder to be set to true. Got '%v' instead.", user.DiscordReminder)
    }
}

func TestGetTasks(t *testing.T) {
    url := fmt.Sprintf("/read/%s/tasks/%s", testuid, testlid1)
    req, _ := http.NewRequest("GET", url, nil)
    //fmt.Printf("Get Task Request: %v\n", req)

    response := executeRequest(req)
    checkResponseCode(t, http.StatusOK, response.Code)
    fmt.Printf("Get Tasks Response: %v\n",response)

    var m map[string]Result
    json.Unmarshal(response.Body.Bytes(), &m)

    result := m["result"]
    tasks := result.Tasks
    if len(tasks) == 0 {
        t.Errorf("Expected data to be in result. Got '%v' instead.", tasks)
    }
}

func TestDestroyUser(t *testing.T) {
    // delete the user
    url := fmt.Sprintf("/destroy/%s", testuid)
    req, _ := http.NewRequest("DELETE", url, nil)
    //fmt.Printf("Destroy User Request: %v\n", req)

    response := executeRequest(req)
    checkResponseCode(t, http.StatusOK, response.Code)
    //fmt.Printf("Destroy User Response: %v\n",response)
}
