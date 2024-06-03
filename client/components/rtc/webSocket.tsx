// webSocket.js
export const initializeWebSocket = (chatId, userId, handleWebSocketMessage) => {
    const webSocket = new WebSocket(`ws://localhost:8080/api/chat/join?chat_id=${chatId}&username=${userId}`);
    
    webSocket.addEventListener('open', () => {
      console.log('WebSocket connection opened');
      webSocket.send(JSON.stringify({ join: true, partnerUsername:userId }));
    });
  
    webSocket.addEventListener('message', handleWebSocketMessage);
  
    return webSocket;
  };
  
  export const sendMessage = (webSocket, message) => {
    webSocket.send(JSON.stringify(message));
  };
  
  export const closeWebSocket = (webSocket) => {
    if (webSocket) {
      webSocket.close();
    }
  };

  