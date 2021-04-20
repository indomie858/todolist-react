# Getting Started with Node Server

Navigate to `node/server/` and run `npm install`

You can then run the server using `npm start`.

This will start the server, serving on port 3003.

##Adding Modules
If you need to add a module or dependency to the project, run `npm install --save 'module-name'`
This will save your dependency to the package.json for others to install with.

To make api calls, you can navigate to `http://localhost:3003/api` in your browser.
Active calls are:
`/userData/{useridcode}`
`/userData/{useridcode}/name`
`/userData/{useridcode}/email`
`/userData/{useridcode}/status`
`/userData/{useridcode}/lists`
