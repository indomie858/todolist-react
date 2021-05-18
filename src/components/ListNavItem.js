import React, {useState} from "react";
import DragIndicatorIcon from '@material-ui/icons/DragIndicator';

const ListNavItem = (props) => {

    return ( 
        <div className="listNavContainer">
            <DragIndicatorIcon/>
            <div className="listDisplay">        
                {props.list}
            </div>
        </div>
    );
}

export default ListNavItem
