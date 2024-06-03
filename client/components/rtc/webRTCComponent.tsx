// WebRTCComponent.js
import { useEffect, useState, useRef } from 'react';
import { useRouter } from 'next/router';
import { initializeWebSocket, sendMessage, closeWebSocket } from './webSocket';
import { createPeerConnection, handleOffer, addIceCandidate, closePeerConnection } from './rtcPeerConnection';
import { openCamera, stopStreamTracks } from './video';

const WebRTCComponent = ({ chatId, userId }) => {
  const router = useRouter();
  const userVideo = useRef();
  const userStream = useRef();
  const partnerVideo = useRef();
  const peerRef = useRef();
  const webSocketRef = useRef();
  const chatEndRef = useRef();
  const [partnerUsername, setPartnerUsername] = useState(null);
  const [hostName, setHostName] = useState(null);
  const [messages, setMessages] = useState([]);
  const [newMessage, setNewMessage] = useState('');

  const handleWebSocketMessage = async (e) => {
    const message = JSON.parse(e.data);

    if (message.join) {
      console.log('Joining Room -> ', message.join);
      setPartnerUsername(message.partnerUsername);
      webSocketRef.current.send(JSON.stringify({ hostUserName: userId })); 
      await createAndCallUser();
    }

    if (message.hostUserName) {
        console.log('Host -> ', message.hostUserName);
        setPartnerUsername(message.hostUserName);
    }

    if (message.offer) {
      console.log('Receiving Offer -> ', message.offer);
      await handleOffer(peerRef.current, message.offer, userStream.current, (msg) => sendMessage(webSocketRef.current, msg));
    }

    if (message.answer) {
      await peerRef.current.setRemoteDescription(new RTCSessionDescription(message.answer));
    }

    if (message.iceCandidate && peerRef.current) {
      console.log("iceCandidate", message.iceCandidate)
      await addIceCandidate(peerRef.current, message.iceCandidate);
    }

    if (message.chatMessage) {
      setMessages((prevMessages) => [...prevMessages, { from: message.chatFrom, text: message.chatMessage }]);
    }

    if (message.disconnect) {
      handlePeerDisconnect();
      setPartnerUsername(null);
    }
  };

  useEffect(() => {
    const start = async () => {
      userStream.current = await openCamera(userVideo, userStream);
      webSocketRef.current = initializeWebSocket(chatId, userId, handleWebSocketMessage);
      peerRef.current = createPeerConnection(handleNegotiationNeeded, handleIceCandidateEvent, handleTrackEvent);
      
    };

    if (chatId) {
      start();
    }
  }, [chatId]);

  useEffect(() => {
    chatEndRef.current.scrollIntoView({ behavior: 'smooth' });
  }, [messages]);

  const createAndCallUser = async () => {
    if (peerRef.current) {
      closePeerConnection(peerRef.current);
    }
    peerRef.current = createPeerConnection(handleNegotiationNeeded, handleIceCandidateEvent, handleTrackEvent);
    callUser();
  };

  const callUser = async () => {
    userStream.current.getTracks().forEach(track => {
      peerRef.current.addTrack(track, userStream.current);
    });
  };

  const handleNegotiationNeeded = async () => {
    console.log("Negotiation Needed")
    const offer = await peerRef.current.createOffer();
    await peerRef.current.setLocalDescription(offer);
    sendMessage(webSocketRef.current, { offer: peerRef.current.localDescription });
  };

  const handleIceCandidateEvent = (e) => {
    console.log("Ice Candidate Event")
    if (e.candidate) {
      sendMessage(webSocketRef.current, { iceCandidate: e.candidate });
    }
  };

  const handleTrackEvent = (e) => {
    console.log("Track Event")
    partnerVideo.current.srcObject = e.streams[0];
  };

  const handlePeerDisconnect = () => {
    if (partnerVideo.current && partnerVideo.current.srcObject) {
      stopStreamTracks(partnerVideo.current.srcObject);
      partnerVideo.current.srcObject = null;
    }
    closePeerConnection(peerRef.current);
  };

  const navigateToHome = () => {
    disconnect();
    router.push({ pathname: '/' });
  };

  useEffect(() => {
    const handlePopState = () => {
      disconnect();
    };
    window.addEventListener('popstate', handlePopState);
    return () => {
      window.removeEventListener('popstate', handlePopState);
    };
  }, []);

  const disconnect = () => {
    sendMessage(webSocketRef.current, { disconnect: true });
    stopStreamTracks(userStream.current);
    closePeerConnection(peerRef.current);
    closeWebSocket(webSocketRef.current);
  };

  const handleSendMessage = () => {
    if (newMessage.trim() !== '') {
      console.log('new Message -> ', newMessage);
      sendMessage(webSocketRef.current, { chatMessage: newMessage, chatFrom: userId});
      setMessages((prevMessages) => [...prevMessages, { from: 'Me', text: newMessage }]);
      setNewMessage('');
    }
  };

  const handleKeyDown = (e) => {
    if (e.key === 'Enter') {
      e.preventDefault(); // Prevent the default behavior of the Enter key
      console.log('Enter Pressed')
      handleSendMessage();
    }
  };

  return (
    <div>
      <button onClick={navigateToHome}>Disconnect</button>
      <div style={{ display: 'flex', justifyContent: 'space-around', alignItems: 'center', top: '100px', right: '100px', borderRadius: '10px', overflow: 'hidden' }}>
        <div style={{ textAlign: 'center' }}>
            {
                partnerUsername ? (
                <h2>Partner Video - {partnerUsername}</h2>
                ) : (
                    <h2>Waiting for Partner</h2>
                )
            }
          <video playsInline autoPlay controls ref={partnerVideo} style={{ width: '500px', height: '350px' }} />
        </div>
        <div style={{ textAlign: 'center' }}>
          <h2>Your Video</h2>
          <video playsInline autoPlay muted controls ref={userVideo} style={{ width: '500px', height: '350px' }} />
        </div>
      </div>

      <div style={{ marginTop: '20px' }}>
        <h2>Chat</h2>
        <div style={{ border: '1px solid black', height: '200px', overflowY: 'scroll', padding: '10px' }}>
          {messages.map((msg, index) => (
            <div key={index} style={{ textAlign: msg.from === 'Me' ? 'right' : 'left', marginBottom: '5px' }}>
              <strong>{msg.from}:</strong> {msg.text}
            </div>
          ))}
          <div ref={chatEndRef}></div>
        </div>
        <input
          type="text"
          value={newMessage}
          onChange={(e) => setNewMessage(e.target.value)}
          onKeyDown={(e) => {
            if (e.key === 'Enter') {
              handleSendMessage()
            }
          }}
          //onKeyDown={handleKeyDown}
          style={{ width: '80%', marginRight: '10px' }}
        />
        <button onClick={handleSendMessage}>Send</button>
      </div>
    </div>
  );
};

export default WebRTCComponent;
