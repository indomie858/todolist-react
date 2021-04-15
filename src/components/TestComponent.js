import { useState } from 'react'

const TestComponent = () => {
    //state for textfield. clears when you start typing
    const [text, setText] = useState('');

    //state for output div
    const [output, setOutput] = useState('');

    const onSubmit = (e) => {
        // stops page from reloading on submit
        e.preventDefault();

        // set output div to whatever is inside input text
        setOutput(text);

        //clear input text
        setText('');


        //testing post request
        fetch('https://jsonplaceholder.typicode.com/posts', {
            method: 'POST',
            headers: {
                Accept: 'application/json',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                //pass in value of input text in body of request
                firstParam: {text},
            })
        });
    }

    return (
        //test form with text input and submit button
        <form onSubmit={onSubmit}>
            <input
                type='text'
                placeholder='type here'
                value={text}
                onChange={(e) => setText(e.target.value)} />
            <input type='submit' value='click me!' />

            <div>{output}</div>
        </form>
    )
}

export default TestComponent
