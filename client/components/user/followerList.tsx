import boardStyles from '@/styles/board-styles.module.css';
import FriendCard from './friendCard';
import FollowerCard from './followerCard';

const FollowerList = ({followers, addFriend}) => (
    <div className={boardStyles.friendList}>
    {followers.length === 0 && 
      <h2> 대기 중인 친구 요청이 없습니다.</h2>
    }
    {followers.map((follower) => (
      <FollowerCard
        key={follower.id}
        follower={follower}
        addFriend={addFriend}
      />
    ))}
  </div>
);
export default FollowerList;