// /pages/[id].js
import { useRouter } from 'next/router';
import { useSession, SessionProvider } from 'next-auth/react';
import { useWebSocket } from '../socket/websocket';
import WebRTCComponent from '@/components/rtc/webRTC';
import { useEffect } from 'react';

export default function Page({ id }) {
  const username = localStorage.getItem('username')
  const { data: session } = useSession();

  return (
    <SessionProvider session={session}>
      <div>
        <WebRTCComponent chatId={id} userId={username} />
      </div>
      <div>
        <h1>Chat Room</h1>
        <p>Chat Room ID: {id}</p>
      </div>
    </SessionProvider>
  );
}

export async function getServerSideProps(context) {
  return { props: { id: context.params.id } };
}
