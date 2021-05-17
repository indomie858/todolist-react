import React, {useState} from "react";


const ListNavItem = (props) => {

    return ( 
        <div className="listDisplay">
            {props.list}
        </div>
    );
}

export default ListNavItem
