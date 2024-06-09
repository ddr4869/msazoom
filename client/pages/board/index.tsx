import { useState, useEffect, FormEvent } from 'react';
import { useRouter } from "next/router";
import { createBoardAxios, getBoardsAxios, recommendBoardAxios, deleteBoardAxios } from '@/server/board';
import BoardList from '@/components/board/boardList';
import CreateBoardForm from '@/ui/board/createBoardForm';
import LoginForm from '@/components/user/loginForm';
import { handleLogin } from '@/utils/auth';

const Home = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [boards, setBoards] = useState([]);
  const [boardReload, setBoardReload] = useState(false);
  const [showCreateBoardForm, setShowCreateBoardForm] = useState(false);
  const [accessToken, setAccessToken] = useState<string>('');
  const router = useRouter();

  const handleCreateBoardClick = () => {
    setShowCreateBoardForm(true);
  };

  const handleSubmitBoardForm = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    try {
      const formData = new FormData(event.currentTarget);
      const boardName = formData.get('board_name') as string;
      const boardPassword = formData.get('board_password') as string;
      await createBoardAxios(boardName, boardPassword);
      setShowCreateBoardForm(false);
      setBoardReload(true);
    } catch (error) {
      console.error('Error creating board:', error);
    }
  };

  const navigateToBoard = (boardId: any, boardName: any) => {
    router.push({
      pathname: `/board/${boardId}`,
      query: { board_name: boardName }
    });
  };

  const handleRecommendBoard = async (board_id: number) => {
    try {
      await recommendBoardAxios(board_id);
      setBoardReload(true);
    } catch (error) {
      console.error('Error recommending board:', error);
    }
  };

  const handleDeleteBoard = async (board_id: number, board_password: string) => {
    try {
      await deleteBoardAxios(board_id, board_password);
      setBoardReload(true);
    } catch (error) {
      console.error('Error deleting board:', error);
    }
  };

  useEffect(() => {
    const token = localStorage.getItem('accessToken') || '';
    setAccessToken(token);
    setIsLoggedIn(!!accessToken);
  }, []);

  useEffect(() => {
    const fetchBoards = async () => {
      try {
        const response = await getBoardsAxios();
        setBoards(response);
      } catch (error) {
        console.error('Error fetching boards:', error);
      }
    };

    const accessToken = localStorage.getItem('accessToken');
    if (isLoggedIn && accessToken) {
      fetchBoards();
    } else {
      setBoards([]);
    }
    setBoardReload(false);
  }, [isLoggedIn, boardReload, showCreateBoardForm]);

  return (
    <div className="chat-board">
      <header>
        <h1>Chat Board</h1>
        <div id="login-section">
          {!isLoggedIn && (
            <LoginForm
              username={username}
              setUsername={setUsername}
              password={password}
              setPassword={setPassword}
              handleLogin={(event) => handleLogin(username, password, setIsLoggedIn, event)}
            />
          ) 
        }
        </div>
      </header>
      <main>
        <h2> 
          {isLoggedIn ? "Container List" : "Please Login" } 
        </h2>
        <BoardList
          boards={boards}
          navigateToBoard={navigateToBoard}
          handleRecommendBoard={handleRecommendBoard}
          handleDeleteBoard={handleDeleteBoard}
        />
        {isLoggedIn && !showCreateBoardForm && (
          <button onClick={handleCreateBoardClick}>Create Board</button>
        )}
        {isLoggedIn && showCreateBoardForm && (
          <CreateBoardForm
            handleSubmitBoardForm={handleSubmitBoardForm}
            setShowCreateBoardForm={setShowCreateBoardForm}
          />
        )}
      </main>
    </div>
  );
};

export default Home;