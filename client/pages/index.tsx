import { useState, useEffect, FormEvent } from 'react';
import { useRouter } from "next/router";
import { createChatAxios, getChatsAxios, getRandomChatIdAxios, getChatAxios } from '@/server/chat';
import CreateChatForm from '@/ui/chat/createChatForm';
import { handleLogin, handleLogout } from '@/utils/auth';
import ChatList from '@/components/chat/chatList';
import LoginComponent from '@/components/user/loginComponent';
import SignUpComponent from '@/components/user/singupComponent';
import { AddFriendAxios, GetFollowerAxios, GetFriendsAxios, RemoveFriendAxios } from '@/server/user';
import FriendsList from '@/components/user/friendList';
import FollowerList from '@/components/user/followerList';

const Home = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [chats, setChats] = useState([]);
  const [friends, setFriends] = useState([]);
  const [followers, setFollowers] = useState([]);
  const [chatReload, setChatReload] = useState(false);
  const [showCreateChatForm, setShowCreateChatForm] = useState(false);
  const router = useRouter();
  
  const handleCreateChatClick = () => {
    setShowCreateChatForm(true);
  };

  const handleSubmitChatForm = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    try {
      const formData = new FormData(event.currentTarget);
      const chat_title = formData.get('title') as string;
      const response = await createChatAxios(localStorage.getItem('accessToken'), localStorage.getItem('username'), chat_title);
      console.log("handleSubmitChatForm response -> ", response)
      setShowCreateChatForm(false);
      setChatReload(true);
      router.push({
        pathname: `/chat/${response.id}`,
      });
    } catch (error) {
      console.error('Error creating chat:', error);
    }
  };


  const navigateToChat = (chatId: any) => {
    getChatAxios(localStorage.getItem('accessToken'), chatId).then((response) => {
      console.log("response -> ", response)
      router.push({
        pathname: `/chat/${chatId}`,
      });
    }).catch((error) => {
      alert('이미 종료된 채팅입니다.');
      setChatReload(true);
    })
  };

  const navigateToRandomChat = () => {
    getRandomChatIdAxios(localStorage.getItem('accessToken')).then((response) => {
      console.log("response -> ", response)
      navigateToChat(response);
    }).catch((error) => {
      console.error('Error fetching random chat:', error);
      alert('No chat available');
    });
  }

  const navigateToFriendChat = (friendId: any) => {
    console.log("friendId -> ", friendId)
    router.push({
      pathname: `/friend`,
      query: { friend: friendId }
    });
  }

  const removeFriend = (friendId: any) => {
    if (confirm('친구목록에서 삭제하겠습니까?')) {
      RemoveFriendAxios(localStorage.getItem('accessToken'), friendId).then(() => {
        console.log("remove friend success")
        setChatReload(true);
    }).catch((error) => {
      console.error('Error removing friend:', error);
      })
    }
  }

  const addFriend = (friendId: any) => {
    if (confirm('친구요청을 수락하겠습니까?')) {
      AddFriendAxios(localStorage.getItem('accessToken'), friendId).then(() => {
        console.log("add friend success")
        setChatReload(true);
    }).catch((error) => {
      console.error('Error removing friend:', error);
      })
    }
  }

  useEffect(() => {
    const accessToken = localStorage.getItem('accessToken');
    setIsLoggedIn(!!accessToken);
  }, []);

  useEffect(() => {
    const fetchChats = async (token: string) => {
      try {
        const response = await getChatsAxios(token);
        setChats(response);
      } catch (error) {
        console.error('Error fetching chats:', error);
      }
    };

    const fetchFriends = async (token: string) => {
      try {
        const response = await GetFriendsAxios(token);
        if (response !== null) {
          setFriends(response);
        }
      } catch (error) {
        console.error('Error fetching friends:', error);
      }
    }

    const fetchFollower = async (token: string) => {
      try {
        const response = await GetFollowerAxios(token);
        if (response !== null) {
          setFollowers(response);
        }
      } catch (error) {
        console.error('Error fetching friends:', error);
      }
    }

    const accessToken = localStorage.getItem('accessToken');
    if (isLoggedIn && accessToken) {
      fetchChats(accessToken);
      fetchFriends(accessToken);
      fetchFollower(accessToken); 
    } else {
      setChats([]);
      setFriends([]);
    }
    setChatReload(false);
  }, [isLoggedIn, chatReload]);

  return (
    <div className="chat-board">
      <header>
        <LoginComponent 
          username={username} 
          setUsername={setUsername} 
          password={password} 
          setPassword={setPassword} 
          handleLogin={(event) => handleLogin(username, password, setIsLoggedIn, event)}
          handleLogout={() => handleLogout(setIsLoggedIn)}
          isLoggedIn={isLoggedIn} 
        />
      </header>
      <main>
        <br></br>
 
          {isLoggedIn ? <h1>Chat List</h1> : <h2>Sign up for an account if you don't have one.</h2> } 
        
        {/* <br></br><br></br> */}
        {!isLoggedIn && ( <SignUpComponent/> )}
        { isLoggedIn && (
          <div>
            <button onClick={() => setChatReload(true) }>Reload Chat List</button> {  }
            <button onClick={() => navigateToRandomChat() }>Random Chat Start</button> 
          </div>
        )}
        <br></br>
        <br></br>
        { isLoggedIn && (
          <ChatList
          chats={chats}
          navigateToChat={navigateToChat}
        />
        )}
        <br></br>
        {isLoggedIn && !showCreateChatForm && (
          <button onClick={handleCreateChatClick}>Create Chat</button>
        )}
        {isLoggedIn && showCreateChatForm && (
          <CreateChatForm
            handleSubmitChatForm={handleSubmitChatForm}
            setShowCreateChatForm={setShowCreateChatForm}
          />
          

        )}
        {isLoggedIn &&
          <div>
          <br></br><br></br>
          <hr></hr>
          </div>
        }
        {isLoggedIn && <h1>Friends List</h1>}
        {isLoggedIn && (          
            <FriendsList 
              friends={friends} 
              navigateToFriendChat={navigateToFriendChat} 
              removeFriend={removeFriend}
            />
          )
        }
        {isLoggedIn &&
          <div>
          <br></br><br></br>
          <hr></hr>
          </div>
        }
        {isLoggedIn && <h1>Following Request</h1>}
        {isLoggedIn && (          
            <FollowerList 
              followers={followers} 
              addFriend={addFriend} 
            />
          )
        }
      </main>
    </div>
  );
};

export default Home;
