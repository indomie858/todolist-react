//component for individual subtask

const SubTask = ({ subTask }) => {
    return (
        // outputs subtask from subtasks array
        <div className="subtask">
            <p>{subTask}</p>
        </div>
    )
}

export default SubTask
