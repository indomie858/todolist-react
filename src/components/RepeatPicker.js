import React from 'react';

const RepeatPicker  = (props) => {
    return (  
        <div>
            <form>
                <select value={props.isEvery} onChange={props.changeRepeats}>
                        <option value="Once">Once</option>
                        <option value="Every">Every</option>
                </select>

                <span>
                    <input type="number" value={props.numDays} onChange={props.onChange} className="numberInput"/>
                    <label>Days</label>
                </span>
                
            </form>
        </div>
    );
}
 
export default RepeatPicker;