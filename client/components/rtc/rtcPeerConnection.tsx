// rtcPeerConnection.js
export const createPeerConnection = (handleNegotiationNeeded, handleIceCandidateEvent, handleTrackEvent) => {
    const peer = new RTCPeerConnection({
      iceServers: [{ urls: 'stun:stun.l.google.com:19302' }],
    });
  
    peer.onnegotiationneeded = handleNegotiationNeeded;
    peer.onicecandidate = handleIceCandidateEvent;
    peer.ontrack = handleTrackEvent;
  
    return peer;
  };
  
  export const handleOffer = async (peer, offer, userStream, sendMessage) => {
    await peer.setRemoteDescription(new RTCSessionDescription(offer));
  
    userStream.getTracks().forEach(track => {
      peer.addTrack(track, userStream);
    });
  
    const answer = await peer.createAnswer();
    await peer.setLocalDescription(answer);
    sendMessage({ answer: peer.localDescription });
  };
  
  export const addIceCandidate = async (peer, candidate) => {
    try {
      await peer.addIceCandidate(candidate);
    } catch (err) {
      console.error("Error adding ICE Candidate", err);
    }
  };
  
  export const closePeerConnection = (peer) => {
    if (peer) {
      peer.close();
    }
  };
  