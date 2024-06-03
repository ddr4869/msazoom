const SignupForm = ( {username, setUsername, password, setPassword, handleSubmit, setShowSignupForm} ) => { 
    return (     
        <form onSubmit={handleSubmit}>
            <input 
            type="text" 
            placeholder="Username" 
            value={username} 
            onChange={(e) => setUsername(e.target.value)} 
            />
            <input 
            type="password" 
            placeholder="Password" 
            value={password} 
            onChange={(e) => setPassword(e.target.value)} 
            />
            <button type="submit">Sign Up</button>
            <button type="button" onClick={() => setShowSignupForm(false)}>Cancel</button>
        </form>
    )
}

export default SignupForm;