import React, { useEffect, useRef, useState } from 'react';
import { closeWebSocket, initializeMessageWebSocket, sendMessage } from './webSocket';
import { getFriendMessageAxios } from '@/server/message';

const ChatFriendComponent = ({ username, friendname }: { username: string, friendname: string }) => {
    const webSocketRef = useRef<WebSocket | null>(null);
    const [messages, setMessages] = useState([]);
    const [newMessage, setNewMessage] = useState('');
    const chatContainerRef = useRef();

    const handleWebSocketMessage = async (e: any) => {
        console.log('WebSocket message received:', e.data);
        const message = JSON.parse(e.data);
        if (message.message) {
            console.log('Received message:', message);
            setMessages(prevMessages => [...prevMessages, message]);
        }
    };

    useEffect(() => {
        const start = async () => {
            try {
                const response = await getFriendMessageAxios(localStorage.getItem("accessToken"), friendname);
                setMessages(response);
            } catch (error) {
                console.error('Error fetching messages:', error);
            }

            webSocketRef.current = initializeMessageWebSocket({username, friendname, handleWebSocketMessage});
        };
        if (username && friendname) {
            start();
        }
        return () => {
            if (webSocketRef.current) {
                console.log('Closing WebSocket connection');
                webSocketRef.current.close();
            }
        };
    }, [username, friendname]);

    const handleSendMessage = () => {
        if (webSocketRef.current && newMessage.trim() !== '') {
            console.log('handleSendMessage -> ', newMessage);
            const messageData = { message: newMessage, writer: username, sender: username, receiver: friendname };
            webSocketRef.current.send(JSON.stringify(messageData));
            setMessages(prevMessages => [...prevMessages, messageData]);
            setNewMessage('');
        }
    };

    // 새로고침, 뒤로가기 check
    useEffect(() => {
        const handleDisconnectState = () => {
            disconnect();
        };
        window.addEventListener('popstate', handleDisconnectState);
        window.addEventListener('beforeunload', handleDisconnectState);
        return () => {
            window.removeEventListener('popstate', handleDisconnectState);
            window.removeEventListener('beforeunload', handleDisconnectState);
        };
    }, []);

    useEffect(() => {
        if (chatContainerRef.current) {
            chatContainerRef.current.scrollTop = chatContainerRef.current.scrollHeight;
        }
    }, [messages]);

    const disconnect = () => {
        sendMessage(webSocketRef.current, { disconnect: true });
        closeWebSocket(webSocketRef.current);
    };

    return (
        <div>
            <div className="chat-container" ref={chatContainerRef}>
                {messages.map((msg, index) => (
                    <div key={index} className={msg.writer === username ? 'my-message' : 'friend-message'}>
                        <span>{msg.message}</span>
                    </div>
                ))}
            </div>
            <div className="input-container">
                <input
                    type="text"
                    value={newMessage}
                    onChange={(e) => setNewMessage(e.target.value)}
                    placeholder="Type a message..."
                    onKeyPress={(e) => {
                        if (e.key === 'Enter') {
                          handleSendMessage()
                        }
                      }}
                />
                <button onClick={handleSendMessage}>Send</button>
            </div>
            <style jsx>{`
                .chat-container {
                    display: flex;
                    flex-direction: column;
                    max-height: 300px;
                    overflow-y: auto;
                    padding: 10px;
                }
                .my-message {
                    align-self: flex-end;
                    background-color: #DCF8C6;
                    padding: 5px 10px;
                    border-radius: 10px;
                    margin: 5px;
                }
                .friend-message {
                    align-self: flex-start;
                    background-color: #FFF;
                    padding: 5px 10px;
                    border-radius: 10px;
                    margin: 5px;
                }
                .input-container {
                    display: flex;
                    margin-top: 10px;
                }
                .input-container input {
                    flex: 1;
                    padding: 10px;
                    border: 1px solid #ccc;
                    border-radius: 5px;
                }
                .input-container button {
                    padding: 10px 20px;
                    border: none;
                    background-color: #4CAF50;
                    color: white;
                    cursor: pointer;
                    border-radius: 5px;
                    margin-left: 10px;
                }
            `}</style>
        </div>
    );
};

export default ChatFriendComponent;
