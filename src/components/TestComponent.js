import { useState } from 'react'
import {Radio, RadioGroup,FormControl, FormControlLabel, } from '@material-ui/core';

// import FormLabel from '@material-ui/core/FormLabel';

const TestComponent = () => {
    //state for textfield. clears when you start typing
    const [text, setText] = useState('');

    //state for output div
    const [output, setOutput] = useState('');

    const [options, setOptions] = useState('');
    const [radio, setRadio] = useState('read');

    
    const onRadioChange = (e) => {
        setRadio(e.target.value);
    }

    const onSubmit = (e) => {
        // stops page from reloading on submit
        e.preventDefault();

        // set output div to whatever is inside input text
        

        //clear input text
        setText('');
        setOptions('');




        //testing post request
        // fetch('https://jsonplaceholder.typicode.com/posts', {
        //     method: 'POST',
        //     headers: {
        //         Accept: 'application/json',
        //         'Content-Type': 'application/json'
        //     },
        //     body: JSON.stringify({
        //         //pass in value of input text in body of request
        //         firstParam: {text},
        //     })
        // });

        switch(radio){
            case "create":
                break;
            case "read":
                if(options===''){
                    fetch('http://localhost:3003/api/'+text,{
                        method: 'GET',
                
                
                    }).then(response => {
                        if(response.status===404){
                            return "Error: 404"
                        }else{
                            return response.json()}
                    }).then(data=>setOutput(JSON.stringify(data)));
                }
                else{
                    fetch('http://localhost:3003/api/userData/'+text+"/"+options,{
                        method: 'GET',
                
                
                    }).then(response => {
                        if(response.status===404){
                            return "Error: 404"
                        }else{
                            return response.json()}
                    }).then(data=>setOutput(JSON.stringify(data)));
                }
            break;
            case "update":
                fetch('http://localhost:3003/api/update/'+text+"/"+options,{
                    method: 'GET',
            
            
                }).then(response => {
                    if(response.status===404){
                        return "Error: 404"
                    }else{
                        return response.json()}
                }).then(data=>setOutput(JSON.stringify(data)));


                
            break;
            case "destroy":
            break;

            default:
                setOutput(text);
            break;
        }

       
        
        
    }

    return (
        //test form with text input and submit button
        <div>
            <FormControl component="fieldset">
        <form onSubmit={onSubmit}>
            <input
                type='text'
                placeholder='type here'
                value={text}
                onChange={(e) => setText(e.target.value)} />
            <input type='text' placeholder='optionals' value={options}
            onChange={(e) => setOptions(e.target.value)}/>
            {/* <FormLabel component="legend">labelPlacement</FormLabel> */}
      <RadioGroup row aria-label="radios" name="radios" value={radio} defaultValue="read" onChange={onRadioChange}>
        <FormControlLabel
                        value="create"
                        control={<Radio color="primary" />}
                        label="Create"
                        labelPlacement="top"
                        />
        <FormControlLabel
                        value="read"
                        control={<Radio color="primary" />}
                        label="Read"
                        labelPlacement="top"
                        />
        <FormControlLabel
                        value="update"
                        control={<Radio color="primary" />}
                        label="Update"
                        labelPlacement="top"
                        />
        <FormControlLabel value="destroy" 
                        control={<Radio color="primary" />} 
                        label="Destroy"
                        labelPlacement="top" 
                        />
      </RadioGroup>
            <input type='submit' value='click me!' />

            <div>{output}</div>
        </form>
        </FormControl>
        </div>
    )
}

export default TestComponent
