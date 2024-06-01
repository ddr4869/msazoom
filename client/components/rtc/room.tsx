import { useEffect, useRef } from 'react';

const Room = ({ roomId }) => {
  const userVideo = useRef();
  const userStream = useRef();
  const partnerVideo = useRef();
  const peerRef = useRef();
  const webSocketRef = useRef();

  const openCamera = async () => {
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
      webSocketRef.current = new WebSocket(`ws://localhost:8080/api/rtc/join?room_iD=${roomId}`);

      webSocketRef.current.addEventListener('open', () => {
        webSocketRef.current.send(JSON.stringify({ join: true }));
      });

      webSocketRef.current.addEventListener('message', async (e) => {
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

        if (message.iceCandidate) {
          console.log('Receiving and Adding ICE Candidate -> ', message.iceCandidate);
          peerRef.current.addIceCandidate(message.iceCandidate).catch(
              (err) => console.log("Error adding ICE Candidate"));
            console.log('*** ICE Candidate Added Success ***');
        }
      });
    };

    if (roomId) {
      start();
    }
  }, [roomId]);

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

  return (
    <div>
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
    </div>
  );
};

export default Room;
