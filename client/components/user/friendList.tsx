import boardStyles from '@/styles/board-styles.module.css';
import FriendCard from './friendCard';
interface FriendListProps {
  friends: { id: number; username: string }[];
  navigateToFriendChat: (id: number) => void;
  removeFriend: (id: number) => void;
}
const FriendsList = ({friends, navigateToFriendChat, removeFriend}:FriendListProps) => (
    <div className={boardStyles.friendList}>
    {friends.length === 0 && 
      <h2> 친구가 없어요 ㅠ.</h2>
    }
    {friends.map((friend) => (
      <FriendCard
        key={friend.id}
        friend={friend}
        navigateToFriendChat={navigateToFriendChat}
        removeFriend={removeFriend}
      />
    ))}
  </div>
);
export default FriendsList;