// component for list of subtasks

import SubTask from './SubTask'

const SubTasks = ({ subTasks }) => {
    return (
        <>
            {/* iterates through subtask array and passes each subtask to SubTask component */}
            {subTasks.map((subTask) => (<SubTask subTask={subTask} />))}
        </>
    )
}

export default SubTasks
