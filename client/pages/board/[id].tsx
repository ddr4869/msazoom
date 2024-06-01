// /pages/[id].js
import { useRouter } from 'next/router';
import { useSession, SessionProvider } from 'next-auth/react';
import { useWebSocket } from '../socket/websocket';

export default function Page({ id }) {
  const router = useRouter();
  const { data: session } = useSession();
  const { messages, sendMessage } = useWebSocket(); // WebSocket 서버 URL

  const navigateToDashboard = () => {
    router.push({
      pathname: '/dashboard',
    });
  };

  const handleSendMessage = () => {
    const message = { id, content: 'Hello from Next.js!' };
    sendMessage(message);
  };

  return (
    <SessionProvider session={session}>
      <button onClick={navigateToDashboard}>뒤로가기</button>
      <button onClick={handleSendMessage}>Send WebSocket Message</button>
      <div>
        <h2>Received Messages:</h2>
        <ul>
          {messages.map((msg, index) => (
            <li key={index}>{msg.content}</li>
          ))}
        </ul>
      </div>
    </SessionProvider>
  );
}

export async function getServerSideProps(context) {
  return { props: { id: context.params.id } };
}
