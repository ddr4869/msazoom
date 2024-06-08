import MsgComponent from "./msgComponent";
import { useRef, useEffect } from "react";

interface ChatComponentProps {
  board_name: string;
  board_id: string;
  messages: { writer: string; message: string }[];
  newMessage: { board_id: string; writer: string; message: string }[];
  me: string;
  value: string;
  setValue: (value: string) => void;
  onSubmit: () => void;
}

const ChatComponent = (props:ChatComponentProps) => (
  <div className="container">
    <div className="header">
      <h1>{props.board_name}</h1>
    </div>
    <div className="messages">
      {props.messages.length === 0 && props.newMessage.length === 0 ? (
        <div>No Messages</div>
      ) : (
        <div>
          {/* 기존 DB Message */}
          {props.messages.slice().reverse().map((msg, index) => (
            <MsgComponent key={index} board_id={props.board_id}  writer={msg.writer} message={msg.message} me={props.me} />
          ))}
        </div>
      )}
      {props.newMessage.length > 0 ? (
        <div>
          {/* 새 메세지 */}
          {props.newMessage.map((msg, index) => {
            if (msg.board_id === props.board_id) {
              return <MsgComponent key={index} writer={msg.writer} message={msg.message} me={props.me} board_id={msg.board_id} />
            } 
          })}
        </div>
      ) : null}
    </div>
    <div className="input-area">
      <input
        type="text"
        value={props.value}
        onChange={(e) => props.setValue(e.target.value)}
        onKeyPress={(e) => {
          if (e.key === "Enter") {
            props.onSubmit();
          }
        }}
        placeholder="Type your message here..."
      />
      <button onClick={props.onSubmit}>Send</button>
    </div>
  </div>
);

export default ChatComponent;