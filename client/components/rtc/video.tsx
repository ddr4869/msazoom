// video.js
export const openCamera = async (userVideoRef, userStreamRef) => {
    const constraints = {
      video: true,
      audio: true,
    };
    const stream = await navigator.mediaDevices.getUserMedia(constraints);
    userVideoRef.current.srcObject = stream;
    userStreamRef.current = stream;
    return stream;
  };
  
  export const stopStreamTracks = (stream) => {
    if (stream) {
      stream.getTracks().forEach(track => track.stop());
    }
  };
  