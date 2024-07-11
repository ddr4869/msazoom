import boardStyles from '@/styles/board-styles.module.css';
interface FriendCardProps {
  friend: { id: number; username: string, Message: { unread_message_count: number }};
  navigateToFriendChat: (username: string) => void;
  removeFriend: (username: string) => void;
}
const FriendCard = ({ friend, navigateToFriendChat, removeFriend }:FriendCardProps) => (
  <div className={boardStyles.boardCard}>
    <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
      <h2 style={{ marginRight: 'auto', letterSpacing: '2px' }}>
        {friend.username}
      </h2>
      <span style={{
        fontSize: '16px',
        fontWeight: 'bold',
        color: 'white',
        backgroundColor: '#007bff',
        padding: '4px 8px',
        borderRadius: '50%',
        margin: "10px"
      }}>
        {friend.Message.unread_message_count}

      </span>
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
