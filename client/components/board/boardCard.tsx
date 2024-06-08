import boardStyles from '@/styles/board-styles.module.css';
interface Board {
  id: number;
  board_name: string;
  board_star: number;
  board_admin: string;
  board_password: string;
}

interface BoardCardProps {
  board: Board;
  navigateToBoard: (id: number, name: string) => void;
  handleRecommendBoard: (id: number) => void;
  handleDeleteBoard: (id: number, password: string) => void;
}

const BoardCard = (props:BoardCardProps) => (
  <div className={boardStyles.boardCard}>
    <div><h2>{props.board.board_name}</h2></div>
    <div>Star: {props.board.board_star}</div>
    <div>Admin: {props.board.board_admin}</div>
    <button onClick={() => props.navigateToBoard(props.board.id, props.board.board_name)}>View Board</button>{" "}
    <button onClick={() => props.handleRecommendBoard(props.board.id)}>Recommend</button>{" "}
    {props.board.board_admin == localStorage.getItem('username') &&
      <button onClick={()   => props.handleDeleteBoard(props.board.id, props.board.board_password)}>Delete</button>
    }
  </div>
);
export default BoardCard;
