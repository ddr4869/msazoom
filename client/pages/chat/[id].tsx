import { useEffect, useState } from 'react';
import { useRouter } from 'next/router';
import { useSession, SessionProvider } from 'next-auth/react';
import WebRTCComponent from '@/components/rtc/webRTCComponent';

export default function Page({ id }:any) {
  const { data: session } = useSession();
  const router = useRouter();
  const [username, setUsername] = useState<string>("");

  useEffect(() => {
    const user = localStorage.getItem('username');
    console.log("user -> ", user)
    if (!user) {
      router.push('/');
    } else {
      setUsername(user);
    }
  }, [router]);

  if (!username) {
    return null; // 또는 로딩 스피너를 반환
  }
  
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
