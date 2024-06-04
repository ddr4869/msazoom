import boardStyles from '@/styles/board-styles.module.css';
import FriendCard from './friendCart';

const FriendsList = ({friends}) => (
    <div className={boardStyles.friendList}>
    {friends.map((friend) => (
      <FriendCard
        //key={friend.id}
        friend={friend}
      />
    ))}
  </div>
);
export default FriendsList;