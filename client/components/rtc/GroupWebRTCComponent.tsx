import { useEffect, useState, useRef } from 'react';
import { useRouter } from 'next/router';
import {
  initializeChatWebSocket,
  sendMessage,
  closeWebSocket,
} from './webSocket';
import {
  createPeerConnection,
  handleOffer,
  addIceCandidate,
  closePeerConnection,
} from './rtcPeerConnection';
import { openCamera, stopStreamTracks } from './video';

interface GroupWebRTCComponentProps {
  chatId: string;
  userId: string;
  password: string;
}

const GroupWebRTCComponent = ({
  chatId,
  userId,
  password,
}: GroupWebRTCComponentProps) => {
  const router = useRouter();
  const userVideo = useRef<HTMLVideoElement>(null);
  const userStream = useRef<MediaStream | null>(null);
  const peerConnections = useRef<{ [key: string]: RTCPeerConnection }>({});
  const webSocketRef = useRef<WebSocket | null>(null);
  const [participants, setParticipants] = useState<string[]>([]);
  const [messages, setMessages] = useState<
    Array<{ from: string; text: string }>
  >([]);
  const [newMessage, setNewMessage] = useState('');

  const handleWebSocketMessage = async (e: MessageEvent) => {
    const message = JSON.parse(e.data);
    console.log('Received Message -> ', message);

    if (message.join) {
      console.log('New participant joined -> ', message.partnerUsername);
      setParticipants((prev) => [...prev, message.partnerUsername]);
      createAndCallUser(message.partnerUsername);
    }

    if (message.offer) {
      console.log('Receiving Offer -> ', message.offer);
      await handleOffer(
        peerConnections.current[message.from],
        message.offer,
        userStream.current!,
        (msg) => sendMessage(webSocketRef.current!, { ...msg, to: message.from })
      );
    }

    if (message.answer) {
      console.log('Receiving Answer -> ', message.answer);
      await peerConnections.current[message.from].setRemoteDescription(
        new RTCSessionDescription(message.answer)
      );
    }

    if (message.iceCandidate) {
      console.log('ICE Candidate received -> ', message.iceCandidate);
      await addIceCandidate(
        peerConnections.current[message.from],
        message.iceCandidate
      );
    }

    if (message.chatMessage) {
      setMessages((prev) => [
        ...prev,
        { from: message.chatFrom, text: message.chatMessage },
      ]);
    }

    if (message.disconnect) {
      console.log('Participant disconnected -> ', message.from);
      handleParticipantDisconnect(message.from);
    }
  };

  const createAndCallUser = (participant: string) => {
    if (!peerConnections.current[participant]) {
      peerConnections.current[participant] = createPeerConnection({
        handleNegotiationNeeded: () => handleNegotiationNeeded(participant),
        handleIceCandidateEvent: (e) => handleIceCandidateEvent(e, participant),
        handleTrackEvent,
      });
    }
    callUser(participant);
  };

  const callUser = (participant: string) => {
    userStream.current?.getTracks().forEach((track) => {
      peerConnections.current[participant].addTrack(track, userStream.current!);
    });
  };

  const handleNegotiationNeeded = async (participant: string) => {
    const peer = peerConnections.current[participant];
    const offer = await peer.createOffer();
    await peer.setLocalDescription(offer);
    sendMessage(webSocketRef.current!, {
      offer: peer.localDescription,
      to: participant,
      from: userId,
    });
  };

  const handleIceCandidateEvent = (
    e: RTCPeerConnectionIceEvent,
    participant: string
  ) => {
    if (e.candidate) {
      sendMessage(webSocketRef.current!, {
        iceCandidate: e.candidate,
        to: participant,
        from: userId,
      });
    }
  };

  const handleTrackEvent = (e: RTCTrackEvent) => {
    const participant = e.streams[0].id; // Assuming stream ID is unique for each participant
    const videoElement = document.getElementById(
      `video-${participant}`
    ) as HTMLVideoElement;
    if (videoElement) {
      videoElement.srcObject = e.streams[0];
    }
  };

  const handleParticipantDisconnect = (participant: string) => {
    const peer = peerConnections.current[participant];
    if (peer) {
      closePeerConnection(peer);
      delete peerConnections.current[participant];
    }
    setParticipants((prev) => prev.filter((p) => p !== participant));
  };

  useEffect(() => {
    const start = async () => {
      userStream.current = await openCamera({
        userVideoRef: userVideo,
        userStreamRef: userStream,
      });
      webSocketRef.current = initializeChatWebSocket({
        chatId,
        userId,
        password,
        handleWebSocketMessage,
      });
    };

    if (chatId) {
      start();
    }

    return () => {
      disconnect();
    };
  }, [chatId]);

  const disconnect = () => {
    sendMessage(webSocketRef.current!, { disconnect: true, from: userId });
    if (userStream.current) {
      stopStreamTracks(userStream.current);
    }
    Object.keys(peerConnections.current).forEach((participant) => {
      closePeerConnection(peerConnections.current[participant]);
    });
    if (webSocketRef.current) {
      closeWebSocket(webSocketRef.current);
    }
  };

  const handleSendMessage = () => {
    if (newMessage.trim() !== '') {
      sendMessage(webSocketRef.current!, {
        chatMessage: newMessage,
        chatFrom: userId,
      });
      setMessages((prev) => [...prev, { from: 'Me', text: newMessage }]);
      setNewMessage('');
    }
  };

  return (
    <div>
      <button onClick={disconnect}>Disconnect</button>
      <div style={{ display: 'flex', justifyContent: 'space-around' }}>
        <div style={{ textAlign: 'center' }}>
          <h2>Your Video</h2>
          <video
            playsInline
            autoPlay
            muted
            controls
            ref={userVideo}
            style={{ width: '500px', height: '350px' }}
          />
        </div>
        {participants.map((participant) => (
          <div key={participant} style={{ textAlign: 'center' }}>
            <h2>{participant}'s Video</h2>
            <video
              id={`video-${participant}`}
              playsInline
              autoPlay
              controls
              style={{ width: '500px', height: '350px' }}
            />
          </div>
        ))}
      </div>

      <div style={{ marginTop: '20px' }}>
        <h1>Chatting</h1>
        <div
          style={{
            border: '1px solid black',
            height: '200px',
            overflowY: 'scroll',
            padding: '10px',
          }}
        >
          {messages.map((msg, index) => (
            <div
              key={index}
              style={{
                textAlign:
                  msg.from === 'System'
                    ? 'center'
                    : msg.from === 'Me'
                    ? 'right'
                    : 'left',
                marginBottom: '5px',
              }}
            >
              <strong>{msg.from}:</strong> {msg.text}
            </div>
          ))}
        </div>
        <input
          type="text"
          value={newMessage}
          onChange={(e) => setNewMessage(e.target.value)}
          onKeyPress={(e) => {
            if (e.key === 'Enter') {
              handleSendMessage();
            }
          }}
          style={{ width: '80%', marginRight: '10px' }}
        />
        <button onClick={handleSendMessage}>Send</button>
      </div>
    </div>
  );
};

export default GroupWebRTCComponent;
