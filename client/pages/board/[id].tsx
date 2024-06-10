// /pages/[id].js
import { useRouter } from 'next/router';
import { useSession, SessionProvider } from 'next-auth/react';
import WebRTCComponent from '@/components/rtc/webRTCComponent';
import { useEffect } from 'react';

export default function Page({ id }:any) {
  const router = useRouter();
  const username = localStorage.getItem('username') || router.push('/login');
  const { data: session } = useSession();
  
  const navigateToDashboard = () => {
    router.push({
      pathname: '/dashboard',
    });
  };

  return (
    <p></p>
  );
}

export async function getServerSideProps(context:any) {
  return { props: { id: context.params.id } };
}
