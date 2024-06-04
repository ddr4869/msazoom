import React from 'react';
import ReactDOM from 'react-dom';

const ChatModal = ({ onClose }) => {
  const handleSendMessage = () => {
    // 메시지 전송 기능 구현
  };

  return ReactDOM.createPortal(
    <div className="modal">
      <div className="modal-content">
        <span className="close" onClick={onClose}>&times;</span>
        <h2>Chat with Partner</h2>
        <textarea placeholder="Type your message here..."></textarea>
        <button onClick={handleSendMessage}>Send</button>
      </div>
    </div>,
    document.getElementById('modal-root') // 모달이 렌더링될 DOM 노드 지정
  );
}

export default ChatModal;
