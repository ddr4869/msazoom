import boardStyles from '@/styles/board-styles.module.css';

const FriendCard = ({ friend }) => (
  <div className={boardStyles.boardCard}>
    <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
      <h2 style={{ marginRight: 'auto', letterSpacing: '2px' }}>{friend.username}</h2>
      <button style={{ marginRight: '10px', padding: '12px 24px', fontSize: '16px' }}>Chat</button>
      <button style={{ padding: '12px 24px', fontSize: '16px' }}>Remove</button>
    </div>
  </div>
);

export default FriendCard;
