import { Component } from "react";

interface MsgComponentProps {
    board_id: string;
    writer: string;
    message: string;
    me: string;
}

const MsgComponent = (props:MsgComponentProps) => (
    <div className="message">
      <p> 

        <strong>{props.writer === props.me ? 'me' : props.writer}</strong>: {props.message}
      </p>
    </div>
  );

export type { MsgComponentProps };
export default MsgComponent;

  