type initializeMessageWebSocketProps = {
  username: string;
  friendname: string;
  handleWebSocketMessage: any;
}

type initializeChatWebSocketProps = {
  chatId: string;
  userId: string;
  password: string;
  handleWebSocketMessage: any;
}

export const initializeMessageWebSocket = (  {username ,  friendname,  handleWebSocketMessage} : initializeMessageWebSocketProps) => {

  const webSocket = new WebSocket(`ws://`+process.env.NEXT_PUBLIC_HOST+`/api/message/connect?user_name=${username}&friend_name=${friendname}`);
  
    webSocket.addEventListener('open', () => {
      console.log('WebSocket connection opened');
    });
  
    webSocket.addEventListener('message', handleWebSocketMessage);
  
    return webSocket;
};

export const initializeChatWebSocket = ( {chatId, userId, password, handleWebSocketMessage}: initializeChatWebSocketProps) => {
  // fix!  
  const webSocket = new WebSocket(`ws://`+process.env.NEXT_PUBLIC_HOST+`/api/chat/join?chat_id=${chatId}&username=${userId}&password=${password}`);
  webSocket.addEventListener('open', () => {
    console.log('WebSocket connection opened');
    webSocket.send(JSON.stringify({ join: true, partnerUsername:userId }));
  });

  webSocket.addEventListener('message', handleWebSocketMessage);

  return webSocket;
};
  
export const sendMessage = (webSocket:WebSocket, message:any) => {
    webSocket.send(JSON.stringify(message));
};
  
export const closeWebSocket = (webSocket:WebSocket) => {
    if (webSocket) {
      console.log('Closing WebSocket connection');
      webSocket.close();
    }
};

  