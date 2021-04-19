const createError = require('http-errors');
const express = require('express');
const path = require('path');
const fs = require('fs')
const http = require('http')
// importing the dependencies
const cookieParser = require('cookie-parser');

// const bodyParser = require('body-parser');
const cors = require('cors');
const helmet = require('helmet');
const logger = require('morgan');

const indexRouter = require('./routes/index');
const usersRouter = require('./routes/users');

const app = express();

let jsonObject = require('./test.json')

// view engine setup
app.set('views', path.join(__dirname, 'views'));
app.set('view engine', 'ejs');

app.use(logger('dev'));
// using bodyParser to parse JSON bodies into JS objects
app.use(express.json());
// app.use(express.urlencoded({ extended: false }));
app.use(cookieParser());
app.use(express.static(path.join(__dirname, 'public')));

app.use('/', indexRouter);
app.use('/users', usersRouter);

// adding Helmet to enhance your API's security
app.use(helmet());

// enabling CORS for all requests
app.use(cors());

// var server = http.createServer(app);
// server.listen(port);
// server.on('error', onError);
// server.on('listening', onListening);







// defining an endpoint to return all ads
app.get('/test1', (req, res) => {
  res.send("you rang?");
  res.end("hello?")
});


app.get('/api/:input',(req,res) =>{
  if(req.params.input in jsonObject.users){
    res.send(JSON.stringify(jsonObject.users[req.params.input]))
  }
  else{
    setCustomError(req,res,"Invalid API Reference. Please check the path again.","Invalid API Reference",404)
    res.render('error')
  }
})


app.get('/api/userData/:id',(req,res)=>{
  

http.get('http://localhost/userData/'+req.params.id, (resp) => {
  let data = '';

  // A chunk of data has been received.
  resp.on('data', (chunk) => {
    data += chunk;
  });

  // The whole response has been received. Print out the result.
  resp.on('end', () => {
    console.log(JSON.parse(data).explanation);
  });

}).on("error", (err) => {
  console.log("Error: " + err.message);
});
})

function setCustomError(req,res,message,name,status){
  const err = createError(status,name)
  res.locals.title=name
  res.locals.message = message
  res.locals.error = req.app.get('env') === 'development' ? err : {};
  res.status(status)
  
}



// catch 404 and forward to error handler
app.use(function(req, res, next) {
  next(createError(404));
});

// error handler
app.use(function(err, req, res, next) {
  // set locals, only providing error in development
  res.locals.title = err.message;
  res.locals.error = req.app.get('env') === 'development' ? err : {};
  
  // render the error page
  res.status(err.status || 500);
  res.locals.message= res.statusCode + " "+err.message
  res.render('error');
});

function errorHandler(err, req, res, next){
  
}

// starting the server
app.listen(3003, () => {
  console.log('listening on port 3003');
});
module.exports = app;
