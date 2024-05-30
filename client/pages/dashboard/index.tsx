// pages/index.js
import { useState, useEffect, FormEvent } from 'react';
import axios from 'axios';
import { signIn, signOut, useSession } from 'next-auth/react';
import { useRouter } from "next/router";
import { createBoardAxios, getBoardsAxios, recommendBoardAxios, deleteBoardAxios } from '../server/board';
import boardStyles from '../styles/board-styles.module.css'
import userStyles from '../styles/userProfile-styles.module.css'

const Home = () => {
  const { data: session } = useSession();
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [boards, setBoards] = useState([]);
  const [boardReload, setBoardReload] = useState(false);
  const [showCreateBoardForm, setShowCreateBoardForm] = useState(false);
  const router = useRouter();

  const handleCreateBoardClick = () => {
    console.log("setShowCreateBoardForm True")
    setShowCreateBoardForm(true); // Create Board 버튼 클릭 시 form 보이도록 상태 변경
  };

  const handleSubmitBoardForm = async (event: FormEvent<HTMLFormElement>) => {
    try {
      event.preventDefault();
      const formData = new FormData(event.currentTarget);
      const boardName = formData.get('board_name') as string;
      const boardPassword = formData.get('board_password') as string;
      await createBoardAxios(localStorage.getItem('accessToken'), boardName, boardPassword);
      setShowCreateBoardForm(false); // form 제출 시 form 숨기도록 상태 변경
    } catch (error) {
      console.error('Error creating board:', error);
    }
  };


  const navigateToBoard = (boardId:any, boardName:any) => {
    router.push({
      pathname: `/board/${boardId}`,
      query: {
        board_name: boardName
      }
    });
  };

  async function recommendBoard(board_id:number) {
      try {
        // API 요청 URL을 적절히 변경하세요.
        console.log("board_id: ", board_id)
        const response = await recommendBoardAxios(board_id);
        console.log("resp: ", response)
        return response
      } catch (error) {
        console.error('Error fetching boards:', error);
      }
  };

  const handleRecommendBoard = async (board_id: number) => {
    try {
        await recommendBoard(board_id); // 추천 요청 보내기
        setBoardReload(true); // 추천 상태 업데이트
    } catch (error) {
        // 추천 요청 실패 시 에러 처리
        console.error('Error recommending board:', error);
    }
  };

  async function deleteBoard(board_id:number, board_password:string) {
    try {
      // API 요청 URL을 적절히 변경하세요.
      console.log("board_id: ", board_id)
      const response = await deleteBoardAxios(localStorage.getItem('accessToken'), board_id, board_password);
      console.log("resp: ", response)
      return response
    } catch (error) {
      console.error('Error fetching boards:', error);
    }
};

  const handleDeleteBoard = async (board_id: number, board_password:string) => {
    try {
        await deleteBoard(board_id, board_password); // 추천 요청 보내기
        setBoardReload(true); // 추천 상태 업데이트
    } catch (error) {
        // 추천 요청 실패 시 에러 처리
        console.error('Error recommending board:', error);
    }
  };

  useEffect(() => {
    if (typeof localStorage === 'undefined') {
        localStorage= {};
        return;
    }
    const accessToken = localStorage.getItem('accessToken');
    setIsLoggedIn(!!accessToken);
  }, []);

  useEffect(() => {
    const fetchBoards = async (token:string) => {
      try {
        // API 요청 URL을 적절히 변경하세요.
        const response = await getBoardsAxios(token);
        setBoards(response); // 'data' 필드 안의 'data' 배열을 상태로 설정
      } catch (error) {
        console.error('Error fetching boards:', error);
      }
    };
    const accessToken = localStorage.getItem('accessToken');
    if (isLoggedIn && accessToken) {
      fetchBoards(accessToken); // fetchData는 보드 정보를 불러오는 함수입니다.
    } else {
      setBoards([])
    }
    setBoardReload(false);
  }, [isLoggedIn, boardReload, showCreateBoardForm]);

  const handleLogin = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    // Here you would replace with your own signIn logic
    const response = await axios.post('http://localhost:8080/api/user/login', {
        username,
        password
      });
      const { data } = response.data;
      const { access_token } = data;
      localStorage.setItem('accessToken', access_token);
      localStorage.setItem('username', username);
      setIsLoggedIn(true);
  };

  const handleLogout = async () => {
    localStorage.removeItem('accessToken');
    localStorage.removeItem('username');
    setIsLoggedIn(false);
  };

  return (
    <div className="chat-board">
      <header>
        <h1>Chat Board</h1>
        <div id="login-section">
          {!isLoggedIn ? (
            <form onSubmit={handleLogin}>
              <input type="text" value={username} onChange={(e) => setUsername(e.target.value)} placeholder="Username" required />
              <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} placeholder="Password" required />
              <button type="submit">Login</button>
            </form>
          ) : (
            <div className={userStyles.gridUserContainer}>
              {/* 사용자 프로필 */}
              <div className={`${userStyles.userStyles} ${userStyles.userProfile}`}>
                <p>Welcome, <strong>{localStorage.getItem('username')}</strong></p>
                <button onClick={handleLogout}>Logout</button>
              </div>
              {/* 추가 정보 */}
              <div className={`${userStyles.userStyles} ${userStyles.additionalInfo}`}>
                <p>Email: example@example.com</p>
                <p>Role: Administrator</p>
              </div>
            </div>
          )}
        </div>
      </header>
      <main>
        <div>
        <h2> 
          {isLoggedIn ? "Container List" : "Please Login" } 
        </h2>
          <ul>
            <div className={boardStyles.gridBoardContainer}>
              {boards.map((board) => (
                <div key={board.id} className={boardStyles.boardCard}>
                {/* <div>ID: {board.board_id}</div> */}
                <div><h2>{board.board_name}</h2></div>
                <div>Star: {board.board_star}</div>
                <div>Admin: {board.board_admin}</div>
                <button onClick={() => navigateToBoard(board.id, board.board_name)}>View Board</button>{" "}
                <button onClick={() => handleRecommendBoard(board.id)}>Recommend</button>{" "}
                { board.board_admin == localStorage.getItem('username') &&
                  <button onClick={() => handleDeleteBoard(board.id, board.board_password)}>Delete</button>
                }
              </div>
              ))}
            </div>
          </ul>
          <div>
          {/* // Add a button to create a new board */}
          { isLoggedIn && !showCreateBoardForm && (
            <button onClick={handleCreateBoardClick}>Create Board</button>
          )}
          {/* 새로운 board 생성 form */}
          { isLoggedIn && showCreateBoardForm && (
            <form onSubmit={handleSubmitBoardForm}>
              {/* 필요한 입력 필드 추가 */}
              <input type="text" placeholder="Board Name" name="board_name" />
              <input type="text" placeholder="Board Password" name="board_password" />
              <button type="submit">Create</button>
            <form onClick={() => setShowCreateBoardForm(false)}>
              <button type="submit">Cancel</button>
            </form>
            </form>
          )}
        </div>
        </div>

      </main>
    </div>
  );
};

export default Home;
