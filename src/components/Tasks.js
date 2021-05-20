//component for list of tasks

import Task from './Task'

const Tasks = (props) => {
    return (
        <>
            <h2>{props.listTitle}</h2>
            {/* iterates through tasks object and passes key/values to task component */}
            {props.tasks.map((task) => (
                <Task key={task.id} id={task.id} task={task} changeTask={(id) => props.changeTask(id)} />
            ))}
        </>
    )
}

export default Tasks