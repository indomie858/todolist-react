import React, {useState} from "react";
import DatePicker from 'react-date-picker';
import TimePicker from 'react-time-picker';
import RepeatPicker from "./RepeatPicker";
import AddSubForm from "./AddSubForm"
import discordImage from "./discord.png";
import MailOutlineIcon from '@material-ui/icons/MailOutline';


const AddTask = (props) => {

    const [userLists, setUserLists] = useState(props.userLists)
    const [templateValue, setTemplateValue] = useState('');
    const [id, setId] = useState(props.id ? props.id : '');
    const [taskValue, setTaskValue] = useState(props.text ? props.text : "");
    const [dateValue, setDateValue] = useState(props.date ? new Date(props.date) : new Date());
    const [listValue, setListValue] = useState(props.list ? props.list : "Main");
    const [willRepeat, setWillRepeat] = useState(props.willRepeat ? props.willRepeat : false);
    const [showTime, toggleShowTime] = useState(props.reminder ? props.reminder : false);
    const [time, setTime] = useState(props.date ? new Date(props.date) : new Date());
    const [repeatFrequency, setRepeatFrequency] = useState(props.repeatFrequency ? props.repeatFrequency : "Never");
    const [emailSelected, setEmailSelected] = useState(props.emailSelected ? props.emailSelected : props.defaultReminders['email']);
    const [discordSelected, setDiscordSelected] = useState(props.discordSelected !== undefined ? props.discordSelected : props.defaultReminders['discord']);
    const [endRepeat, setEndRepeat] = useState(props.endRepeat ? new Date(props.endRepeat) : new Date());
    const [showSubtasks, setShowSubtasks] = useState(false);
    const [subtaskValue, setSubtaskValue] = useState('');
    const [subtaskArr, setSubtaskArr] = useState([]);

    return ( 
        <div>
            <div className="popover" >
                <span>Add New Task:</span>
                <label className="templateInput">Template: &nbsp;
                    <select value={templateValue} onChange={(e) => {
                        //sorry for how nasty this select element is
                        setTemplateValue(e.target.value);
                        if (e.target.value !== ''){
                            toggleShowTime(true);
                            setTaskValue(e.target.value);
                        } else {
                            toggleShowTime(false);
                            setTaskValue('');
                        }                        
                        }}>
                        <option value=""></option>
                        <option value="Study">Study</option>
                        <option value="Clean Room">Clean Room</option>
                        <option value="Do Laundry">Do Laundry</option>
                        <option value="Buy Groceries">Buy Groceries</option>
                        <option value="Relax">Relax</option>
                    </select>
                </label>
                <input 
                    className="textInput"
                    type="text" 
                    value={taskValue} 
                    onChange={(e) => setTaskValue(e.target.value)} autoFocus/>
                <DatePicker value={dateValue}  onChange={setDateValue} className="addTaskInput"/>
                <div>Repeat? &nbsp;
                    <input type="checkbox" checked={willRepeat} onChange={() => setWillRepeat(!willRepeat)}/>&nbsp;
                    Remind? &nbsp;
                    <input type="checkbox"
                        checked={showTime} 
                        onChange={() => toggleShowTime(!showTime)}/>
                </div>
                {showTime && <TimePicker onChange={setTime} value={time} className="addTaskInput" disableClock={true} clearIcon={null} clockIcon={null}/>}
                {showTime && <MailOutlineIcon className={emailSelected ? "outlined icon" : "icon"} onClick={() => setEmailSelected(!emailSelected)}/>}
                {showTime && <img alt="discord icon" className={discordSelected ? "discord outlined icon" : "discord icon"} onClick={() => setDiscordSelected(!discordSelected)} src={discordImage}/>}
                {willRepeat && <RepeatPicker 
                    frequency={repeatFrequency} changeRepeats={(e) => setRepeatFrequency(e.target.value)} endDate={endRepeat}
                    changeEndDate={setEndRepeat}/>}
                <label className="addTaskInput">List: &nbsp;
                    <select value={listValue} onChange={(e) => setListValue(e.target.value)}>
                        {userLists.map((name, id) => (
                            <option value={name[0]} key={id[1]}>{name[0]}</option>
                        ))}
                    </select>
                </label>
                <div>
                    Subtasks? &nbsp;
                    <input type="checkbox" 
                        checked={showSubtasks}
                        onChange={() => setShowSubtasks(!showSubtasks)}/> &nbsp;
                </div>
                <div className="subtaskSectionAdd">
                    {showSubtasks && <AddSubForm setSubtaskValue={setSubtaskValue} subtaskValue=    {subtaskValue} subtaskArr={subtaskArr} setSubtaskArr={setSubtaskArr}/>}
                    {(subtaskArr && showSubtasks) ? subtaskArr.map((subTask) => (<p>{subTask}</p>)) : ''}
                </div>
                <div className="addTaskInput">
                    <button className="addTaskInput addTaskButton" onClick={() => props.onAdd({
                        id: id,
                        date: dateValue,
                        discordSelected: discordSelected,
                        emailSelected: emailSelected,
                        end_repeat: endRepeat,
                        list: listValue,
                        remind: showTime,
                        reminder_time: time,
                        repeatFrequency: repeatFrequency,
                        text: taskValue,
                        willRepeat: willRepeat
                    })}>Add</button>
                    <button className="addTaskInput addTaskButton" onClick={() => props.onCancel()}>Cancel</button>
                </div>
            </div>
            <div className="popoverTag1 popoverRight1"></div>
            <div className="popoverTag2 popoverRight2"></div>
        </div>
    );
}

export default AddTask
