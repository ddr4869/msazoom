// pages/index.js
import { useState, useEffect, FormEvent } from 'react';
import axios from 'axios';
import { signIn, signOut, useSession } from 'next-auth/react';
import { useRouter } from "next/router";
import { createBoardAxios, getBoardsAxios, recommendBoardAxios, deleteBoardAxios } from '@/server/board';
import boardStyles from '@/styles/board-styles.module.css'
import userStyles from '@/styles/userProfile-styles.module.css'

const Home = () => {

  return (
    <h1>Home</h1>
  );
};

export default Home;
