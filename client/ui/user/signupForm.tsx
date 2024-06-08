interface SignupFormProps {
    username: string;
    setUsername: (username: string) => void;
    password: string;
    setPassword: (password: string) => void;
    handleSubmit: (e: React.FormEvent) => void;
    setShowSignupForm: (showSignupForm: boolean) => void;
}

const SignupForm = ( {username, setUsername, password, setPassword, handleSubmit, setShowSignupForm}:SignupFormProps ) => { 
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