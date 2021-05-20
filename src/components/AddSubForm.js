//component for adding a sub task from the AddTask component form.

const AddSubForm = (props) => {

    const handleSubmit = (e) => {
        e.preventDefault();
        //call function to submit with new task here
        console.log(props.subtaskValue);
        //adds subtask to subtask array in AddTask
        props.setSubtaskArr([...props.subtaskArr, props.subtaskValue]);
    }

    return (
        <>
        <form className="addSubForm" onSubmit={handleSubmit}>
            <input type="text" onChange={(e) => {
                props.setSubtaskValue(e.target.value)
                }}/>
            <button className="addTaskButton" type="submit">Add subtask</button>
        </form>
        </>
    )
}

export default AddSubForm
