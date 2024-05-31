import userStyles from '@/styles/userProfile-styles.module.css';

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

export default UserProfile;
