import { useEffect, useState } from 'react';
import { useRouter } from 'next/router';
import { useSession, SessionProvider } from 'next-auth/react';
import WebRTCComponent from '@/components/rtc/webRTCComponent';

export default function Page({ id }:any) {
  const router = useRouter();
  //const password = Array.isArray(router.query.password) ? router.query.password[0] : router.query.password || "";
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
    <>
      <div>
        <WebRTCComponent chatId={id} userId={username} password=''/>
      </div>
      <div>
        <h1>Chat Room</h1>
        <p>Chat Room ID: {id}</p>
      </div>
    </>
  );
}

export async function getServerSideProps(context:any) {
  return { props: { id: context.params.id } };
}
