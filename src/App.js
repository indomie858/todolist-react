import TestComponent from './components/TestComponent'

import { useState } from 'react'
import Header from './components/Header'
import Tasks from './components/Tasks'
import BottomNavBar from './components/BottomNavBar'
import AddTask from './components/AddTask.js'

import Container from '@material-ui/core/Container';
import { SettingsApplications } from '@material-ui/icons'

const App = () => {
  //state for displaying tasks. currently has placeholder objects. will replace with tasks from database
  const [tasks, setTasks] = useState(
    [
      {
        id: 1,
        text: 'Doctors appt',
        day: 'Feb 5th at 2:30pm',
        reminder: true,
      },
      {
        id: 2,
        text: 'School meeting',
        day: 'Feb 6th at 1:30pm',
        reminder: true,
      },
      {
        id: 3,
        text: 'Food shopping',
        day: 'Feb 5th at 5:30pm',
        reminder: false,
      },
      {
        id: 1,
        text: 'Doctors appt',
        day: 'Feb 5th at 2:30pm',
        reminder: true,
      },
      {
        id: 2,
        text: 'School meeting',
        day: 'Feb 6th at 1:30pm',
        reminder: true,
      },
      {
        id: 3,
        text: 'Food shopping',
        day: 'Feb 5th at 5:30pm',
        reminder: false,
      },
    ]
  )

  const [showAddTask, setAddTask] = useState(true);

  return (
    <>
    <TestComponent />
    {/* Container is from material-ui library */}
    <Container maxWidth="xs">
    <Header />
      <div className='listContainer'>
        {/* displays placeholder list and title "Today" */}
        {tasks.length > 0 ? (<Tasks tasks={tasks} listTitle='Today' />) : ('No tasks to show')}
      </div>
    </Container>
    {showAddTask && <AddTask onAdd={() => setAddTask(false)}/>}
    <BottomNavBar onAddTask={() => setAddTask(true)} />
    </>
  );
}

export default App;
