TOC
- [Getting Started with Database Server](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#getting-started-with-database-server)
- [Datbase API Requests](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#database-api-requests)
   - [Adding Data to a Collection](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#adding-data-to-a-collection)
      - [Add User w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#add-user)
      - [Add List w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#add-list)
      - [Add Task](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#add-task)
      - [Add Subtask](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#add-subtask)
   - [Removing Data](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#removing-data-from-a-collection)
      - [Destroy User](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#destroy-user)
      - [Destroy List](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#destroy-list)
      - [Destroy Task](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#destroy-task)
   - [Reading Data](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#reading-data-from-a-collection)
      - [Read User w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#read-user)
      - [Read List](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#read-list)
      - [Read Task](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#read-task)
   - [Editing Data](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#editing-values-in-the-database)
      - [Update User](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#edit-user)
      - [Update List w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#edit-list)
      - [Update Task](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#edit-task)


# Getting Started with Database Server

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

## Adding data to a Collection

### Add User
`http://localhost:10000/create/user/{name}`

**Example**

URL

`http://localhost:10000/create/user/sabra`

Return
(updated list)
```json
{
    "Id": "hpeOH5GDaYRelq51m4XP",
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
 New user's name: sabra

 PAYLOAD PARAMATERS
```

### Add List
`http://localhost:10000/create/{uid}/list/{name}?<params>`

**Example**

URL

`http://localhost:10000/create/a3a1hWUx5geKB8qeR6fbk5LZZGI2/list/testaddlist`

`http://localhost:10000/create/a3a1hWUx5geKB8qeR6fbk5LZZGI2/list/test_add_list?lock=false`

Return 
(updated list)

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
`http://localhost:10000/create/{uid}/task/{name}`

### Add Subtask
`http://localhost:10000/create/{uid}/subtask/{name}`

## Removing data from a Collection
### Destroy User
`http://localhost:10000/destroy/{uid}`

### Destory List
`http://localhost:10000/destroy/list/{lists}`

### Destroy Task 
`http://localhost:10000/destroy/task/{tasks}`

## Reading data from a Collection
**Read User**: http://localhost:10000/read/{uid}

**Example**

URL

`http://localhost:10000/read/a3a1hWUx5geKB8qeR6fbk5LZZGI2`

Return
(updated list)

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
`http://localhost:10000/read/{uid}/task/{name}`

### Read Task
`http://localhost:10000/read/{uid}/list/{name}`

## Editing values in the database
### Edit User
`http://localhost:10000/update/{uid}?<params>`

### Edit List
`http://localhost:10000/update/{uid}/list/{list}?<params>`

**Example**

URL

`http://localhost:10000/update/a3a1hWUx5geKB8qeR6fbk5LZZGI2/list/list1?list_name=list1updated&lock=false`

Return
(updated list)

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

Return
(updated list)

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
