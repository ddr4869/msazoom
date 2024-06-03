import LoginForm from '@/ui/user/LoginForm';
import userStyles from '@/styles/user-styles.module.css';

const LoginComponent = ({ username, setUsername, password, setPassword, handleLogin, handleLogout, isLoggedIn }) => {
  return (
    <header>
      <h1>Chat Board</h1>
      <div id="login-section">
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
    </header>
  );
};

const UserProfile = ({ handleLogout }) => (
    <div className={userStyles.gridUserContainer}>
      <div className={`${userStyles.userStyles} ${userStyles.userProfile}`}>
        <p>Welcome, <strong>{localStorage.getItem('username')}</strong></p>
        <button onClick={handleLogout}>Logout</button>
      </div>
      <div className={`${userStyles.userStyles} ${userStyles.additionalInfo}`}>
        <p>Email: example@example.com</p>
        <p>Role: Administrator</p>
      </div>
    </div>
  );

export default LoginComponent;