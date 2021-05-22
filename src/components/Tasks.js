//component for list of tasks

import Task from './Task'

const Tasks = (props) => {
    return (
        <>
            <h2>{props.listTitle}</h2>
            {/* iterates through tasks object and passes key/values to task component */}
            {props.tasks.filter((task) => {
                // we filter out completed tasks so they aren't shown
                if (task.isComplete === true) {
                    return false;
                }
                return true;
            }).map((task) => (
                <Task key={task.id} id={task.id} markCompleted={props.markCompleted} task={task} changeTask={(id) => props.changeTask(id)} />
            ))}
            
        </>
    )
}

export default Tasks