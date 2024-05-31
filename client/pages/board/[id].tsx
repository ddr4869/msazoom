import { useRouter } from 'next/router'
import { useSession, SessionProvider } from "next-auth/react";
import { socket, WebSocketComponent } from '../socket/websocket';
import { Websocket } from '../socket/websocket';
import axios from '../../server/axios'


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
      {/* <WebSocketComponent value={socket}>
        <Websocket board_id={id} board_name={router.query.board_name}/>
      </WebSocketComponent> */}
      <WebSocketComponent>
        <Websocket board_id={id} board_name={router.query.board_name}/>
      </WebSocketComponent>
    </SessionProvider>
  )
}

export async function getServerSideProps(context:any) {
  return { props: { id: context.params.id } };
}

