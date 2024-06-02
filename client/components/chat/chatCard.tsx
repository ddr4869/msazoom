import boardStyles from '@/styles/board-styles.module.css';

const ChatCard = ({ chat, navigateToChat }) => (
  <div className={boardStyles.boardCard}>
    <div><h2>{chat.board_name}</h2></div>
    <div>Star: {chat.board_star}</div>
    <div>Admin: {chat.board_admin}</div>
    <button onClick={() => navigateToChat(chat.id, chat.board_name)}>View Chat</button>{" "}
  </div>
);

export default ChatCard;
