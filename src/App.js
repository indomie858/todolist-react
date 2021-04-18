import TestComponent from './components/TestComponent'

import { useState } from 'react'
import Header from './components/Header'
import Tasks from './components/Tasks'
import BottomNavBar from './components/BottomNavBar'

import Container from '@material-ui/core/Container';

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

  return (
    <>
    <TestComponent />

    <Container maxWidth="xs" >
      
        <Header />

        {tasks.length > 0 ? (<Tasks tasks={tasks} />) : ('No tasks to show')}

    </Container>
    <BottomNavBar />
    </>
  );
}

export default App;
