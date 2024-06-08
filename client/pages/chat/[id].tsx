// /pages/[id].js
import { useSession, SessionProvider } from 'next-auth/react';
import WebRTCComponent from '@/components/rtc/webRTCComponent';

export default function Page({ id }:any) {
  const { data: session } = useSession();
  const username = localStorage.getItem('username') || '';
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

export async function getServerSideProps(context:any) {
  return { props: { id: context.params.id } };
}
