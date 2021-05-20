# TESTS
navigate to `bin/database`

run `go test -v`

return (so far) -

```
(base) sabra@Sabras-MacBook-Pro database % go test -v
=== RUN   TestCreateUser
--- PASS: TestCreateUser (1.38s)
=== RUN   TestCreateList
--- PASS: TestCreateList (1.06s)
=== RUN   TestCreateListWithPayload
--- PASS: TestCreateListWithPayload (1.18s)
=== RUN   TestCreateTaskWithPaylod
--- PASS: TestCreateTaskWithPaylod (0.54s)
=== RUN   TestCreateSubTask
--- PASS: TestCreateSubTask (0.35s)
=== RUN   TestGetUser
--- PASS: TestGetUser (0.98s)
=== RUN   TestGetList
--- PASS: TestGetList (0.47s)
=== RUN   TestGetLists
--- PASS: TestGetLists (0.81s)
=== RUN   TestGetTask
--- PASS: TestGetTask (0.21s)
=== RUN   TestUpdateTask
--- PASS: TestUpdateTask (0.39s)
=== RUN   TestGetTasks
--- PASS: TestGetTasks (0.37s)
=== RUN   TestDestroyUser
--- PASS: TestDestroyUser (2.50s)
PASS
ok  	database/bin/database	10.261s
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
*full documentation coming soon*

## AddUser(name string, fields url.Values) (*TaskJSON, error)
Adds a new user to the Users Collection in Firebase, setting any fields that are provided

Possible `fields` are:

|        field     |   type    | required | notes                                                                  |
| :--------------: | :-------: | :------: | ---------------------------------------------------------------------- |
| name             | string    |   NO     | Not required in the payload                                            |
| email            | string    |   NO     | Users email that they signed up with                                   |
| status           | string    |   NO     | Users status to be shown to friends *later feature*                    |
| lists            | []string  |   NO     | the ids of the users lists                                             |
| default_list     | string    |   NO     | the default list to add tasks to                                       |
| discord_reminder | bool      |   NO     | Whether or not discord is the default reminder                         |
| email_reminder   | bool      |   NO     | Whether or not email is the default reminder                           |



# LISTS

## AddList(list_name string, fields url.Values) (*TaskJSON, error)
Adds a new list to the Lists Collection in Firebase, setting any fields that are provided

Possible `fields` are:

|     field     |   type    | required | notes                                                                  |
| :-----------: | :-------: | :------: | ---------------------------------------------------------------------- |
| list_name     | string    |   NO     | Not required in the payload                                            |
| list_owner    | string    |   YES    | Must be given the id of the parent list, or the parent task if subtask |
| lock          | bool      |   NO     | default = false                                                        |
| tasks         | []string  |   NO     | Tasks in the list                                                      |
| shared        | bool      |   NO     | default = `false`                                                      |
| shared_users  | []string  |   NO     | default = [""]                                                         |

## GetListByName(listname string) (*ListJSON, error)
Returns a list using the list name. Ensures we get the correct list by specifying the list owner

## GetListByID(lid string) (*ListJSON, error)
Returns a list using the lists id. Checks that the user requesting the list has proper access before returning it.

## GetLists() ([]*ListJSON, error)
Returns all of a users lists

## GetSharedLists() ([]*ListJSON, error)
Returns all lists shared with the requesting user

## UpdateList(id string, fields url.Values) (*ListJSON, error)
Updates the list with the given fields and returns the updated list

## DestroyList(name string) error
Destroys the list and any of its tasks, returning any error that occurred

## DestroyListById(id string) error
Destroys the list and any of its tasks, returning any error that occurred

## ListToJSON() *ListJSON
Parses the list structure into a JSON structure

## UpdateListTasks(listid, id string)
Updates the tasks array in the list

# TASKS

## AddTask(task_name, parentid string, fields url.Values) (*TaskJSON, error)
Adds a new task to the Tasks Collection in Firebase, setting any fields that are provided

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
| sub_tasks     | []string  |   NO     | default = [""]                                                         |

Fields must be listed exactly as you see them above.

## GetTaskByName(taskname, parentid string) (*TaskJSON, error)
The parent of the task is required to ensure we get the correct task

## GetTaskByID(id string) (*TaskJSON, error)
Returns the specified task

## GetTasks(parentid string) ([]*TaskJSON, error)
Returns all tasks that have the provided parentid

## UpdateTask(id string, fields url.Values) (*TaskJSON, error)
The id of the task is required

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

## UpdateTaskSubtasks(taskid, id string) (*TaskJSON, error)
Update the tasks subtasks array
