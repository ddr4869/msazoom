import React from 'react';
import BoardCard from './boardCard';
import boardStyles from '@/styles/board-styles.module.css';

interface Board {
  id: number;
  board_name: string;
  board_star: number;
  board_admin: string;
  board_password: string;
}

interface BoardListProps {
  boards: Board[];
  navigateToBoard: (id: number, name: string) => void;
  handleRecommendBoard: (id: number) => void;
  handleDeleteBoard: (id: number, password: string) => void;
}

const BoardList: React.FC<BoardListProps> = ({ boards, navigateToBoard, handleRecommendBoard, handleDeleteBoard }) => (
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
