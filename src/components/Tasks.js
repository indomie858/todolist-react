//component for list of tasks

import Task from './Task'

const Tasks = ({ tasks, listTitle }) => {
    return (
        <>
            <h2>{listTitle}</h2>
            {/* iterates through tasks object and passes key/values to task component */}
            {tasks.map((task) => (
                <Task key={task.id} task={task} />
            ))}
        </>
    )
}

export default Tasks