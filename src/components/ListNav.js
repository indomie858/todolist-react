import React, {useState} from "react";
import ListNavItem from "./ListNavItem.js"
import {DragDropContext, Droppable,Draggable} from "react-beautiful-dnd"


const ListNav = (props) => {

    const [lists, updateLists] = useState(props.lists)

    function handleOnDragEnd(result) {
        if (!result.destination) return;
        const items = Array.from(lists);
        const [reorderedItem] = items.splice(result.source.index, 1);
        items.splice(result.destination.index, 0, reorderedItem);

        updateLists(items);
    }

    return ( 
        <div>
            <div className="popover">
                <div className="listHeader">Change List:</div>
                <DragDropContext onDragEnd={handleOnDragEnd}>    
                    <Droppable droppableId="listOfLists">
                        {(provided) => (
                            <ul className="listOfLists" {...provided.droppableProps} ref={provided.innerRef}>
                                {lists.map(({name}, index) => (
                                    <Draggable key={name} draggableId={name} index={index}>
                                        {(provided) => (                                            
                                            <li {...provided.draggableProps} {...provided.dragHandleProps} ref={provided.innerRef}>
                                                <ListNavItem list={name}/>
                                            </li>
                                        )}
                                    </Draggable>
                                ))}
                                {provided.placeholder}
                            </ul>
                        )}
                    </Droppable>
                </DragDropContext>
            </div>
            <div className="popoverTag1 popoverLeft1"></div>
            <div className="popoverTag2 popoverLeft2"></div>
        </div>
    );
}

export default ListNav
