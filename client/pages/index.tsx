import React, { useState, useEffect } from 'react';
import { w3cwebsocket as W3CWebSocket } from 'websocket';

const client = new W3CWebSocket('ws://localhost:8080/ws');

export default function WebSocketPage() {
  const [message, setMessage] = useState('');
  const [response, setResponse] = useState('');

  useEffect(() => {
    client.onopen = () => {
      console.log('WebSocket Client Connected');
    };
    client.onmessage = (message) => {
      console.log('Received message:', message.data);
      setResponse(message.data);
    };
    client.onclose = () => {
      console.log('WebSocket Client Closed');
    };
    client.onerror = (error) => {
      console.error('WebSocket Client Error:', error);
    };
  }, []);

  const sendMessage = () => {
    if (client.readyState === client.OPEN) {
      console.log('Sending message:', message);
      client.send(message);
    }
  };

  return (
    <div>
      <h1>WebSocket Example</h1>
      <input type="text" value={message} onChange={(e) => setMessage(e.target.value)} />
      <button onClick={sendMessage}>Send Message</button>
      <p>Response from server: {response}</p>
    </div>
  );
}
