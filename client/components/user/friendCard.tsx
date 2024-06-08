import boardStyles from '@/styles/board-styles.module.css';
interface FriendCardProps {
  friend: { id: number; username: string };
  navigateToFriendChat: (username: string) => void;
  removeFriend: (username: string) => void;
}
const FriendCard = ({ friend, navigateToFriendChat, removeFriend }:FriendCardProps) => (
  <div className={boardStyles.boardCard}>
    <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
      <h2 style={{ marginRight: 'auto', letterSpacing: '2px' }}>
        {friend.username}
      </h2>
      <button 
        style={{ marginRight: '10px', padding: '12px 24px', fontSize: '16px' }}
        onClick={() => navigateToFriendChat(friend.username)}>
          Chat
      </button>
      <button 
        style={{ padding: '12px 24px', fontSize: '16px' }}
        onClick={() => removeFriend(friend.username)}>
          Remove
      </button>
    </div>
  </div>
);
export default FriendCard;
