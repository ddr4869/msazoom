import { useState, useEffect, FormEvent } from 'react';
import { useRouter } from "next/router";
import { createChatAxios, getChatsAxios } from '@/server/chat';
import CreateChatForm from '@/ui/chat/createChatForm';
import { handleLogin, handleLogout } from '@/utils/auth';
import ChatList from '@/components/chat/chatList';
import LoginComponent from '@/components/user/loginComponent';
import SignUpComponent from '@/components/user/singupComponent';
import { GetFriendsAxios, RemoveFriendAxios } from '@/server/user';
import FriendsList from '@/components/user/friendList';

const Home = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [chats, setChats] = useState([]);
  const [friends, setFriends] = useState([]);
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
      const response = await createChatAxios(localStorage.getItem('username'), chat_title);
      console.log("response -> ", response)
      setShowCreateChatForm(false);
      setChatReload(true);
      navigateToChat(response.id);
    } catch (error) {
      console.error('Error creating chat:', error);
    }
  };


  const navigateToChat = (chatId: any) => {
    console.log("chatId -> ", chatId)
    router.push({
      pathname: `/chat/${chatId}`,
    });
  };

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
        setFriends(response);
      } catch (error) {
        console.error('Error fetching friends:', error);
      }
    }

    const accessToken = localStorage.getItem('accessToken');
    if (isLoggedIn && accessToken) {
      fetchChats(accessToken);
      fetchFriends(accessToken);
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
        { isLoggedIn && <button onClick={() => setChatReload(true) }>Reload Chat List</button> }
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
        <br></br><br></br>
        <hr></hr>

        {isLoggedIn && <h1>Friends List</h1>}
        {isLoggedIn && (          
            <FriendsList 
              friends={friends} 
              navigateToFriendChat={navigateToFriendChat} 
              removeFriend={removeFriend}
            />
          )
        }
      </main>
    </div>
  );
};

export default Home;
