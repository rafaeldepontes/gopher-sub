import SubButton from "@/components/SubButton"

const HomePage = () => {
    // Adding some type of context for holding the users data after the log in...

    return (
        <>
            <h1>Hi random user!</h1>
            <br />
            <p>click here to receive news about our product:</p>
            <SubButton />
        </>
    )
}

export default HomePage