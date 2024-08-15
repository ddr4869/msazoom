import { useState, useEffect, FormEvent } from 'react';
import { useRouter } from "next/router";
import { createChatAxios, getChatsAxios, getRandomChatIdAxios, getChatAxios, checkChatPasswordAxios } from '@/server/chat';
import CreateChatForm from '@/ui/chat/CreateChatForm';
import { handleLogin, handleLogout } from '@/utils/auth';
import ChatList from '@/components/chat/chatList';
import LoginComponent from '@/components/user/LoginComponent';
import SignUpComponent from '@/components/user/singupComponent';
import { AddFriendAxios, GetFollowerAxios, GetFriendsAxios, RemoveFriendAxios } from '@/server/user';
import FriendsList from '@/components/user/friendList';
import FollowerList from '@/components/user/followerList';

const Home = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [chats, setChats] = useState([]);
  const [friends, setFriends] = useState<any>([]);
  const [followers, setFollowers] = useState([]);
  const [chatReload, setChatReload] = useState(false);
  const [showCreateChatForm, setShowCreateChatForm] = useState(false);
  const [accessToken, setAccessToken] = useState<string>('');
  const router = useRouter();
  const handleCreateChatClick = () => {
    setShowCreateChatForm(true);
  };

  const handleSubmitChatForm = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    try {
      const formData = new FormData(event.currentTarget);
      const chat_title = formData.get('title') as string;
      const password = formData.get('password') ? formData.get('password') as string : '';
      const response = await createChatAxios(username, chat_title, password);
      setShowCreateChatForm(false);
      setChatReload(true);
      router.push({
        pathname: `/chat/${response.id}`,
        query: { password: password } 
      });
    } catch (error) {
      console.error('Error creating chat:', error);
    }
  };


  const navigateToChat = (chatId: any) => {
    getChatAxios(chatId).then((response) => {
      console.log("response -> ", response)
      if (response.Private) {
        const password = prompt('Password Required');
        if (password) {
          checkChatPasswordAxios(chatId, password).then(() => {
            router.push({ pathname: `/chat/${chatId}` });
          }).catch((error) => {
            alert('Invalid Password.');
          })
        }
          return;
        }
        router.push({ pathname: `/chat/${chatId}` });
    }).catch((error) => {
      alert('이미 종료된 채팅입니다.');
      setChatReload(true);
    })
  };

  const navigateToRandomChat = () => {
    getRandomChatIdAxios().then((response) => {
      console.log("response -> ", response)
      navigateToChat(response);
    }).catch((error) => {
      console.error('Error fetching random chat:', error);
      alert('No chat available');
    });
  }

  const navigateToFriendChat = (friendId: string) => {
    console.log("friendId -> ", friendId)
    router.push({
      pathname: `/friend`,
      query: { friend: friendId }
    });
  }

  const removeFriend = (friendId: string) => {
    if (confirm('친구목록에서 삭제하겠습니까?')) {
      RemoveFriendAxios(friendId).then(() => {
        console.log("remove friend success")
        setChatReload(true);
    }).catch((error) => {
      console.error('Error removing friend:', error);
      })
    }
  }

  const addFriend = (friendId: any) => {
    if (confirm('친구요청을 수락하겠습니까?')) {
      AddFriendAxios(friendId).then(() => {
        console.log("add friend success")
        setChatReload(true);
    }).catch((error) => {
      console.error('Error removing friend:', error);
      })
    }
  }

  useEffect(() => {
    const token = localStorage.getItem('accessToken') || '';
    const userId = localStorage.getItem('username') || '';
    setAccessToken(token);
    setUsername(userId);
    setIsLoggedIn(!!token);
  }, [router]);

  useEffect(() => {
    const fetchChats = async () => {
      try {
        const response = await getChatsAxios();
        setChats(response);
      } catch (error) {
        console.error('Error fetching chats:', error);
      }
    };

    const fetchFriends = async () => {
      try {
        const response = await GetFriendsAxios();
        if (response !== null) {
          setFriends(response);
        }
      } catch (error) {
        console.error('Error fetching friends:', error);
      }
    }

    const fetchFollower = async () => {
      try {
        const response = await GetFollowerAxios();
        if (response !== null) {
          setFollowers(response);
        }
      } catch (error) {
        console.error('Error fetching friends:', error);
      }
    }

    const accessToken = localStorage.getItem('accessToken');
    if (isLoggedIn && accessToken) {
      fetchChats();
      fetchFriends();
      fetchFollower(); 
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
          handleLogin={(e) => handleLogin(username, password, setIsLoggedIn, e)}
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
              key={friends.id}
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
