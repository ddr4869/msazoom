import ChatCard from './chatCard';
import boardStyles from '@/styles/board-styles.module.css';

const ChatList = ({ chats, navigateToChat }) => (
  <div className={boardStyles.gridBoardContainer}>
    {chats.length === 0 && 
    // apply css in h2 tag
    // if there is no chat, display this message
    
      <h2> 진행중인 미팅이 없어요.</h2>
    }
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
