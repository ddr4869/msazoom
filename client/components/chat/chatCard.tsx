import boardStyles from '@/styles/board-styles.module.css';

const ChatCard = ({ chat, navigateToChat }) => (
  <div className={boardStyles.boardCard}>
    <div><h2>{chat.chat_name}</h2></div>
    <div>Admin: {chat.chat_user}</div>
    <button onClick={() => navigateToChat(chat.id)}>Enter Chat</button>{" "}
  </div>
);


export default ChatCard;
