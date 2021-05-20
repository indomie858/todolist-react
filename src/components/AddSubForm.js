const AddSubForm = (props) => {
    return (
        <div className="addSubForm">
            <input type="text" onChanged={(e) => props.setSubtaskValue(e.target.value)}/>
            <button>Add subtask</button>
        </div>
    )
}

export default AddSubForm
