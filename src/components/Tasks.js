import Task from './Task'

const Tasks = ({ tasks, listTitle }) => {
    return (
        <>
            <h2>{listTitle}</h2>
            {tasks.map((task) => (
                <Task key={task.id} task={task} />
            ))}
        </>
    )
}

export default Tasks