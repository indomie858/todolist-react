# TESTS
navigate to `bin/database`

run `go test -v`

return (so far) -

```
(base) sabra@Sabras-MacBook-Pro database % go test -v
=== RUN   TestCreateUser
--- PASS: TestCreateUser (0.95s)
=== RUN   TestCreateList
--- PASS: TestCreateList (0.83s)
=== RUN   TestCreateListWithPayload
--- PASS: TestCreateListWithPayload (0.86s)
=== RUN   TestCreateTaskWithPaylod
--- PASS: TestCreateTaskWithPaylod (0.51s)
=== RUN   TestCreateSubTask
--- PASS: TestCreateSubTask (0.57s)
=== RUN   TestGetUser
--- PASS: TestGetUser (0.93s)
=== RUN   TestGetList
--- PASS: TestGetList (0.28s)
=== RUN   TestGetLists
--- PASS: TestGetLists (0.17s)
=== RUN   TestGetTask
--- PASS: TestGetTask (0.19s)
=== RUN   TestUpdateTask
--- PASS: TestUpdateTask (0.37s)
=== RUN   TestGetTasks
--- PASS: TestGetTasks (0.41s)
=== RUN   TestDestroyUser
--- PASS: TestDestroyUser (2.39s)
PASS
ok  	database/bin/database	8.479s
```

***I DO NOT GUARANTEE ANY OTHER FUNCTION WILL WORK IF IT DOES NOT HAVE A PASSING TEST***


## CREATE
~~TestCreateUser(t *testing.T)~~

~~TestCreateList(t *testing.T)~~

~~TestCreateListWithPayload(t *testing.T)~~

~~TestCreateTaskWithPaylod(t *testing.T)~~

~~TestCreateSubTask(t *testing.T)~~

## DESTROY
~~TestDestroyUser(t *testing.T)~~

~~TestDestroyList(t *testing.T)~~ *Destroy user destroys lists as well*

~~TestDestroyTask(t *testing.T)~~ *Destroy list (user) destroys tasks as well*

## READ
~~TestGetUser(t *testing.T)~~

~~TestGetList(t *testing.T)~~

~~TestGetLists(t *testing.T)~~

~~TestGetTask(t *testing.T)~~

~~TestGetTasks(t *testing.T)~~

## UPDATE
TestEditUser(t *testing.T)

TestEditList(t *testing.T)

~~TestEditTask(t *testing.T)~~

# API REQUESTS

# USERS
*documentation coming soon*

# LISTS
*full documentatin coming soon*

## AddList(list_name string, fields url.Values) (*TaskJSON, error)
Adds a new list to the List Collection in Firebase, setting any fields that are provided

Possible `fields` are:

|     field     |   type    | required | notes                                                                  |
| :-----------: | :-------: | :------: | ---------------------------------------------------------------------- |
| list_name     | string    |   NO     | Not required in the payload                                            |
| list_owner    | string    |   YES    | Must be given the id of the parent list, or the parent task if subtask |
| lock          | bool      |   NO     | default = false                                                        |
| tasks         | []string  |   NO     | tasks in the list                                                      |
| shared        | bool      |   NO     | default = `false`                                                      |
| shared_users  | []string  |   NO     | default = [""]                                                         |

# TASKS

## AddTask(task_name string, fields url.Values) (*TaskJSON, error)
Adds a new task to the Task Collection in Firebase, setting any fields that are provided

Possible `fields` are:

|     field     |   type    | required | notes                                                                  |
| :-----------: | :-------: | :------: | ---------------------------------------------------------------------- |
| task_name     | string    |   NO     | Not required in the *payload*                                          |
| parent_id     | string    |   NO     | Not required in the *payload*                                          |
| lock          | bool      |   NO     | default = false                                                        |
| list          | string    |   NO     | list name                                                              |
| date_due      | date      |   YES    | Must be given BEFORE end_repeat date, format: `01/02/2006 3:04:05 PM`  |
| done          | bool      |   NO     | Whether or not it's done - default false                               |
| repeating     | bool      |   NO     | default = `false`  - autoset when given repeat                         |
| repeat        | string    |   NO     | default = `never` example: `every week`                                |
| end_repeat    | date      |   NO     | format: `01/02/2006`                                                   |
| discord       | bool      |   NO     | Whether or not discord was selected as a reminder                      |
| email         | bool      |   NO     | Whether or not email was selected as a reminder                        |
| reminder      | string    |   NO     | default = `false`                                                      |
| priority      | string    |   NO     | default = `none`                                                       |
| location      | string    |   NO     | default = ""                                                           |
| description   | string    |   NO     | default = ""                                                           |
| url           | string    |   NO     | default = ""                                                           |
| shared        | bool      |   NO     | default = `false`                                                      |
| shared_users  | []string  |   NO     | default = [""]                                                         |
| sub_task      | bool      |   NO     | default = false                                                        |
| sub_tasks     | []string  |   NO     | default = [""]                                                         |

Fields must be listed exactly as you see them above.

## GetTaskByName(taskname, parentid string) (*TaskJSON, error)
The parent of the task is required to ensure we get the correct task

## GetTaskByID(id string) (*TaskJSON, error)
Called on by DestroyTaskById

## GetTasks(parentid string) ([]*TaskJSON, error)
Returns all tasks that have the provided parentid

## UpdateTask(name, parentid string, fields url.Values) (*TaskJSON, error)
The parent of the task is required to ensure we get the correct task

## DestroyTask(name, parentid string) error
Checks for any subtasks and deletes those as well

The parent of the task is required to ensure we get the correct task

## DestroyTaskById(id string) error
Called on by DestroyTask to remove subtasks

## ParseTaskFields(fields url.Values) map[string]interface{}
Called on by AddTask and UpdateTask.

Parses the url fields into a map that we can send to Firestore

## TaskToJSON() *TaskJSON
Converts the Task structure we use for firestore into a Task structure encoded for JSON
