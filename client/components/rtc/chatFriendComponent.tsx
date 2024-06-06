import React, { useEffect, useRef, useState,  } from 'react';
import { useRouter } from 'next/router';
import { closeWebSocket, initializeMessageWebSocket, sendMessage } from './webSocket';
import { getFriendMessageAxios } from '@/server/message';

const ChatFriendComponent = ({ username, friendname }: { username: string, friendname: string }) => {
    const router = useRouter();
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
            const messageData = { message: newMessage, sender: username, receiver: friendname };
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
        router.push({ pathname: '/' });
    };

    return (
        
        <div>
            <h1>To {friendname}</h1>
            <div style={{ border: '1px solid black', height: '600px', overflowY: 'scroll', padding: '10px' }} ref={chatContainerRef}>
            {messages.map((msg, index) => (
                <div key={index} style={{ textAlign: msg.sender === username ? 'right' : 'left', marginBottom: '5px' }}>
                <strong>{msg.sender}:</strong> {msg.message}
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
                      style={{ width: '80%', marginRight: '10px' }}
                />
                <button onClick={handleSendMessage}>Send</button>
            </div>
                <br></br>
            <button onClick={disconnect}>Back</button>

        </div>
    );
};

export default ChatFriendComponent;
