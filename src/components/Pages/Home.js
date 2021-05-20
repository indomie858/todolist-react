import React from 'react'
import { useState } from 'react';
import { Redirect } from 'react-router-dom';
import Header from '../Header';
import Tasks from '../Tasks';
import BottomNavBar from '../BottomNavBar';
import AddTask from '../AddTask.js'
import ListNav from '../ListNav.js'
import Options from '../Options.js'
import Container from '@material-ui/core/Container';

const Home = () => {

  const getToken = () => {
    //tokens are stored locally so user doesn't have to keep logging in
    const token = sessionStorage.getItem('token');
    return token
  };

  //state for displaying tasks. currently has placeholder objects. will replace with tasks from database
  const [tasks, setTasks] = useState(
    [
      {
        id: 1,
        text: 'Doctors appt',
        date: '2021-02-05T14:00',
        list: 'Main',
        willRepeat: true,
        repeatFrequency: "Every",
        repeatNumDays: 1,
        emailSelected: true,
        discordSelected: false,
        reminder: true,
        isComplete: false,
        subTasks:
          [
            'Get in car',
            'Drive to Doctor'
          ],
      },
      {
        id: 2,
        text: 'School meeting',
        date: '2021-02-05T14:00',
        list: 'Main',
        willRepeat: true,
        repeatFrequency: "Every",
        repeatNumDays: 1,
        emailSelected: true,
        discordSelected: false,
        reminder: true,
        isComplete: true,
        subTasks:
          [
            'Take two shots',
            'Put on some pants',
            'Get a ride to school'
          ],
      },
      {
        id: 3,
        text: 'Food shopping',
        date: '2021-02-05T14:00',
        list: 'Main',
        willRepeat: true,
        repeatFrequency: "Every",
        repeatNumDays: 1,
        emailSelected: true,
        discordSelected: false,
        reminder: true,
        isComplete: false,
        subTasks:
          [
            'Make shopping list',
            'Drive to costco',
            'Buy some shit'
          ],
      },
      {
        id: 4,
        text: 'something something darkside',
        date: '2021-02-05T14:00',
        list: 'Main',
        willRepeat: true,
        repeatFrequency: "Every",
        repeatNumDays: 1,
        emailSelected: true,
        discordSelected: false,
        reminder: true,
        isComplete: false,
        subTasks: [
        ],
      },
    ]
  )


  const [showListNav, setListNav] = useState(false);
  const [showOptions, setOptions] = useState(false);
  const [showAddTask, setAddTask] = useState(false);
  const [showChangeTask, setChangeTask] = useState(false);
  const [changingTask, setChangingTask] = useState(0);
  const [email, setEmail] = useState('');

  if (!getToken()) {
    console.log('/home token does not exist');
    return (<Redirect to="/login" />);
  } else {
    console.log('/home token exists');
    if (email === '') {
      setEmail(sessionStorage.getItem('email'));
    }
  }
  return (
    <>
      {/* <Container maxWidth="xs"> */}
      <p>Welcome {email}</p>
      <div className="mainContainer">
        {showAddTask && <AddTask onAdd={() => setAddTask(false)} defaultReminders={{ "discord": true, "email": false }} onCancel={() => setAddTask(false)}/>}
        {showChangeTask && <AddTask onAdd={() => setAddTask(false)} defaultReminders={{ "discord": true, "email": false }} onCancel={() => setAddTask(false)}
          date={changingTask.date}
          text={changingTask.text}
          list={changingTask.list}
          willRepeat={changingTask.willRepeat}
          reminder={changingTask.reminder}
          repeatFrequency={changingTask.repeatFrequency}
          numDays={changingTask.repeatNumDays}
          emailSelected={changingTask.emailSelected}
          discordSelected={changingTask.discordSelected}
        />}
        {showListNav && <ListNav onChooseList={() => setListNav(false)} lists={[{ name: "Main List" }, { name: "Some Shared List" }, { name: "Some Other List" }]} />}
        {showOptions && <Options defaultList={"Shared"} defaultReminders={{ "discord": true, "email": false }} />}
        <Header />
        <div className='listContainer'>
          {/* displays placeholder list and title "Today" */}
          {tasks.length > 0 ? (<Tasks tasks={tasks} listTitle='Today' changeTask={(id) => {
            // console.log("Changing task")
            // console.log("id: " + id)
            for (let i = 0; i < tasks.length; i++) {
              if (tasks[i].id === id) {
                setChangingTask(tasks[i])
                // console.log(changingTask)
              }
            }
            setChangeTask(!showChangeTask);
            setListNav(false);
            setOptions(false);
            setAddTask(false);
          }} />) : ('No tasks to show')}
        </div>
      </div>
      <BottomNavBar onAddTask={() => { setAddTask(!showAddTask); setChangeTask(false); setListNav(false); setOptions(false) }} onListNav={() => { setListNav(!showListNav); setChangeTask(false); setAddTask(false); setOptions(false) }} onOptions={() => { setListNav(false); setChangeTask(false); setAddTask(false); setOptions(!showOptions) }} />
      {/* </Container> */}
    </>
  )
}

export default Home
