# TESTS
navigate to `bin/database`

run `go test -v`

return (so far) -

```
(base) sabra@Sabras-MacBook-Pro database % go test -v
=== RUN   TestCreateUser
--- PASS: TestCreateUser (1.08s)
=== RUN   TestCreateList
--- PASS: TestCreateList (0.85s)
=== RUN   TestCreateListWithPayload
--- PASS: TestCreateListWithPayload (0.87s)
=== RUN   TestCreateTaskWithPaylod
--- PASS: TestCreateTaskWithPaylod (0.53s)
=== RUN   TestCreateSubTask
--- PASS: TestCreateSubTask (0.53s)
=== RUN   TestGetUser
--- PASS: TestGetUser (0.21s)
=== RUN   TestGetList
--- PASS: TestGetList (0.18s)
=== RUN   TestGetLists
--- PASS: TestGetLists (0.19s)
=== RUN   TestDestroyUser
--- PASS: TestDestroyUser (0.27s)
PASS
ok  	database/bin/database	4.730s
```

***DESTROY FUNCTIONS DEF DO NOT WORK PROPERLY YET***

***I DO NOT GUARANTEE ANY OTHER FUNCTION WILL WORK IF IT DOES NOT HAVE A PASSING TEST***


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


# TASKS
*Request Methods*

## AddTask(task_name string, fields url.Values) (*TaskJSON, error)
    TODO: if parent_id field is set, update parent reference array
    if subtask, parent is task, else parent is list

## GetTaskByName(taskname string) (*TaskJSON, error)
    Called on by UpdateTask, Destroy Task

## GetTaskByID(id string) (*TaskJSON, error)
    Called on by UpdateTask

## GetTasks(parentid string) ([]*TaskJSON, error)

## UpdateTask(name string, fields url.Values) (*TaskJSON, error)

## DestroyTask(name string) error
    Checks for any subtasks and deletes those as well

## DestroyTaskById(id string) error
    Called on by DestroyTask to remove subtasks

## ParseTaskFields(fields url.Values) map[string]interface{}
    Called on by **AddTask** and **UpdateTask**

    Parses the url fields into a map that we can send to Firestore

## TaskToJSON() *TaskJSON
