interface createPeerConnectionProps {
  handleNegotiationNeeded: () => void;
  handleIceCandidateEvent: (event: RTCPeerConnectionIceEvent) => void;
  handleTrackEvent: (event: RTCTrackEvent) => void;
}

export const createPeerConnection = ({ handleNegotiationNeeded, handleIceCandidateEvent, handleTrackEvent }: createPeerConnectionProps): RTCPeerConnection => {
  console.log('Creating Peer Connection');
  const peer = new RTCPeerConnection({
    iceServers: [{ urls: 'stun:stun.l.google.com:19302' }],
  });

  peer.onnegotiationneeded = handleNegotiationNeeded;
  peer.onicecandidate = handleIceCandidateEvent;
  peer.ontrack = handleTrackEvent;

  return peer;
};

interface RTCSessionDescriptionInit {
  type: RTCSdpType;
  sdp: string;
}

interface SendMessage {
  (message: { answer: RTCSessionDescription | null }): void;
}

export const handleOffer = async (
  peer: RTCPeerConnection,
  offer: RTCSessionDescriptionInit,
  userStream: MediaStream,
  sendMessage: SendMessage
): Promise<void> => {
  await peer.setRemoteDescription(new RTCSessionDescription(offer));

  userStream.getTracks().forEach(track => {
    peer.addTrack(track, userStream);
  });

  const answer = await peer.createAnswer();
  await peer.setLocalDescription(answer);
  sendMessage({ answer: peer.localDescription });
};

export const addIceCandidate = async (peer: RTCPeerConnection, candidate: RTCIceCandidateInit): Promise<void> => {
  try {
    await peer.addIceCandidate(new RTCIceCandidate(candidate));
  } catch (err) {
    console.error("Error adding ICE Candidate", err);
  }
};

export const closePeerConnection = (peer: RTCPeerConnection | null): void => {
  if (peer) {
    console.log('Closing Peer Connection');
    peer.close();
  }
};
