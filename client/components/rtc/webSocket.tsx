export const initializeMessageWebSocket = (username, friendname, handleWebSocketMessage) => {

  const webSocket = new WebSocket(`ws://`+process.env.NEXT_PUBLIC_HOST+`/api/message?user_name=${username}&friend_name=${friendname}`);
  
    webSocket.addEventListener('open', () => {
      console.log('WebSocket connection opened');
      webSocket.send(JSON.stringify({ join: true, partnerUsername:userId }));
    });
  
    webSocket.addEventListener('message', handleWebSocketMessage);
  
    return webSocket;
};

export const initializeChatWebSocket = (chatId, userId, handleWebSocketMessage) => {
  // fix!  
  
  const webSocket = new WebSocket(`ws://`+process.env.NEXT_PUBLIC_HOST+`/api/chat/join?chat_id=${chatId}&username=${userId}`);
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
      console.log('Closing WebSocket connection');
      webSocket.close();
    }
};

  