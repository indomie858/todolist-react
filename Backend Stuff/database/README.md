# TOC
- [Getting Started with Database Server](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#getting-started-with-database-server)
- [Datbase API Requests](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#database-api-requests)
   - [Adding Data to a Collection](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#adding-data-to-a-collection)
      - [Add User w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#add-user)
      - [Add List w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#add-list)
      - [Add Task w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#add-task)
      - [Add Subtask w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#add-subtask)
   - [Removing Data From a Collection](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#removing-data-from-a-collection)
      - [Destroy User w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#destroy-user)
      - [Destroy List w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#destroy-list)
      - [Destroy Task w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#destroy-task)
   - [Reading Data From a Collection](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#reading-data-from-a-collection)
      - [Read User w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#read-user)
      - [Read List w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#read-list)
      - [Read Lists w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#read-lists)
      - [Read Task w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#read-task)
      - [Read Tasks w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#read-tasks)
   - [Editing Data](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#editing-values-in-the-database)
      - [Edit User w/ ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#edit-user)
      - [Edit List](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#edit-list)
      - [Edit Task w/ Ex!](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#edit-task)
- [Database Testing](https://github.com/indomie858/todolist-react/tree/dev/Backend%20Stuff/database#tests)


# Getting Started with Database Server

## .env File
I believe you need to make a .env file in `database/bin` ... I don't know how to make it so it just auto works like the server *r i p*

File should contain the same fields from the service account private key, which can be generated from the [project settings page]( https://console.firebase.google.com/u/0/project/friday-584/settings/serviceaccounts/adminsdk)

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

Payload keys must exactly match the fields for each collection type.\
Possible `fields` are

## USER

|        field     |   type    | required | notes                                                                  |
| :--------------: | :-------: | :------: | ---------------------------------------------------------------------- |
| first_name       | string    |   NO     | Not required in the payload                                            |
| last_name        | string    |   NO     | Not required in the payload                                            |
| email            | string    |   NO     | Users email that they signed up with                                   |
| status           | string    |   NO     | Users status to be shown to friends *later feature*                    |
| lists            | []string  |   NO     | the ids of the users lists                                             |
| default_list     | string    |   NO     | the default list to add tasks to                                       |
| discord_reminder | bool      |   NO     | Whether or not discord is the default reminder                         |
| email_reminder   | bool      |   NO     | Whether or not email is the default reminder                           |

## LIST

|     field     |   type    | required | notes                                                                  |
| :-----------: | :-------: | :------: | ---------------------------------------------------------------------- |
| list_name     | string    |   NO     | Not required in the payload                                            |
| list_owner    | string    |   YES    | Must be given the id of the parent list, or the parent task if subtask |
| lock          | bool      |   NO     | default = false                                                        |
| tasks         | []string  |   NO     | Tasks in the list                                                      |
| shared        | bool      |   NO     | default = `false`                                                      |
| shared_users  | []string  |   NO     | default = [""]                                                         |

## TASK

|     field       |   type    | required | notes                                                                  |
| :-----------:   | :-------: | :------: | ---------------------------------------------------------------------- |
| text            | string    |   NO     | Not required in the *payload*                                          |
| parent_id       | string    |   NO     | Not required in the *payload*                                          |
| lock            | bool      |   NO     | default = false                                                        |
| list            | string    |   NO     | list name                                                              |
| date            | date      |   YES    | Must be given BEFORE end_repeat date, format: `01/02/2006 3:04:05 PM`  |
| done            | bool      |   NO     | Whether or not it's done - default false                               |
| willRepeat      | bool      |   NO     | default = `false`  - autoset when given repeat                         |
| repeatFrequency | string    |   NO     | default = `never` example: `every week`                                |
| end_repeat      | date      |   NO     | format: `01/02/2006`                                                   |
| discordSelected | bool      |   NO     | Whether or not discord was selected as a reminder                      |
| emailSelected   | bool      |   NO     | Whether or not email was selected as a reminder                        |
| reminder        | string    |   NO     | default = `false` options: `at time of event`, `days/mins/weeks before`|
| remind          | bool      |   NO     | Whether or not we should remind the user                               |
| reminder_time   | date      |   NO     | What time to remind the user at - auto determined when given reminder  |
| priority        | string    |   NO     | default = `none`;                                                      |
| location        | string    |   NO     | default = ""                                                           |
| description     | string    |   NO     | default = ""                                                           |
| url             | string    |   NO     | default = ""                                                           |
| shared          | bool      |   NO     | default = `false`                                                      |
| shared_users    | []string  |   NO     | default = [""]                                                         |
| sub_tasks       | []string  |   NO     | default = [""]                                                         |


## Adding Data to a Collection

### Add User
`http://localhost:10000/create/user/{uid}?<params>`

**Example**

URL

`http://localhost:10000/create/user/q0nQCfPpjneYyCXAbXkK?first_name=Sabra&last_name=Bilodeau&email=testing.email.352@my.csun.edu`

Return

```json
"result": {
    "User": {
        "id": "q0nQCfPpjneYyCXAbXkK",
        "first_name": "Sabra",
        "last_name": "Bilodeau",
        "email": "testing.email.352@my.csun.edu",
        "lists": [
            "iBYGoFECohgBloBCEgC4",
            "BhG0RwO5vx1DZxNvyPik"
        ],
        "discord_reminder": false,
        "email_reminder": false
    },
    "List": null,
    "Lists": [
        {
            "id": "BhG0RwO5vx1DZxNvyPik",
            "list_name": "Shared",
            "list_owner": "q0nQCfPpjneYyCXAbXkK",
            "shared": true,
            "shared_users": [
                ""
            ],
            "tasks": [
                "Kwb57EsDB68QiXguacEd"
            ]
        },
        {
            "id": "iBYGoFECohgBloBCEgC4",
            "list_name": "Main",
            "list_owner": "q0nQCfPpjneYyCXAbXkK",
            "shared_users": [
                ""
            ],
            "tasks": [
                "gs7rdjwRrBQwGahGjLYB"
            ]
        }
    ],
    "Task": null,
    "Tasks": null,
    "AllTasks": [
        [
            {
                "id": "Kwb57EsDB68QiXguacEd",
                "text": "First Task !",
                "task_owner": "q0nQCfPpjneYyCXAbXkK",
                "parent_id": "BhG0RwO5vx1DZxNvyPik",
                "date": "0001-01-01T00:00:00Z",
                "isComplete": false,
                "willRepeat": false,
                "repeatFrequency": "never",
                "end_repeat": "0001-01-01T00:00:00Z",
                "remind": false,
                "emailSelected": false,
                "discordSelected": false,
                "reminder": "none",
                "reminder_time": "0001-01-01T00:00:00Z",
                "priority": "none",
                "shared": false,
                "subTasks": [
                    ""
                ]
            }
        ],
        [
            {
                "id": "gs7rdjwRrBQwGahGjLYB",
                "text": "First Task !",
                "task_owner": "q0nQCfPpjneYyCXAbXkK",
                "parent_id": "iBYGoFECohgBloBCEgC4",
                "date": "0001-01-01T00:00:00Z",
                "isComplete": false,
                "willRepeat": false,
                "repeatFrequency": "never",
                "end_repeat": "0001-01-01T00:00:00Z",
                "remind": false,
                "emailSelected": false,
                "discordSelected": false,
                "reminder": "none",
                "reminder_time": "0001-01-01T00:00:00Z",
                "priority": "none",
                "shared": false,
                "subTasks": [
                    ""
                ]
            }
        ]
    ]
}
```

### Add List
`http://localhost:10000/create/{uid}/list/{name}?<params>`

**Example**

URL

`http://localhost:10000/create/8MFkaIrLbLjkxpzGMCwH/list/test_list_1`

Return

```json
{
    "result": {
        "User": null,
        "List": {
            "id": "ZhDmSusiguNCIp9A9H8x",
            "list_name": "test_list_1",
            "list_owner": "q0nQCfPpjneYyCXAbXkK",
            "shared_users": [
                ""
            ],
            "tasks": [
                "CwLns7muiQ0u3PKA8OkD"
            ]
        },
        "Lists": null,
        "Task": null,
        "Tasks": [
            {
                "id": "CwLns7muiQ0u3PKA8OkD",
                "text": "First Task !",
                "task_owner": "q0nQCfPpjneYyCXAbXkK",
                "parent_id": "ZhDmSusiguNCIp9A9H8x",
                "date": "0001-01-01T00:00:00Z",
                "isComplete": false,
                "willRepeat": false,
                "repeatFrequency": "never",
                "end_repeat": "0001-01-01T00:00:00Z",
                "remind": false,
                "emailSelected": false,
                "discordSelected": false,
                "reminder": "none",
                "reminder_time": "0001-01-01T00:00:00Z",
                "priority": "none",
                "shared": false,
                "subTasks": [
                    ""
                ]
            }
        ],
        "AllTasks": null
    }
}
```

Create list with payload\
URL

`http://localhost:10000/create/gNMA6TlIOCdB52LPSuL5/list/test_list_2?lock=true&shared=false`

Return

```json
{
    "result": {
        "User": null,
        "List": {
            "id": "s4LFfNmxchYpSops4zlU",
            "list_name": "test_list_2",
            "list_owner": "q0nQCfPpjneYyCXAbXkK",
            "lock": true,
            "shared_users": [
                ""
            ],
            "tasks": [
                "3vg86jNhau3Szgo9EwTc"
            ]
        },
        "Lists": null,
        "Task": null,
        "Tasks": [
            {
                "id": "3vg86jNhau3Szgo9EwTc",
                "text": "First Task !",
                "task_owner": "q0nQCfPpjneYyCXAbXkK",
                "parent_id": "s4LFfNmxchYpSops4zlU",
                "date": "0001-01-01T00:00:00Z",
                "isComplete": false,
                "willRepeat": false,
                "repeatFrequency": "never",
                "end_repeat": "0001-01-01T00:00:00Z",
                "remind": false,
                "emailSelected": false,
                "discordSelected": false,
                "reminder": "none",
                "reminder_time": "0001-01-01T00:00:00Z",
                "priority": "none",
                "shared": false,
                "subTasks": [
                    ""
                ]
            }
        ],
        "AllTasks": null
    }
}
```

### Add Task
`http://localhost:10000/create/{uid}/task/{name}/parent/{pid}?<params>`

**Example**

URL

`http://localhost:10000/create/gNMA6TlIOCdB52LPSuL5/task/test_task_1/parent/fQksVGJzgTUc6FervXa4?sub_task=false&lock=false&date_due=01/02/2006 3:04:05 PM`


Return

```json
{
    "result": {
        "User": null,
        "List": null,
        "Lists": null,
        "Task": {
            "id": "cMYyUwSbLDZ9Vl2WZIGe",
            "text": "test_task_1",
            "task_owner": "q0nQCfPpjneYyCXAbXkK",
            "parent_id": "ZhDmSusiguNCIp9A9H8x",
            "date": "2006-01-02T15:04:05Z",
            "isComplete": false,
            "willRepeat": false,
            "repeatFrequency": "never",
            "end_repeat": "0001-01-01T00:00:00Z",
            "remind": false,
            "emailSelected": false,
            "discordSelected": false,
            "reminder": "none",
            "reminder_time": "0001-01-01T00:00:00Z",
            "priority": "none",
            "shared": false,
            "subTasks": [
                ""
            ]
        },
        "Tasks": null,
        "AllTasks": null
    }
}
```

### Add Subtask
`http://localhost:10000/create/{uid}/subtask/{name}/parent/{pid}`

**Example**

URL

`http://localhost:10000/create/gNMA6TlIOCdB52LPSuL5/subtask/sub_task_1/parent/mOCohcha1i6sInCJPeEp`

Return

```json
{
    "result": {
        "User": null,
        "List": null,
        "Lists": null,
        "Task": {
            "id": "cMYyUwSbLDZ9Vl2WZIGe",
            "text": "test_task_1",
            "task_owner": "q0nQCfPpjneYyCXAbXkK",
            "parent_id": "ZhDmSusiguNCIp9A9H8x",
            "date": "2006-01-02T15:04:05Z",
            "isComplete": false,
            "willRepeat": false,
            "repeatFrequency": "never",
            "end_repeat": "0001-01-01T00:00:00Z",
            "remind": false,
            "emailSelected": false,
            "discordSelected": false,
            "reminder": "none",
            "reminder_time": "0001-01-01T00:00:00Z",
            "priority": "none",
            "shared": false,
            "subTasks": [
                "sub_task_1"
            ]
        },
        "Tasks": null,
        "AllTasks": null
    }
}
```

## Removing Data From a Collection
### Destroy User
`http://localhost:10000/destroy/{uid}`

**Example**

URL

`http://localhost:10000/destroy/8MFkaIrLbLjkxpzGMCwH`

Return

```json
"result": "user successfully deleted"
```

### Destroy List
`http://localhost:10000/destroy/{uid}/list/{id}`

**Example**

URL

`http://localhost:10000/destroy/MIUVfleqSkxAtzwNeW0W/list/364DgExvwpE4lNC7JV59`

Return

```json
"result": "list successfully deleted"
```

### Destroy Task
`http://localhost:10000/destroy/{uid}/task/{id}`

**Example**

URL

`http://localhost:10000/destroy/a3a1hWUx5geKB8qeR6fbk5LZZGI2/task/MKEUu0LxHZtMOd6KfmsB`

Return

```json
"result": "task successfully deleted"
```

## Reading Data From a Collection
### Read User
`http://localhost:10000/read/{uid}`

**Example**

URL

`http://localhost:10000/read/gNMA6TlIOCdB52LPSuL5`

Return
```json
{
    "result": {
        "User": {
            "id": "q0nQCfPpjneYyCXAbXkK",
            "first_name": "Sabra",
            "last_name": "Bilodeau",
            "email": "testing.email.352@my.csun.edu",
            "lists": [
                "iBYGoFECohgBloBCEgC4",
                "BhG0RwO5vx1DZxNvyPik",
                "ZhDmSusiguNCIp9A9H8x",
                "s4LFfNmxchYpSops4zlU"
            ],
            "discord_reminder": false,
            "email_reminder": false
        },
        "List": null,
        "Lists": [
            {
                "id": "BhG0RwO5vx1DZxNvyPik",
                "list_name": "Shared",
                "list_owner": "q0nQCfPpjneYyCXAbXkK",
                "shared": true,
                "shared_users": [
                    ""
                ],
                "tasks": [
                    "Kwb57EsDB68QiXguacEd"
                ]
            },
            {
                "id": "ZhDmSusiguNCIp9A9H8x",
                "list_name": "test_list_1",
                "list_owner": "q0nQCfPpjneYyCXAbXkK",
                "shared_users": [
                    ""
                ],
                "tasks": [
                    "CwLns7muiQ0u3PKA8OkD",
                    "cMYyUwSbLDZ9Vl2WZIGe"
                ]
            },
            {
                "id": "iBYGoFECohgBloBCEgC4",
                "list_name": "Main",
                "list_owner": "q0nQCfPpjneYyCXAbXkK",
                "shared_users": [
                    ""
                ],
                "tasks": [
                    "gs7rdjwRrBQwGahGjLYB"
                ]
            },
            {
                "id": "s4LFfNmxchYpSops4zlU",
                "list_name": "test_list_2",
                "list_owner": "q0nQCfPpjneYyCXAbXkK",
                "lock": true,
                "shared_users": [
                    ""
                ],
                "tasks": [
                    "3vg86jNhau3Szgo9EwTc"
                ]
            }
        ],
        "Task": null,
        "Tasks": null,
        "AllTasks": [
            [
                {
                    "id": "Kwb57EsDB68QiXguacEd",
                    "text": "First Task !",
                    "task_owner": "q0nQCfPpjneYyCXAbXkK",
                    "parent_id": "BhG0RwO5vx1DZxNvyPik",
                    "date": "0001-01-01T00:00:00Z",
                    "isComplete": false,
                    "willRepeat": false,
                    "repeatFrequency": "never",
                    "end_repeat": "0001-01-01T00:00:00Z",
                    "remind": false,
                    "emailSelected": false,
                    "discordSelected": false,
                    "reminder": "none",
                    "reminder_time": "0001-01-01T00:00:00Z",
                    "priority": "none",
                    "shared": false,
                    "subTasks": [
                        ""
                    ]
                }
            ],
            [
                {
                    "id": "CwLns7muiQ0u3PKA8OkD",
                    "text": "First Task !",
                    "task_owner": "q0nQCfPpjneYyCXAbXkK",
                    "parent_id": "ZhDmSusiguNCIp9A9H8x",
                    "date": "0001-01-01T00:00:00Z",
                    "isComplete": false,
                    "willRepeat": false,
                    "repeatFrequency": "never",
                    "end_repeat": "0001-01-01T00:00:00Z",
                    "remind": false,
                    "emailSelected": false,
                    "discordSelected": false,
                    "reminder": "none",
                    "reminder_time": "0001-01-01T00:00:00Z",
                    "priority": "none",
                    "shared": false,
                    "subTasks": [
                        ""
                    ]
                },
                {
                    "id": "cMYyUwSbLDZ9Vl2WZIGe",
                    "text": "test_task_1",
                    "task_owner": "q0nQCfPpjneYyCXAbXkK",
                    "parent_id": "ZhDmSusiguNCIp9A9H8x",
                    "date": "2006-01-02T15:04:05Z",
                    "isComplete": false,
                    "willRepeat": false,
                    "repeatFrequency": "never",
                    "end_repeat": "0001-01-01T00:00:00Z",
                    "remind": false,
                    "emailSelected": false,
                    "discordSelected": false,
                    "reminder": "none",
                    "reminder_time": "0001-01-01T00:00:00Z",
                    "priority": "none",
                    "shared": false,
                    "subTasks": [
                        "sub_task_1"
                    ]
                }
            ],
            [
                {
                    "id": "gs7rdjwRrBQwGahGjLYB",
                    "text": "First Task !",
                    "task_owner": "q0nQCfPpjneYyCXAbXkK",
                    "parent_id": "iBYGoFECohgBloBCEgC4",
                    "date": "0001-01-01T00:00:00Z",
                    "isComplete": false,
                    "willRepeat": false,
                    "repeatFrequency": "never",
                    "end_repeat": "0001-01-01T00:00:00Z",
                    "remind": false,
                    "emailSelected": false,
                    "discordSelected": false,
                    "reminder": "none",
                    "reminder_time": "0001-01-01T00:00:00Z",
                    "priority": "none",
                    "shared": false,
                    "subTasks": [
                        ""
                    ]
                }
            ],
            [
                {
                    "id": "3vg86jNhau3Szgo9EwTc",
                    "text": "First Task !",
                    "task_owner": "q0nQCfPpjneYyCXAbXkK",
                    "parent_id": "s4LFfNmxchYpSops4zlU",
                    "date": "0001-01-01T00:00:00Z",
                    "isComplete": false,
                    "willRepeat": false,
                    "repeatFrequency": "never",
                    "end_repeat": "0001-01-01T00:00:00Z",
                    "remind": false,
                    "emailSelected": false,
                    "discordSelected": false,
                    "reminder": "none",
                    "reminder_time": "0001-01-01T00:00:00Z",
                    "priority": "none",
                    "shared": false,
                    "subTasks": [
                        ""
                    ]
                }
            ]
        ]
    }
}
```

### Read All Users
`http://localhost:10000/readusers`

Return

```json
{
    "users": [
        {
            "id": "MTREdIGdhUhoT5RyhlR7",
            "name": "testing_user_1",
            "lists": [
                "haTzgO8uPAsXCAYLSrDH",
                "ptbhy7C4xFiHmhYtm9EP",
                "NtFDVgrb5JzMRaOKBLAW"
            ],
            "discord_reminder": false,
            "email_reminder": false
        },
        {
            "id": "a3a1hWUx5geKB8qeR6fbk5LZZGI2",
            "name": "max",
            "lists": [
                "NIcoux7atd3A8Lv7guUO",
                "ahsdfhhf"
            ],
            "default_list": "Shared",
            "discord_reminder": false,
            "email_reminder": false
        },
        {
            "id": "f9oXnGYUlUADNIDambFG",
            "name": "testing_user_1",
            "lists": [
                "5rtFkYURIxBeil4NEdoM",
                "hsHYrOZeeAAuIAOSWaLk",
                "GDQ0gcEqqftU2iQdu3Ae"
            ],
            "discord_reminder": false,
            "email_reminder": false
        },
        {
            "id": "kF9VV9rep89BpjMcf1n0",
            "name": "testing_user_1",
            "lists": [
                "u6krZBxi2eeJ1QcG6bWH",
                "MEC7q1TSixOG4EJj1ApA",
                "lSA8K6KXk98UbexyRNRz"
            ],
            "discord_reminder": false,
            "email_reminder": false
        },
        {
            "id": "mxYIIQrKBgZKIhdPKHmh",
            "name": "testing_user_1",
            "lists": [
                "3l5v1uSc96RvBaTfuVjD",
                "H3uam9oYWvyZPR5bzYxr",
                "6R5c97TTTygB3hBFd32M"
            ],
            "discord_reminder": false,
            "email_reminder": false
        }
    ],
    "AllTasks": null
}
```

### Read List
`http://localhost:10000/read/{uid}/list/{id}`

**Example**

URL

`http://localhost:10000/read/gNMA6TlIOCdB52LPSuL5/list/fQksVGJzgTUc6FervXa4`

Return

```json
{
    "result": {
        "User": null,
        "List": {
            "id": "lAbqHRKJ9OH2Dw9G3TVq",
            "list_name": "test_list_2",
            "list_owner": "Fbk67C0uIfQ1Q3EepFBv",
            "lock": true,
            "shared_users": [
                ""
            ],
            "tasks": [
                "mLIyPR7T8VBfBj9Fzu8D"
            ]
        },
        "Lists": null,
        "Task": null,
        "Tasks": [
            {
                "id": "mLIyPR7T8VBfBj9Fzu8D",
                "text": "first_task",
                "task_owner": "Fbk67C0uIfQ1Q3EepFBv",
                "parent_id": "lAbqHRKJ9OH2Dw9G3TVq",
                "date": "0001-01-01T00:00:00Z",
                "isComplete": false,
                "willRepeat": false,
                "repeatFrequency": "never",
                "end_repeat": "0001-01-01T00:00:00Z",
                "remind": false,
                "emailSelected": false,
                "discordSelected": false,
                "reminder": "none",
                "reminder_time": "0001-01-01T00:00:00Z",
                "priority": "none",
                "shared": false,
                "subTasks": [
                    ""
                ]
            }
        ],
        "AllTasks": null
    }
}
```

### Read List*s*
( emphasis on the s )

`http://localhost:10000/read/{uid}/lists`

**Example**

URL
`http://localhost:10000/create/gNMA6TlIOCdB52LPSuL5/list/test_list_2?lock=true&shared=false`

Return

```json
{
    "result": {
        "User": null,
        "List": null,
        "Lists": [
            {
                "id": "BhG0RwO5vx1DZxNvyPik",
                "list_name": "Shared",
                "list_owner": "q0nQCfPpjneYyCXAbXkK",
                "shared": true,
                "shared_users": [
                    ""
                ],
                "tasks": [
                    "Kwb57EsDB68QiXguacEd"
                ]
            },
            {
                "id": "ZhDmSusiguNCIp9A9H8x",
                "list_name": "test_list_1",
                "list_owner": "q0nQCfPpjneYyCXAbXkK",
                "shared_users": [
                    ""
                ],
                "tasks": [
                    "CwLns7muiQ0u3PKA8OkD",
                    "cMYyUwSbLDZ9Vl2WZIGe"
                ]
            },
            {
                "id": "iBYGoFECohgBloBCEgC4",
                "list_name": "Main",
                "list_owner": "q0nQCfPpjneYyCXAbXkK",
                "shared_users": [
                    ""
                ],
                "tasks": [
                    "gs7rdjwRrBQwGahGjLYB"
                ]
            },
            {
                "id": "s4LFfNmxchYpSops4zlU",
                "list_name": "test_list_2",
                "list_owner": "q0nQCfPpjneYyCXAbXkK",
                "lock": true,
                "shared_users": [
                    ""
                ],
                "tasks": [
                    "3vg86jNhau3Szgo9EwTc"
                ]
            }
        ],
        "Task": null,
        "Tasks": null,
        "AllTasks": [
            [
                {
                    "id": "Kwb57EsDB68QiXguacEd",
                    "text": "First Task !",
                    "task_owner": "q0nQCfPpjneYyCXAbXkK",
                    "parent_id": "BhG0RwO5vx1DZxNvyPik",
                    "date": "0001-01-01T00:00:00Z",
                    "isComplete": false,
                    "willRepeat": false,
                    "repeatFrequency": "never",
                    "end_repeat": "0001-01-01T00:00:00Z",
                    "remind": false,
                    "emailSelected": false,
                    "discordSelected": false,
                    "reminder": "none",
                    "reminder_time": "0001-01-01T00:00:00Z",
                    "priority": "none",
                    "shared": false,
                    "subTasks": [
                        ""
                    ]
                }
            ],
            [
                {
                    "id": "CwLns7muiQ0u3PKA8OkD",
                    "text": "First Task !",
                    "task_owner": "q0nQCfPpjneYyCXAbXkK",
                    "parent_id": "ZhDmSusiguNCIp9A9H8x",
                    "date": "0001-01-01T00:00:00Z",
                    "isComplete": false,
                    "willRepeat": false,
                    "repeatFrequency": "never",
                    "end_repeat": "0001-01-01T00:00:00Z",
                    "remind": false,
                    "emailSelected": false,
                    "discordSelected": false,
                    "reminder": "none",
                    "reminder_time": "0001-01-01T00:00:00Z",
                    "priority": "none",
                    "shared": false,
                    "subTasks": [
                        ""
                    ]
                },
                {
                    "id": "cMYyUwSbLDZ9Vl2WZIGe",
                    "text": "test_task_1",
                    "task_owner": "q0nQCfPpjneYyCXAbXkK",
                    "parent_id": "ZhDmSusiguNCIp9A9H8x",
                    "date": "2006-01-02T15:04:05Z",
                    "isComplete": false,
                    "willRepeat": false,
                    "repeatFrequency": "never",
                    "end_repeat": "0001-01-01T00:00:00Z",
                    "remind": false,
                    "emailSelected": false,
                    "discordSelected": false,
                    "reminder": "none",
                    "reminder_time": "0001-01-01T00:00:00Z",
                    "priority": "none",
                    "shared": false,
                    "subTasks": [
                        "sub_task_1"
                    ]
                }
            ],
            [
                {
                    "id": "gs7rdjwRrBQwGahGjLYB",
                    "text": "First Task !",
                    "task_owner": "q0nQCfPpjneYyCXAbXkK",
                    "parent_id": "iBYGoFECohgBloBCEgC4",
                    "date": "0001-01-01T00:00:00Z",
                    "isComplete": false,
                    "willRepeat": false,
                    "repeatFrequency": "never",
                    "end_repeat": "0001-01-01T00:00:00Z",
                    "remind": false,
                    "emailSelected": false,
                    "discordSelected": false,
                    "reminder": "none",
                    "reminder_time": "0001-01-01T00:00:00Z",
                    "priority": "none",
                    "shared": false,
                    "subTasks": [
                        ""
                    ]
                }
            ],
            [
                {
                    "id": "3vg86jNhau3Szgo9EwTc",
                    "text": "First Task !",
                    "task_owner": "q0nQCfPpjneYyCXAbXkK",
                    "parent_id": "s4LFfNmxchYpSops4zlU",
                    "date": "0001-01-01T00:00:00Z",
                    "isComplete": false,
                    "willRepeat": false,
                    "repeatFrequency": "never",
                    "end_repeat": "0001-01-01T00:00:00Z",
                    "remind": false,
                    "emailSelected": false,
                    "discordSelected": false,
                    "reminder": "none",
                    "reminder_time": "0001-01-01T00:00:00Z",
                    "priority": "none",
                    "shared": false,
                    "subTasks": [
                        ""
                    ]
                }
            ]
        ]
    }
}

```

### Read Task
`http://localhost:10000/read/{uid}/task/{id}`

**Example**

URL

`http://localhost:10000/read/gNMA6TlIOCdB52LPSuL5/task/mOCohcha1i6sInCJPeEp`

Return

```json
{
    "result": {
        "User": null,
        "List": null,
        "Lists": null,
        "Task": {
            "id": "cMYyUwSbLDZ9Vl2WZIGe",
            "text": "test_task_1",
            "task_owner": "q0nQCfPpjneYyCXAbXkK",
            "parent_id": "ZhDmSusiguNCIp9A9H8x",
            "date": "2006-01-02T15:04:05Z",
            "isComplete": false,
            "willRepeat": false,
            "repeatFrequency": "never",
            "end_repeat": "0001-01-01T00:00:00Z",
            "remind": false,
            "emailSelected": false,
            "discordSelected": false,
            "reminder": "none",
            "reminder_time": "0001-01-01T00:00:00Z",
            "priority": "none",
            "shared": false,
            "subTasks": [
                "sub_task_1"
            ]
        },
        "Tasks": null,
        "AllTasks": null
    }
}
```

### Read Task*s*
( emphasis on the s )

`http://localhost:10000/read/{uid}/tasks/{pid}`

**Example**

URL

`http://localhost:10000/read/pSdsua0LrFB4IiGIsI3j/tasks/Tju8b4Deg2f5lHvgE4PJ`

Return

```json
{
    "result": {
        "User": null,
        "List": null,
        "Lists": null,
        "Task": null,
        "Tasks": [
            {
                "id": "CwLns7muiQ0u3PKA8OkD",
                "text": "First Task !",
                "task_owner": "q0nQCfPpjneYyCXAbXkK",
                "parent_id": "ZhDmSusiguNCIp9A9H8x",
                "date": "0001-01-01T00:00:00Z",
                "isComplete": false,
                "willRepeat": false,
                "repeatFrequency": "never",
                "end_repeat": "0001-01-01T00:00:00Z",
                "remind": false,
                "emailSelected": false,
                "discordSelected": false,
                "reminder": "none",
                "reminder_time": "0001-01-01T00:00:00Z",
                "priority": "none",
                "shared": false,
                "subTasks": [
                    ""
                ]
            },
            {
                "id": "cMYyUwSbLDZ9Vl2WZIGe",
                "text": "test_task_1",
                "task_owner": "q0nQCfPpjneYyCXAbXkK",
                "parent_id": "ZhDmSusiguNCIp9A9H8x",
                "date": "2006-01-02T15:04:05Z",
                "isComplete": true,
                "willRepeat": false,
                "repeatFrequency": "never",
                "end_repeat": "0001-01-01T00:00:00Z",
                "remind": false,
                "emailSelected": false,
                "discordSelected": false,
                "reminder": "none",
                "reminder_time": "0001-01-01T00:00:00Z",
                "priority": "none",
                "shared": false,
                "subTasks": [
                    "sub_task_1"
                ]
            }
        ],
        "AllTasks": null
    }
}
```

## Editing values in the database
### Edit User
`http://localhost:10000/update/{uid}?<params>`

**Example**

Set discord reminders to true

URL

`http://localhost:10000/update/Fbk67C0uIfQ1Q3EepFBv?discord_reminder=true`

Return

```json
{
    "result": {
        "User": {
            "id": "q0nQCfPpjneYyCXAbXkK",
            "first_name": "Sabra",
            "last_name": "Bilodeau",
            "email": "testing.email.352@my.csun.edu",
            "lists": [
                "iBYGoFECohgBloBCEgC4",
                "BhG0RwO5vx1DZxNvyPik",
                "ZhDmSusiguNCIp9A9H8x",
                "s4LFfNmxchYpSops4zlU"
            ],
            "discord_reminder": true,
            "email_reminder": false
        },
        "List": null,
        "Lists": [
            {
                "id": "BhG0RwO5vx1DZxNvyPik",
                "list_name": "Shared",
                "list_owner": "q0nQCfPpjneYyCXAbXkK",
                "shared": true,
                "shared_users": [
                    ""
                ],
                "tasks": [
                    "Kwb57EsDB68QiXguacEd"
                ]
            },
            {
                "id": "ZhDmSusiguNCIp9A9H8x",
                "list_name": "test_list_1",
                "list_owner": "q0nQCfPpjneYyCXAbXkK",
                "shared_users": [
                    ""
                ],
                "tasks": [
                    "CwLns7muiQ0u3PKA8OkD",
                    "cMYyUwSbLDZ9Vl2WZIGe"
                ]
            },
            {
                "id": "iBYGoFECohgBloBCEgC4",
                "list_name": "Main",
                "list_owner": "q0nQCfPpjneYyCXAbXkK",
                "shared_users": [
                    ""
                ],
                "tasks": [
                    "gs7rdjwRrBQwGahGjLYB"
                ]
            },
            {
                "id": "s4LFfNmxchYpSops4zlU",
                "list_name": "test_list_2",
                "list_owner": "q0nQCfPpjneYyCXAbXkK",
                "lock": true,
                "shared_users": [
                    ""
                ],
                "tasks": [
                    "3vg86jNhau3Szgo9EwTc"
                ]
            }
        ],
        "Task": null,
        "Tasks": null,
        "AllTasks": [
            [
                {
                    "id": "Kwb57EsDB68QiXguacEd",
                    "text": "First Task !",
                    "task_owner": "q0nQCfPpjneYyCXAbXkK",
                    "parent_id": "BhG0RwO5vx1DZxNvyPik",
                    "date": "0001-01-01T00:00:00Z",
                    "isComplete": false,
                    "willRepeat": false,
                    "repeatFrequency": "never",
                    "end_repeat": "0001-01-01T00:00:00Z",
                    "remind": false,
                    "emailSelected": false,
                    "discordSelected": false,
                    "reminder": "none",
                    "reminder_time": "0001-01-01T00:00:00Z",
                    "priority": "none",
                    "shared": false,
                    "subTasks": [
                        ""
                    ]
                }
            ],
            [
                {
                    "id": "CwLns7muiQ0u3PKA8OkD",
                    "text": "First Task !",
                    "task_owner": "q0nQCfPpjneYyCXAbXkK",
                    "parent_id": "ZhDmSusiguNCIp9A9H8x",
                    "date": "0001-01-01T00:00:00Z",
                    "isComplete": false,
                    "willRepeat": false,
                    "repeatFrequency": "never",
                    "end_repeat": "0001-01-01T00:00:00Z",
                    "remind": false,
                    "emailSelected": false,
                    "discordSelected": false,
                    "reminder": "none",
                    "reminder_time": "0001-01-01T00:00:00Z",
                    "priority": "none",
                    "shared": false,
                    "subTasks": [
                        ""
                    ]
                },
                {
                    "id": "cMYyUwSbLDZ9Vl2WZIGe",
                    "text": "test_task_1",
                    "task_owner": "q0nQCfPpjneYyCXAbXkK",
                    "parent_id": "ZhDmSusiguNCIp9A9H8x",
                    "date": "2006-01-02T15:04:05Z",
                    "isComplete": true,
                    "willRepeat": false,
                    "repeatFrequency": "never",
                    "end_repeat": "0001-01-01T00:00:00Z",
                    "remind": false,
                    "emailSelected": false,
                    "discordSelected": false,
                    "reminder": "none",
                    "reminder_time": "0001-01-01T00:00:00Z",
                    "priority": "none",
                    "shared": false,
                    "subTasks": [
                        "sub_task_1"
                    ]
                }
            ],
            [
                {
                    "id": "gs7rdjwRrBQwGahGjLYB",
                    "text": "First Task !",
                    "task_owner": "q0nQCfPpjneYyCXAbXkK",
                    "parent_id": "iBYGoFECohgBloBCEgC4",
                    "date": "0001-01-01T00:00:00Z",
                    "isComplete": false,
                    "willRepeat": false,
                    "repeatFrequency": "never",
                    "end_repeat": "0001-01-01T00:00:00Z",
                    "remind": false,
                    "emailSelected": false,
                    "discordSelected": false,
                    "reminder": "none",
                    "reminder_time": "0001-01-01T00:00:00Z",
                    "priority": "none",
                    "shared": false,
                    "subTasks": [
                        ""
                    ]
                }
            ],
            [
                {
                    "id": "3vg86jNhau3Szgo9EwTc",
                    "text": "First Task !",
                    "task_owner": "q0nQCfPpjneYyCXAbXkK",
                    "parent_id": "s4LFfNmxchYpSops4zlU",
                    "date": "0001-01-01T00:00:00Z",
                    "isComplete": false,
                    "willRepeat": false,
                    "repeatFrequency": "never",
                    "end_repeat": "0001-01-01T00:00:00Z",
                    "remind": false,
                    "emailSelected": false,
                    "discordSelected": false,
                    "reminder": "none",
                    "reminder_time": "0001-01-01T00:00:00Z",
                    "priority": "none",
                    "shared": false,
                    "subTasks": [
                        ""
                    ]
                }
            ]
        ]
    }
}
```

### Edit List
`http://localhost:10000/update/{uid}/list/{id}?<params>`

**Example**

URL

`http://localhost:10000/update/a3a1hWUx5geKB8qeR6fbk5LZZGI2/list/364DgExvwpE4lNC7JV59?list_name=list1updated&lock=false`

Return

```json

```

### Edit Task
`http://localhost:10000/update/{uid}/task/{id}?<params>`

**Example**

URL

`http://localhost:10000/update/SFCCBJMyEA3NyBnCRe4j/task/jHExtupOREWA1CyHcEUX?done=true&discord=true`

Return

```json
{
    "result": {
        "User": null,
        "List": null,
        "Lists": null,
        "Task": {
            "id": "cMYyUwSbLDZ9Vl2WZIGe",
            "text": "test_task_1",
            "task_owner": "q0nQCfPpjneYyCXAbXkK",
            "parent_id": "ZhDmSusiguNCIp9A9H8x",
            "date": "2006-01-02T15:04:05Z",
            "isComplete": true,
            "willRepeat": false,
            "repeatFrequency": "never",
            "end_repeat": "0001-01-01T00:00:00Z",
            "remind": false,
            "emailSelected": false,
            "discordSelected": false,
            "reminder": "none",
            "reminder_time": "0001-01-01T00:00:00Z",
            "priority": "none",
            "shared": false,
            "subTasks": [
                "sub_task_1"
            ]
        },
        "Tasks": null,
        "AllTasks": null
    }
}
```

# TESTS
navigate to `bin/database`

run `go test -v`

return (so far) -

```
(base) sabra@Sabras-MacBook-Pro database % go test -v
=== RUN   TestCreateUser
--- PASS: TestCreateUser (2.57s)
=== RUN   TestCreateList
--- PASS: TestCreateList (1.20s)
=== RUN   TestCreateListWithPayload
--- PASS: TestCreateListWithPayload (1.10s)
=== RUN   TestCreateTaskWithPaylod
--- PASS: TestCreateTaskWithPaylod (0.58s)
=== RUN   TestCreateSubTask
--- PASS: TestCreateSubTask (0.40s)
=== RUN   TestGetUser
--- PASS: TestGetUser (1.13s)
=== RUN   TestGetList
--- PASS: TestGetList (0.54s)
=== RUN   TestGetLists
--- PASS: TestGetLists (1.09s)
=== RUN   TestGetTask
--- PASS: TestGetTask (0.20s)
=== RUN   TestUpdateTask
--- PASS: TestUpdateTask (0.37s)
=== RUN   TestUpdateUser
--- PASS: TestUpdateUser (1.45s)
=== RUN   TestGetTasks
--- PASS: TestGetTasks (0.47s)
=== RUN   TestDestroyUser
--- PASS: TestDestroyUser (3.45s)
PASS
ok  	database/bin/database	14.556s
```

***I DO NOT GUARANTEE ANY OTHER FUNCTION WILL WORK IF IT DOES NOT HAVE A PASSING TEST***
