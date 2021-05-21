# Getting Started with Node Server

Navigate to `node/server/` and run `npm install`

You can then run the server using `npm start`.

This will start the server, serving on port 3003.

##Adding Modules
If you need to add a module or dependency to the project, run `npm install --save 'module-name'`
This will save your dependency to the package.json for others to install with.

To make api calls, you can navigate to `http://localhost:3003/api` in your browser.
Active calls are:





#Create:
URL: `/api/create/{useridcode}`
Method: post
Necessary JSON Objects in Body:
create: {user||list||task||subtask}
if list:
    list_name: "listname"
if task:
    task_name: "taskname"
    parentId: "parentidcode"
if subtask:
    subtask_name: "subtask_name"
    parentId: "parentidcode"

#Read:
Method: get
URL's:
    All Lists: 
        `/api/userData/{useridcode}/lists`
    One List:
        `/api/userData/{useridcode}/list/{listidcode}`
    User:
        `/api/userData/{useridcode}`
    User Name:
        `/api/userData/{useridcode}/name`
    User Email:
        `/api/userData/{useridcode}/email`
    User Status:
        `/api/userData/{useridcode}/status`


#Update:
URL: `/api/update/{useridcode}`
Method: post
Necessary JSON Objects in Body:
update: {userSettings||listSettings||taskSettings||subtasks}
if listSettings:
    listId: "listidcode"
if taskSettings:
    taskId: "taskidcode"
if subtasks: //note: MUST ONLY HAVE `update, taskID, sub_tasks`.
    taskId: "taskId"
    sub_tasks: ["array", "of", "string", "subtasks"]

#Delete:
URL: `/api/delete/{useridcode}`
Method: delete
delete: {user||list||task||subtask}
if list:
    listId: "listid"
if task:
    taskId: "taskid"
if subtask:
    sub_tasks: ["array", "of", "string", "subtasks"]
