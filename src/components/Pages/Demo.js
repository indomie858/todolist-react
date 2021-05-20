const Demo = () => {

    const handleSubmit = (e) => {
        e.preventDefault();

    }

    return (
        <div>
            <form onSubmit={() => handleSubmit()}>
                <input type="text" />
                <input type="submit" />
            </form>
        </div>
    )
}

export default Demo
