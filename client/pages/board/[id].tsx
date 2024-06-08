// /pages/[id].js
import { useRouter } from 'next/router';
import { useSession, SessionProvider } from 'next-auth/react';
import WebRTCComponent from '@/components/rtc/webRTCComponent';
import { useEffect } from 'react';

export default function Page({ id }:any) {
  const router = useRouter();
  const username = localStorage.getItem('username') || '';
  const { data: session } = useSession();
  
  const navigateToDashboard = () => {
    router.push({
      pathname: '/dashboard',
    });
  };

  return (
    <SessionProvider session={session}>
      <button onClick={navigateToDashboard}>뒤로가기</button>
      <div>
        <WebRTCComponent chatId={id} userId={username} />
      </div>
    </SessionProvider>
  );
}

export async function getServerSideProps(context:any) {
  return { props: { id: context.params.id } };
}
