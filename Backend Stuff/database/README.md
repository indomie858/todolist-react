# TOC
- [Getting Started with Database Server](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#getting-started-with-database-server)
- [Datbase API Requests](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#database-api-requests)
   - [Adding Data to a Collection](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#adding-data-to-a-collection)
      - [Add User w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#add-user)
      - [Add List w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#add-list)
      - [Add Task w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#add-task)
      - [Add Subtask](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#add-subtask)
   - [Removing Data From a Collection](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#removing-data-from-a-collection)
      - [Destroy User w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#destroy-user)
      - [Destroy List w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#destroy-list)
      - [Destroy Task w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#destroy-task)
   - [Reading Data From a Collection](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#reading-data-from-a-collection)
      - [Read User w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#read-user)
      - [Read List](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#read-list)
      - [Read Lists w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#read-lists)
      - [Read Task w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#read-task)
      - [Read Tasks w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#read-tasks)
   - [Editing Data](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#editing-values-in-the-database)
      - [Edit User w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#edit-user)
      - [Edit List w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#edit-list)
      - [Edit Task](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#edit-task)


# Getting Started with Database Server

## The GOPATH environment variable
[snippet from here](https://golang.org/doc/gopath_code#GOPATH)

The GOPATH environment variable specifies the location of your workspace. It defaults to a directory named go inside your home directory, so $HOME/go on Unix, $home/go on Plan 9, and %USERPROFILE%\go (usually C:\Users\YourName\go) on Windows.

_If you would like to work in a different location, you will need to set GOPATH to the path to that directory._ (Another common setup is to set GOPATH=$HOME). Note that GOPATH must not be the same path as your Go installation.

The command `go env GOPATH` prints the effective current GOPATH; it prints the default location if the environment variable is unset.

For convenience, add the workspace's bin subdirectory to your PATH:

`$ export PATH=$PATH:$(go env GOPATH)/bin`

To learn more about the GOPATH environment variable, see 'go help gopath'.

## .env File
I believe you need to make a .env file in `database/bin` ... I don't know how to make it so it just auto works like the server *r i p*

File should contain
```env
   # .env file
   # Configuration for Firestore SDK

   TYPE=service_account
   PROJECT_ID=friday-584
   PRIVATE_KEY_ID=
   PRIVATE_KEY=
   CLIENT_EMAIL=
   CLIENT_ID=
   AUTH_URI=
   TOKEN_URI=
   AUTH_PROVIDER=
   CLIENT_X509_CERT_URL=
```

In terminal navigate to `database/bin` and run `go run api_server.go`

# Database API Requests

`http://localhost:10000/`

## Adding Data to a Collection

### Add User
`http://localhost:10000/create/user/{name}`

**Example**

URL

`http://localhost:10000/create/user/sabra`

Return (new user)

```json
{
    "Id": "MIUVfleqSkxAtzwNeW0W",
    "Name": "sabra",
    "Email": "",
    "Status": "",
    "Lists": null,
    "Settings": ""
}
```

In terminal

```bash
 Endpoint Hit: createUser
 New users name: sabra

 PAYLOAD PARAMATERS
```

### Add List
`http://localhost:10000/create/{uid}/list/{name}?<params>`

**Example**

URL

`http://localhost:10000/create/a3a1hWUx5geKB8qeR6fbk5LZZGI2/list/test_add_list`

`http://localhost:10000/create/a3a1hWUx5geKB8qeR6fbk5LZZGI2/list/test_add_list?lock=false`

Return (updated list)

```json
{
   "Id": "dcWbqvKvU3fUYzcCumbb",
   "Name": "test_add_list",
   "Owner": "a3a1hWUx5geKB8qeR6fbk5LZZGI2",
   "Lock": false,
   "Shared": false,
   "SharedUsers": null,
   "Tasks": null
}
```

In terminal
```bash
 Endpoint Hit: createList
 list_name: test_add_list

 PAYLOAD PARAMATERS
 lock => [false]
 &{dcWbqvKvU3fUYzcCumbb test_add_list a3a1hWUx5geKB8qeR6fbk5LZZGI2 false false [] []}
```

### Add Task
`http://localhost:10000/create/{uid}/task/{name}?<params>`

Will eventually be changed to so we can add the task to the list and the user at the same time
`http://localhost:10000/create/{uid}/list/{name}/task/{name}`

**Example**

URL
`http://localhost:10000/create/a3a1hWUx5geKB8qeR6fbk5LZZGI2/task/test_task_1`


Return
```json
{
    "Id": "ykyMNNOAU9RWF2NBgghQ",
    "Name": "test_task_1",
    "Owner": "a3a1hWUx5geKB8qeR6fbk5LZZGI2",
    "Parent": "",
    "Lock": false,
    "DueDate": "0001-01-01T00: 00: 00Z",
    "IdealStart": "0001-01-01T00: 00: 00Z",
    "StartDate": "0001-01-01T00: 00: 00Z",
    "Repeating": false,
    "Repeat": "",
    "Remind": false,
    "Reminder": "",
    "TimeFrame": 0,
    "Location": "",
    "Description": "",
    "Url": "",
    "Subtasks": [
        ""
    ]
}
```

Eventual URL to test ..

`http://localhost:10000/create/a3a1hWUx5geKB8qeR6fbk5LZZGI2/list/test_add_list/task/test_task_1`


### Add Subtask
`http://localhost:10000/create/{uid}/subtask/{name}`

## Removing Data From a Collection
### Destroy User
`http://localhost:10000/destroy/{uid}`

**Example**

URL

`http://localhost:10000/destroy/MIUVfleqSkxAtzwNeW0W`

Return
`user successfully deleted`

In terminal
```bash
   Endpoint Hit: destroyUser
   user successfully deleted
```

### Destroy List
`http://localhost:10000/destroy/{uid}/list/{name}`

**Example**

URL

`http://localhost:10000/destroy/MIUVfleqSkxAtzwNeW0W/list/first_list`

Return

`list successfully deleted`

### Destroy Task
`http://localhost:10000/destroy/{uid}/task/{name}`

**Example**

URL

`http://localhost:10000/destroy/a3a1hWUx5geKB8qeR6fbk5LZZGI2/task/test_task_1`

Return

`task successfully deleted`

## Reading Data From a Collection
### Read User
`http://localhost:10000/read/{uid}`

**Example**

URL

`http://localhost:10000/read/a3a1hWUx5geKB8qeR6fbk5LZZGI2`

Return (user)

```json
{
    "Id": "a3a1hWUx5geKB8qeR6fbk5LZZGI2",
    "Name": "max",
    "Email": "",
    "Status": "",
    "Lists": [
        "NIcoux7atd3A8Lv7guUO",
        "ahsdfhhf"
    ],
    "Settings": ""
}
```

In terminal

```bash
 Endpoint Hit: getUser
 &{a3a1hWUx5geKB8qeR6fbk5LZZGI2 max   [NIcoux7atd3A8Lv7guUO ahsdfhhf] }
```

### Read List
`http://localhost:10000/read/{uid}/list/{name}`

### Read List*s*
( emphasis on the s )

`http://localhost:10000/read/{uid}/lists`

**Example**

URL

`http://localhost:10000/read/a3a1hWUx5geKB8qeR6fbk5LZZGI2/lists`

Return (all users lists)
```json
[
    {
        "Id": "NIcoux7atd3A8Lv7guUO",
        "Name": "list1",
        "Owner": "a3a1hWUx5geKB8qeR6fbk5LZZGI2",
        "Lock": false,
        "Shared": false,
        "SharedUsers": null,
        "Tasks": null
    },
    {
        "Id": "WBZgf5dM2YEi2V6aDspd",
        "Name": "list2",
        "Owner": "a3a1hWUx5geKB8qeR6fbk5LZZGI2",
        "Lock": false,
        "Shared": false,
        "SharedUsers": null,
        "Tasks": null
    },
    {
        "Id": "dcWbqvKvU3fUYzcCumbb",
        "Name": "test_add_list",
        "Owner": "a3a1hWUx5geKB8qeR6fbk5LZZGI2",
        "Lock": false,
        "Shared": false,
        "SharedUsers": null,
        "Tasks": null
    },
    {
        "Id": "nlk0zDpnuatUU7bIHhms",
        "Name": "test_add_list",
        "Owner": "a3a1hWUx5geKB8qeR6fbk5LZZGI2",
        "Lock": false,
        "Shared": false,
        "SharedUsers": null,
        "Tasks": null
    }
]

```

### Read Task
`http://localhost:10000/read/{uid}/task/{name}`

**Example**

URL

`http://localhost:10000/read/a3a1hWUx5geKB8qeR6fbk5LZZGI2/task/task1`

Return

```json
{
    "Id": "RfQWRaILs6LhFg2PgUJq",
    "Name": "task1",
    "Owner": "a3a1hWUx5geKB8qeR6fbk5LZZGI2",
    "Parent": "NIcoux7atd3A8Lv7guUO",
    "Lock": false,
    "DueDate": "0001-01-01T00:00:00Z",
    "IdealStart": "0001-01-01T00:00:00Z",
    "StartDate": "0001-01-01T00:00:00Z",
    "Repeating": false,
    "Repeat": "",
    "Remind": false,
    "Reminder": "",
    "TimeFrame": 0,
    "Location": "",
    "Description": "",
    "Url": "",
    "Subtasks": null
}
```

### Read Task*s*
( emphasis on the s )

`http://localhost:10000/read/{uid}/tasks/{parent_id}`

**Example**

URL

`http://localhost:10000/read/a3a1hWUx5geKB8qeR6fbk5LZZGI2/tasks/NIcoux7atd3A8Lv7guUO`

Return (all tasks in the list)

```json
[
    {
        "Id": "RfQWRaILs6LhFg2PgUJq",
        "Name": "task1",
        "Owner": "a3a1hWUx5geKB8qeR6fbk5LZZGI2",
        "Parent": "NIcoux7atd3A8Lv7guUO",
        "Lock": false,
        "DueDate": "0001-01-01T00:00:00Z",
        "IdealStart": "0001-01-01T00:00:00Z",
        "StartDate": "0001-01-01T00:00:00Z",
        "Repeating": false,
        "Repeat": "",
        "Remind": false,
        "Reminder": "",
        "TimeFrame": 0,
        "Location": "",
        "Description": "",
        "Url": "",
        "Subtasks": null
    }
]
```

## Editing values in the database
### Edit User
`http://localhost:10000/update/{uid}?<params>`

**Example**

Add list to list array

URL

`http://localhost:10000/update/MIUVfleqSkxAtzwNeW0W?lists=qqEkD06oFudIRrCVPAc5`

Return (updated user)

```json
{
    "Id": "MIUVfleqSkxAtzwNeW0W",
    "Name": "sabra",
    "Email": "",
    "Status": "",
    "Lists": [
        "qqEkD06oFudIRrCVPAc5"
    ],
    "Settings": ""
}
```

In terminal

```bash
Endpoint Hit: updateUser

PAYLOAD PARAMATERS
lists => [[33TPlBCXI1DhXksyWtdm]]
&{g1tAZTJgfVOqTtDpvAAz    [[33TPlBCXI1DhXksyWtdm]] }s
```

### Edit List
`http://localhost:10000/update/{uid}/list/{list}?<params>`

**Example**

URL

`http://localhost:10000/update/a3a1hWUx5geKB8qeR6fbk5LZZGI2/list/list1?list_name=list1updated&lock=false`

Return (updated list)

```json
{
    "Id": "NIcoux7atd3A8Lv7guUO",
    "Name": "list1updated",
    "Owner": "a3a1hWUx5geKB8qeR6fbk5LZZGI2",
    "Lock": false,
    "Shared": false,
    "SharedUsers": null,
    "Tasks": null
}
```

In terminal

```bash
 Endpoint Hit: updateList
 listname: list1

 PAYLOAD PARAMATERS
 list_name => [list1updated]
 lock => [false]
 &{NIcoux7atd3A8Lv7guUO list1updated a3a1hWUx5geKB8qeR6fbk5LZZGI2 false false [] []}
```

> REMINDER: once you go to that URL once, it won't work again unless u edit the list name :)

URL

`http://localhost:10000/update/a3a1hWUx5geKB8qeR6fbk5LZZGI2/list/list1updated?list_name=list1`

Return (updated list)

```json
{
    "Id": "NIcoux7atd3A8Lv7guUO",
    "Name": "list1",
    "Owner": "a3a1hWUx5geKB8qeR6fbk5LZZGI2",
    "Lock": false,
    "Shared": false,
    "SharedUsers": null,
    "Tasks": null
}
```

In terminal

```bash
Endpoint Hit: updateList
listname: list1updated

PAYLOAD PARAMATERS
list_name => [list1]
&{NIcoux7atd3A8Lv7guUO list1 a3a1hWUx5geKB8qeR6fbk5LZZGI2 false false [] []}
```

### Edit Task
`http://localhost:10000/update/{uid}/task/{task}?<params>`
