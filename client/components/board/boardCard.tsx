import boardStyles from '@/styles/board-styles.module.css';

const BoardCard = ({ board, navigateToBoard, handleRecommendBoard, handleDeleteBoard }) => (
  <div className={boardStyles.boardCard}>
    <div><h2>{board.board_name}</h2></div>
    <div>Star: {board.board_star}</div>
    <div>Admin: {board.board_admin}</div>
    <button onClick={() => navigateToBoard(board.id, board.board_name)}>View Board</button>{" "}
    <button onClick={() => handleRecommendBoard(board.id)}>Recommend</button>{" "}
    {board.board_admin == localStorage.getItem('username') &&
      <button onClick={() => handleDeleteBoard(board.id, board.board_password)}>Delete</button>
    }
  </div>
);

export default BoardCard;
