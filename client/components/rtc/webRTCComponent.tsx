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
  const [partnerUsername, setPartnerUsername] = useState(null);

  const handleWebSocketMessage = async (e) => {
    const message = JSON.parse(e.data);

    if (message.join) {
      callUser();
      console.log('Joining -> ', message.partnerUsername);
      setPartnerUsername(message.partnerUsername);
      webSocketRef.current.send(JSON.stringify({ hostname: userId }));
    }

    if (message.hostname) {
        console.log('Host -> ', message.hostname);
        setPartnerUsername(message.hostname);
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

  const callUser = async () => {
    userStream.current.getTracks().forEach(track => {
      peerRef.current.addTrack(track, userStream.current);
    });
  };

  const handleNegotiationNeeded = async () => {
    const offer = await peerRef.current.createOffer();
    await peerRef.current.setLocalDescription(offer);
    sendMessage(webSocketRef.current, { offer: peerRef.current.localDescription });
  };

  const handleIceCandidateEvent = (e) => {
    if (e.candidate) {
      sendMessage(webSocketRef.current, { iceCandidate: e.candidate });
    }
  };

  const handleTrackEvent = (e) => {
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
          <video playsInline autoPlay controls ref={partnerVideo} style={{ width: '600px', height: '450px' }} />
        </div>
        <div style={{ textAlign: 'center' }}>
          <h2>Your Video</h2>
          <video playsInline autoPlay muted controls ref={userVideo} style={{ width: '600px', height: '450px' }} />
        </div>
      </div>
    </div>
  );
};

export default WebRTCComponent;
