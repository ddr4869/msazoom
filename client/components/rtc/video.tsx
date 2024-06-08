export const openCamera = async ({ userVideoRef, userStreamRef }: { userVideoRef: React.RefObject<HTMLVideoElement>, userStreamRef: React.MutableRefObject<MediaStream | null> }): Promise<MediaStream> => {
  const constraints = {
    video: true,
    audio: true,
  };
  const stream = await navigator.mediaDevices.getUserMedia(constraints);
  if (userVideoRef.current) {
    userVideoRef.current.srcObject = stream;
  }
  userStreamRef.current = stream;
  return stream;
};

export const stopStreamTracks = (stream: MediaStream | null) => {
  if (stream) {
    stream.getTracks().forEach(track => track.stop());
  }
};