import React from 'react';
import ReactDOM from 'react-dom';

interface ChatModalProps {
  onClose: () => void;
}

const ChatModal: React.FC<ChatModalProps> = ({ onClose }) => {
  const handleSendMessage = () => {
    // 메시지 전송 기능 구현
  };

  const modalRoot = document.getElementById('modal-root');
  if (!modalRoot) {
    console.error('The element with id "modal-root" was not found.');
    return null;
  }

  return ReactDOM.createPortal(
    <div className="modal">
      <div className="modal-content">
        <span className="close" onClick={onClose}>&times;</span>
        <h2>Chat with Partner</h2>
        <textarea placeholder="Type your message here..."></textarea>
        <button onClick={handleSendMessage}>Send</button>
      </div>
    </div>,
    modalRoot
  );
};

export default ChatModal;
