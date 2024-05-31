import BoardCard from './BoardCard';
import boardStyles from '@/styles/board-styles.module.css';

const BoardList = ({ boards, navigateToBoard, handleRecommendBoard, handleDeleteBoard }) => (
  <div className={boardStyles.gridBoardContainer}>
    {boards.map((board) => (
      <BoardCard
        key={board.id}
        board={board}
        navigateToBoard={navigateToBoard}
        handleRecommendBoard={handleRecommendBoard}
        handleDeleteBoard={handleDeleteBoard}
      />
    ))}
  </div>
);

export default BoardList;
