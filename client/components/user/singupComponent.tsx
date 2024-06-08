import React, { useState } from 'react';
import '@/styles/user-styles.module.css';
import { SignupAxios } from '@/server/user';
import SignupForm from '@/ui/user/signupForm';


const SignUpComponent = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [showSignupForm, setShowSignupForm] = useState(false);

  const handleSubmit = (e:any) => {
    e.preventDefault();
    console.log("username -> ", username)
    SignupAxios(username, password).then((res) => {
      alert("User created successfully")
      setUsername(''); setPassword('');
    }).catch((err) => {
      alert("User created failed")
      setUsername(''); setPassword('');
      console.log("err -> ", err)
    }
    );
  };

  const handleCreateChatClick = () => {
    setShowSignupForm(true);
  }

  return (
    <div className="signup-form">
      {showSignupForm ? (  
        <SignupForm
          username={username}
          setUsername={setUsername}
          password={password}
          setPassword={setPassword}
          handleSubmit={handleSubmit}
          setShowSignupForm={setShowSignupForm}
        /> 
        ) : <button onClick={handleCreateChatClick}>Sign Up</button> }
        
      {/* <form onSubmit={handleSubmit}>
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
      </form> */}
    </div>
  );
};

export default SignUpComponent;
