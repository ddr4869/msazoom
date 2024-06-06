import boardStyles from '@/styles/board-styles.module.css';

const FollowerCard = ({ follower, addFriend }) => (
  <div className={boardStyles.boardCard}>
    <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
      <h2 style={{ marginRight: 'auto', letterSpacing: '2px' }}>
        {follower.username}
      </h2>
      <button 
        style={{ marginRight: '10px', padding: '12px 24px', fontSize: '16px' }}
        onClick={() => addFriend(follower.username)}>
          Add
      </button>
    </div>
  </div>
);
export default FollowerCard;
