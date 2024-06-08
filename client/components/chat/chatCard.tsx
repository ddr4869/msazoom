import boardStyles from '@/styles/board-styles.module.css';
export interface Chat {
  id: number;
  title: string;
  admin: string;
  created_at: string;
}

interface ChatCardProps {
  chat: Chat;
  navigateToChat: (id: number) => void;
}

const ChatCard = (props:ChatCardProps) => (
  <div className={boardStyles.boardCard}>
    <div><h2>{props.chat.title}</h2></div>
    <div>Admin: {props.chat.admin}</div>
    <div>Created: {props.chat.created_at}</div>
    <button onClick={() => props.navigateToChat(props.chat.id)}>Enter Chat</button>{" "}
  </div>
);


export default ChatCard;
