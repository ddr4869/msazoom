import { useState, useEffect, FormEvent } from 'react';
import { useRouter } from "next/router";
import { createChatAxios, getChatsAxios } from '@/server/chat';
import CreateChatForm from '@/ui/chat/CreateChatForm';
import { handleLogin, handleLogout } from '@/utils/auth';
import ChatList from '@/components/chat/chatList';
import LoginComponent from '@/components/user/loginComponent';

const Home = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [chats, setChats] = useState([]);
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

  // TODO !!
  const navigateToChat = (chatId: any) => {
    console.log("chatId -> ", chatId)
    router.push({
      pathname: `/chat/${chatId}`,
      //query: { board_name: boardName }
    });
  };

  useEffect(() => {
    const accessToken = localStorage.getItem('accessToken');
    setIsLoggedIn(!!accessToken);
  }, []);

  useEffect(() => {
    const fetchBoards = async (token: string) => {
      try {
        const response = await getChatsAxios(token);
        setChats(response);
      } catch (error) {
        console.error('Error fetching boards:', error);
      }
    };

    const accessToken = localStorage.getItem('accessToken');
    if (isLoggedIn && accessToken) {
      fetchBoards(accessToken);
    } else {
      setChats([]);
    }
    setChatReload(false);
  }, [isLoggedIn, chatReload, showCreateChatForm]);

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
        <h2> 
          {isLoggedIn ? "Container List" : "Please Login" } 
        </h2>
        <ChatList
          chats={chats}
          navigateToChat={navigateToChat}
        />
        {isLoggedIn && !showCreateChatForm && (
          <button onClick={handleCreateChatClick}>Create Chat</button>
        )}
        {isLoggedIn && showCreateChatForm && (
          <CreateChatForm
            handleSubmitChatForm={handleSubmitChatForm}
            setShowCreateChatForm={setShowCreateChatForm}
          />
        )}
      </main>
    </div>
  );
};

export default Home;