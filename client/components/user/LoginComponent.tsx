import LoginForm from '@/ui/user/loginForm';
import userStyles from '@/styles/user-styles.module.css';

interface LoginComponentProps {
  username: string;
  setUsername: (username: string) => void;
  password: string;
  setPassword: (password: string) => void;
  handleLogin: (
    username: string, password: string, 
    setIsLoggedIn: (isLoggedIn: boolean) => void, 
    event: React.FormEvent<HTMLFormElement> | React.MouseEvent<HTMLButtonElement>
  ) => void;
  handleLogout: () => void;
  isLoggedIn: boolean;
}

const LoginComponent = ({ username, setUsername, password, setPassword, handleLogin, handleLogout, isLoggedIn }:LoginComponentProps) => {
  return (
    <header>
      <h1>Chat Board</h1>
      <div id="login-section">
      
        {!isLoggedIn && <h2>Login</h2>}
        {!isLoggedIn ? (
          <LoginForm
            username={username}
            setUsername={setUsername}
            password={password}
            setPassword={setPassword}
            handleLogin={handleLogin}
          />
        ) : (
          <UserProfile handleLogout={handleLogout} />
        )}
      </div>
      <br></br>
      <hr></hr>
    </header>
  );
};

const UserProfile = ({handleLogout}:any ) => (
    <div className={userStyles.gridUserContainer}>
      <div className={`${userStyles.userStyles} ${userStyles.userProfile}`}>
        <p>Welcome, <strong>{localStorage.getItem('username')}</strong></p>
        <button onClick={handleLogout}>Logout</button>
      </div>
      {/* <div className={`${userStyles.userStyles} ${userStyles.additionalInfo}`}>
        <p>Email: example@example.com</p>
        <p>Role: Administrator</p>
      </div> */}
    </div>
  );

export default LoginComponent;