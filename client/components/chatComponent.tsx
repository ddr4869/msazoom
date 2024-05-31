import MsgComponent from "./msgComponent";
import { useRef, useEffect } from "react";

const ChatComponent = ({ board_name, board_id, messages, newMessage, me, value, setValue, onSubmit }) => (
  <div className="container">
    <div className="header">
      <h1>{board_name}</h1>
    </div>
    <div className="messages">
      {messages.length === 0 && newMessage.length === 0 ? (
        <div>No Messages</div>
      ) : (
        <div>
          {/* 기존 DB Message */}
          {messages.slice().reverse().map((msg, index) => (
            <MsgComponent key={index} board_id={board_id}  writer={msg.writer} message={msg.message} me={me} />
          ))}
        </div>
      )}
      {newMessage.length > 0 ? (
        <div>
          {/* 새 메세지 */}
          {newMessage.map((msg, index) => {
            if (msg.board_id === board_id) {
              return <MsgComponent key={index} writer={msg.writer} message={msg.message} me={me} board_id={msg.board_id} />
            } 
          })}
        </div>
      ) : null}
    </div>
    <div className="input-area">
      <input
        type="text"
        value={value}
        onChange={(e) => setValue(e.target.value)}
        onKeyPress={(e) => {
          if (e.key === "Enter") {
            onSubmit();
          }
        }}
        placeholder="Type your message here..."
      />
      <button onClick={onSubmit}>Send</button>
    </div>
  </div>
);

export default ChatComponent;