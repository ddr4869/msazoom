import React from 'react';
import ChatCard from './chatCard';
import boardStyles from '@/styles/board-styles.module.css';

export interface Chat {
  id: number;
  title: string;
  admin: string;
  created_at: string;
}

interface ChatListProps {
  chats: Chat[];
  navigateToChat: (id: number) => void;
}

const ChatList: React.FC<ChatListProps> = ({ chats, navigateToChat }) => (
  <div className={boardStyles.gridBoardContainer}>
    {chats.length === 0 && 
      <h2>진행중인 미팅이 없어요.</h2>
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
