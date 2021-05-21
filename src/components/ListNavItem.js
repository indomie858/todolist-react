import React, {useState} from "react";
import DragIndicatorIcon from '@material-ui/icons/DragIndicator';

const ListNavItem = (props) => {

    return ( 
        <div className="listNavContainer" onClick={() => props.onChooseList(props.list)}>
            <DragIndicatorIcon/>
            <div className="listDisplay">        
                {props.list}
            </div>
        </div>
    );
}

export default ListNavItem
