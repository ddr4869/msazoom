import { useRouter } from 'next/router'
import { useSession, SessionProvider } from "next-auth/react";
import { socket, WebsocketProvider } from '../socket/websocketContext';
import { Websocket } from '../socket/websocket';
import axios from '../server/axios'


export default function Page({id}:any) {
  const router = useRouter()  
  const navigateToDashboard = () => {
    router.push({
      pathname: `/dashboard`,
    });
  };

  return (
      <SessionProvider>
        <button onClick={navigateToDashboard}>뒤로가기</button>
      <WebsocketProvider value={socket}>
        <Websocket board_id={id} board_name={router.query.board_name}/>
      </WebsocketProvider>
    </SessionProvider>
  )
}

export async function getServerSideProps(context:any) {
  return { props: { id: context.params.id } };
}

