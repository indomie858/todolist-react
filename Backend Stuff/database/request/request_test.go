package request

import "testing"

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

    TESTS TO WRITE:
    CREATE
    TestCreateUser(t *testing.T)
    TestCreateList(t *testing.T)
    TestCreateListWithPayload(t *testing.T)
    TestCreateTask(t *testing.T)
    TestCreateTaskWithPaylod(t *testing.T)
    TestCreateSubTask(t *testing.T)

    DESTROY
    TestDestroyUser(t *testing.T)
    TestDestroyList(t *testing.T)
    TestDestroyTask(t *testing.T)

    READ
    TestGetUser(t *testing.T)
    TestGetList(t *testing.T)
    TestGetLists(t *testing.T)
    TestGetTask(t *testing.T)
    TestGetTasks(t *testing.T)

    UPDATE
    TestEditUser(t *testing.T)
    TestEditList(t *testing.T)
    TestEditTask(t *testing.T)


func TestGetUser(t *testing.T) {
    req, _ := http.NewRequest("GET", "/read/hLVgv6QshC9NBhX9lisl", nil)
    response := executeRequest(req)
    var originalUser map[string]interface{}
    json.Unmarshal(response.Body.Bytes(), &originalUser)
}
*/
