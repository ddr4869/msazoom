// pages/index.js
import { useState, useEffect, FormEvent } from 'react';
import axios from 'axios';
import { signIn, signOut, useSession } from 'next-auth/react';
import { useRouter } from "next/router";
import { createBoardAxios, getBoardsAxios, recommendBoardAxios, deleteBoardAxios } from '@/server/board';

import Room from '@/components/rtc/webRTC_backup';
const RoomPage = () => {
    const router = useRouter();
    const { id } = router.query;
    return (
      <div>
        <h1>Golang & React</h1>
        <Room roomId={id} />
      </div>
    );
  };
  
  export default RoomPage;