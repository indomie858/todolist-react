import React from 'react'
import { useState, useEffect } from 'react';
import { Redirect } from 'react-router-dom';
import Header from '../Header';
import Tasks from '../Tasks';
import BottomNavBar from '../BottomNavBar';
import AddTask from '../AddTask.js'
import ListNav from '../ListNav.js'
import Options from '../Options.js'
import Container from '@material-ui/core/Container';
import { LensTwoTone } from '@material-ui/icons';

const Home = () => {

  const userId = "a3a1hWUx5geKB8qeR6fbk5LZZGI2"; // TODO: Get this from them being logged in
  // const listId = 'updated_isolated_list';

  let viewingList = 'mJmia9sdFy6yfb134ygs';
  
  
  // fetch(`http://localhost:3003/api/userData/${userId}/list/${listId}`).then(
  //     data => {
  //       console.log(data); 
  //       data.text().then(
  //         value => { 
  //           console.log(value)
  //           let response = JSON.parse(value); 
  //           console.log(response)
  //           console.log(JSON.parse(response).result)
  //         }
  //       );
  //     }
  // );


  //setEmail(JSON.stringify(JSON.parse(value).result.name))


  const getToken = () => {
    //tokens are stored locally so user doesn't have to keep logging in
    const token = sessionStorage.getItem('token');
    return token
  };

  //state for displaying tasks. currently has placeholder objects. will replace with tasks from database
  const [tasks, setTasks] = useState(
    [
      {
        id: "FAKE1",
        text: "Starter 1",
        date: "0001-01-01T00:00:00Z",
        parent_id: "mJmia9sdFy6yfb134ygs",
        list: "updated_isolated_list",
        willRepeat: false,
        repeatFrequency: "Never",
        emailSelected: false,
        discordSelected: false,
        end_repeat: "0001-01-01T00:00:00Z",
        isComplete: false,
        lock: true,
        priority: "none",
        remind: false,
        reminder: "none",
        reminder_time: "0001-01-01T00:00:00Z",
        shared: false,
        subTasks: [],
        sub_task: false,
        task_owner: "a3a1hWUx5geKB8qeR6fbk5LZZGI2"
      },
      {
        date: "0001-01-01T00:00:00Z",
        discordSelected: false,
        emailSelected: false,
        end_repeat: "0001-01-01T00:00:00Z",
        id: "FAKE2",
        isComplete: false,
        parent_id: "mJmia9sdFy6yfb134ygs",
        list: "updated_isolated_list",
        priority: "none",
        remind: false,
        reminder: "none",
        reminder_time: "0001-01-01T00:00:00Z",
        repeatFrequency: "never",
        shared: false,
        sub_task: false,
        subTasks: [],
        task_owner: "a3a1hWUx5geKB8qeR6fbk5LZZGI2",
        text: "Starter 2",
        willRepeat: false
      }
    ]
  )


  const [userLists, setUserLists] = useState([]);
  const [discordDefault, setDiscordDefault] = useState(false);
  const [emailDefault, setEmailDefault] = useState(false);

  function refreshTasks() {
    fetch(`http://localhost:3003/api/userData/${userId}`).then(
      data => data.text().then(
        value => {
          const userData = JSON.parse(value).result;
          const listsFromDb = userData.Lists;
          let listNames = []; 
          listsFromDb.forEach(list => {
            listNames.push([list.list_name, list.id])
          })
          setUserLists(listNames);
          let newTasks = []
          userData.AllTasks[0].forEach(task => {
            let parentList = null;
            listsFromDb.forEach(list => {
              if (list.id == task.parent_id) {
                parentList = list.list_name
              }
            })
            task.list = parentList;
            task.subTasks = [];
            newTasks.push(task);
          });
          setTasks(newTasks);
        }
      )
    );
  }

  useEffect(() => {
    refreshTasks();
    if (!getToken()) {
      console.log('/home token does not exist');
      return (<Redirect to="/login" />);
    } else {
      console.log('/home token exists');
      if (email === '') {
        setEmail(sessionStorage.getItem('email'));
        
      }
    }
  }, []);

  const [showListNav, setListNav] = useState(false);
  const [showOptions, setOptions] = useState(false);
  const [showAddTask, setAddTask] = useState(false);
  const [showChangeTask, setChangeTask] = useState(false);
  const [changingTask, setChangingTask] = useState(0);
  const [email, setEmail] = useState('');

  function updateTask(taskObject) {
    console.log("updating")
    setChangeTask(false);
    let parentId;
    userLists.forEach(([name, id]) => {
      if (name == taskObject.list) {
        parentId = id;
      }
    })

    console.log(taskObject)
    fetch('http://localhost:3003/api/update/'+userId,{
                    method: 'POST',
                    headers: {
                        Accept: 'application/json',
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({
                        
                        update: 'taskSettings', 
                        date: taskObject.date,
                        taskId: taskObject.id,
                        discordSelected: taskObject.discordSelected,
                        emailSelected: taskObject.emailSelected,
                        end_repeat: taskObject.end_repeat,
                        parent_id: parentId,
                        remind: taskObject.remind,
                        reminder_time: taskObject.reminder_time,
                        repeatFrequency: taskObject.repeatFrequency,
                        text: taskObject.text,
                        willRepeat: taskObject.willRepeat
                        
                    })
            
                }).then(response => {
                    if(response.status===404){
                        return "Error: 404"
                    }else{
                        return response.json()}
                }).then(data=>JSON.stringify(data));
  }

  
  return (
    <>
      {/* <Container maxWidth="xs"> */}
      <p>Welcome {email}</p>
      <div className="mainContainer">
        {showAddTask && <AddTask userLists={userLists} onAdd={() => {setAddTask(false); refreshTasks()}} defaultReminders={{ "discord": true, "email": false }} onCancel={() => setAddTask(false)} />}
        {showChangeTask && <AddTask userLists={userLists} onAdd={updateTask} defaultReminders={{ "discord": true, "email": false }} onCancel={() => setChangeTask(false)}
          id={changingTask.id}
          date={changingTask.date}
          text={changingTask.text}
          list={changingTask.list}
          willRepeat={changingTask.willRepeat}
          reminder={changingTask.remind}
          repeatFrequency={changingTask.repeatFrequency}
          emailSelected={changingTask.emailSelected}
          discordSelected={changingTask.discordSelected}
        />}
        {showListNav && <ListNav onChooseList={() => setListNav(false)} lists={[{ name: "Main List" }, { name: "Some Shared List" }, { name: "Some Other List" }]} />}
        {showOptions && <Options defaultList={"Shared"} defaultReminders={{ "discord": true, "email": false }} />}
        <Header />
        <div className='listContainer'>
          {/* displays placeholder list and title "Today" */}
          {tasks.length > 0 ? (<Tasks tasks={tasks} listTitle='Today' changeTask={
            (id) => {
              for (let i = 0; i < tasks.length; i++) {
                if (tasks[i].id === id) {
                  setChangingTask(tasks[i])
                }
              }
              setChangeTask(!showChangeTask);
              setListNav(false);
              setOptions(false);
              setAddTask(false);
          }
          }
           />) : ('No tasks to show')}
        </div>
      </div>
      <BottomNavBar onAddTask={() => { setAddTask(!showAddTask); setChangeTask(false); setListNav(false); setOptions(false) }} onListNav={() => { setListNav(!showListNav); setChangeTask(false); setAddTask(false); setOptions(false) }} onOptions={() => { setListNav(false); setChangeTask(false); setAddTask(false); setOptions(!showOptions) }} />
      {/* </Container> */}
    </>
  )
}

export default Home
