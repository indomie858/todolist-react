//component for list of tasks

import Task from './Task'

const Tasks = ({ tasks, listTitle, changeTask }) => {
    return (
        <>
            <h2>{listTitle}</h2>
            {/* iterates through tasks object and passes key/values to task component */}
            {tasks.map((task) => (
                <Task key={task.id} id={task.id} task={task} changeTask={(id) => changeTask(id)} />
            ))}
        </>
    )
}

export default Tasks