import ChatCard from './chatCard';
import boardStyles from '@/styles/board-styles.module.css';

const ChatList = ({ chats, navigateToChat }) => (
  <div className={boardStyles.gridBoardContainer}>
    {chats.map((chat) => (
      <ChatCard
        key={chat.id}
        chat={chat}
        navigateToChat={navigateToChat}
      />
    ))}
  </div>
);

export default ChatList;
