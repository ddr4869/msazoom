import { LoginAxios } from '@/server/user';
import { toNamespacedPath } from 'path';

export const handleLogin = async (username: string, password: string, setIsLoggedIn: (value: boolean) => void, event: FormEvent<HTMLFormElement>) => {
  event.preventDefault();
  await LoginAxios(username, password);
  setIsLoggedIn(true);
};

export const handleLogout = (setIsLoggedIn: (value: boolean) => void) => {
  localStorage.removeItem('accessToken');
  localStorage.removeItem('username');
  setIsLoggedIn(false);
};
