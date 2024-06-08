import { LoginAxios, SignupAxios } from '@/server/user';
import { FormEvent } from 'react';

export const handleLogin = async (username: string, password: string, setIsLoggedIn: (value: boolean) => void, event: any) => {
  event.preventDefault();
  await LoginAxios(username, password);
  setIsLoggedIn(true);
};

export const handleLogout = (setIsLoggedIn: (value: boolean) => void) => {
  localStorage.removeItem('accessToken');
  localStorage.removeItem('username');
  setIsLoggedIn(false);
};

export const handleSignup = async (username: string, password: string ) => {
  //event.preventDefault();
  console.log("handleSignup -> ", username)
  await SignupAxios(username, password);
};