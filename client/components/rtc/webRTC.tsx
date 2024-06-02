import { useEffect, useRef } from 'react';
import { useRouter } from 'next/router';

const WebRTCComponent = ({ chatId, userId }) => {
  const router = useRouter();
  const userVideo = useRef();
  const userStream = useRef();
  const partnerVideo = useRef();
  const peerRef = useRef();
  const webSocketRef = useRef();

  const openCamera = async () => {
    console.log('Opening Camera');
    const constraints = {
      video: true,
      audio: true,
    };
    const stream = await navigator.mediaDevices.getUserMedia(constraints);
    userVideo.current.srcObject = stream;
    userStream.current = stream;
  };

  useEffect(() => {
    const start = async () => {
      await openCamera();
      
      webSocketRef.current = new WebSocket(`ws://localhost:8080/api/chat/join?chat_id=${chatId}&username=${userId}`);
      webSocketRef.current.addEventListener('open', () => {
        console.log('!! addEventListner open !!');
        webSocketRef.current.send(JSON.stringify({ join: true }));
      });

      webSocketRef.current.addEventListener('message', async (e) => {
        console.log('!! addEventListner message !!');
        console.log('Received data -> ', e.data);
        // if message is empty, return
        if (!e.data) return;    
        const message = JSON.parse(e.data);
        //console.log('message: ', message);

        if (message.join) {
          console.log('Join Message Received, Calling User');
          callUser();
        }

        if (message.offer) {
          console.log('Receiving Offer');
          handleOffer(message.offer);
        }

        if (message.answer) {
          console.log('Receiving Answer');
          await peerRef.current.setRemoteDescription(
            new RTCSessionDescription(message.answer)
          );
        }

        if (message.iceCandidate && peerRef.current) {
          console.log('Receiving and Adding ICE Candidate -> ', message.iceCandidate);
          peerRef.current.addIceCandidate(message.iceCandidate).catch(
              (err) => console.log("Error adding ICE Candidate"));
            console.log('*** ICE Candidate Added Success ***');
        }

        if (message.disconnect) {
          console.log('Peer disconnected');
          handlePeerDisconnect();
        }
      });
    };

    if (chatId) {
      start();
    }
  }, [chatId]);

  const handleOffer = async (offer) => {
    console.log('Received Offer, Creating Answer');
    peerRef.current = createPeer();

    await peerRef.current.setRemoteDescription(
      new RTCSessionDescription(offer)
    );

    userStream.current.getTracks().forEach((track) => {
      console.log('Adding Track to Peer -> ', track);
      peerRef.current.addTrack(track, userStream.current);
    });

    const answer = await peerRef.current.createAnswer();
    console.log('* Answer(createAnswer): ', answer);  
    await peerRef.current.setLocalDescription(answer);
    console.log('* peerRef.current.localDescription: ', peerRef.current.localDescription);  
    webSocketRef.current.send(
      JSON.stringify({ answer: peerRef.current.localDescription })
    );
  };

  const callUser = async () => {
    console.log('Calling Other User');
    peerRef.current = createPeer();

    userStream.current.getTracks().forEach((track) => {
      peerRef.current.addTrack(track, userStream.current);
    });
  };

  const createPeer = () => {
    console.log('Creating Peer Connection');
    const peer = new RTCPeerConnection({
      iceServers: [{ urls: 'stun:stun.l.google.com:19302' }],
    });

    peer.onnegotiationneeded = handleNegotiationNeeded;
    peer.onicecandidate = handleIceCandidateEvent;
    peer.ontrack = handleTrackEvent;

    return peer;
  };

  const handleNegotiationNeeded = async () => {
    console.log('Creating Offer');

    try {
      const myOffer = await peerRef.current.createOffer();
      await peerRef.current.setLocalDescription(myOffer);
      console.log('handleNegotiationNeeded localDescription -> ', peerRef.current.localDescription);
      webSocketRef.current.send(
        JSON.stringify({ offer: peerRef.current.localDescription })
      );
    } catch (err) {
      console.error(err);
    }
  };

  const handleIceCandidateEvent = async (e) => {
    console.log('Found Ice Candidate');
    if (e.candidate) {
      console.log(e.candidate);
      webSocketRef.current.send(
        JSON.stringify({ iceCandidate: e.candidate })
      );
    }
  };

  const handleTrackEvent = (e) => {
    console.log('Received Tracks, e.stream[0]->', e.streams[0]);
    partnerVideo.current.srcObject = e.streams[0]; 
  };

  const handlePeerDisconnect = () => {
    console.log('Handling peer disconnection');
    // Close the partner video stream
    if (partnerVideo.current && partnerVideo.current.srcObject) {
      partnerVideo.current.srcObject.getTracks().forEach((track) => track.stop());
      partnerVideo.current.srcObject = null;
    }
  
    // Close the peer connection
    if (peerRef.current) {
      peerRef.current.close();
      peerRef.current = null;
    }
  };
  
    // 뒤로가기
    const navigateToHome = () => {
      disconnect();
      router.push({
        pathname: '/',
      });
    };

    // disconnect 감지 
    useEffect(() => {
      const handlePopState = () => {
        disconnect();
      };  
      window.addEventListener('popstate', handlePopState);
      // Cleanup the event listener on component unmount
      return () => {
        window.removeEventListener('popstate', handlePopState);
      };
    }, []);

  // disconnect webRTC
  const disconnect = () => {
    console.log('Disconnecting');
    if (webSocketRef.current) {
      webSocketRef.current.send(JSON.stringify({ disconnect: true }));
    }

    userStream.current.getTracks().forEach((track) => {
      track.stop();
    });

    if (peerRef.current) {
      peerRef.current.close();
    }

    if (webSocketRef.current) {
      webSocketRef.current.close();
    }
  };

  return (
    <div>
      <button onClick={navigateToHome}>뒤로가기</button>
      <div
        style={{
          display: 'flex',
          justifyContent: 'center',
          alignItems: 'center',
          top: '100px',
          right: '100px',
          borderRadius: '10px',
          overflow: 'hidden',
        }}
      >
        <video playsInline autoPlay muted controls ref={userVideo} />
        <video playsInline autoPlay controls ref={partnerVideo} />
      </div>
      <div>
        <button onClick={disconnect}>Disconnect</button>

      </div>
    </div>
  );
};

export default WebRTCComponent;
