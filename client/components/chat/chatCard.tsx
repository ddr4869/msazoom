import boardStyles from '@/styles/board-styles.module.css';

const ChatCard = ({ chat, navigateToChat }) => (
  <div className={boardStyles.boardCard}>
    <div><h2>{chat.title}</h2></div>
    <div>Admin: {chat.admin}</div>
    <div>Created: {chat.created_at}</div>
    <button onClick={() => navigateToChat(chat.id)}>Enter Chat</button>{" "}
  </div>
);


export default ChatCard;
