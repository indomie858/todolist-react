//component for repeat option in add task form component
import {React} from 'react';
import DatePicker from 'react-date-picker';

const RepeatPicker  = (props) => {

    return (  
        <div className="centerAlignContainer">
            <select value={props.frequency} onChange={props.changeRepeats}>
                    <option value="Daily">Daily</option>
                    <option value="Weekly">Weekly</option>
                    <option value="Monthly">Monthly</option>
                    <option value="Yearly">Yearly</option>
            </select>

            <div className="verticalSpace">End Repeat: </div>

            <DatePicker value={props.endDate}  onChange={props.changeEndDate} className="addTaskInput"/>
        </div>
    );
}
 
export default RepeatPicker;