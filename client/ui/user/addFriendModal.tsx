import React, { useEffect, useState } from 'react';
import ReactDOM from 'react-dom';
import styles from '@/styles/user-styles.module.css';

const AddFriendModal = ({ onClose }) => {
  const [isClient, setIsClient] = useState(false);

  useEffect(() => {
    setIsClient(true);
  }, []);

  const handleAddFriend = () => {
    // 친구 추가 기능 구현
    // 추가한 후에는 모달을 닫습니다.
    onClose();
  };

  if (!isClient) {
    return null;
  }

  return ReactDOM.createPortal(
    <div className={styles['modal-overlay']}>
      <div className={styles.modal}>
        <div className={styles['modal-content']}>
          <span className={styles.close} onClick={onClose}>&times;</span>
          <h2>Add Friend</h2>
          <button onClick={handleAddFriend}>Add Friend</button>
        </div>
      </div>
    </div>,
    document.getElementById('modal-root')
  );
}

export default AddFriendModal;
