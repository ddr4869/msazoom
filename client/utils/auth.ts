import axios from 'axios';
import { toNamespacedPath } from 'path';

export const handleLogin = async (username: string, password: string, setIsLoggedIn: (value: boolean) => void, event: FormEvent<HTMLFormElement>) => {
  event.preventDefault();
  try {
    console.log("handleLogin called");
    const response = await axios.post('http://localhost:8080/api/user/login', {
      username,
      password
    });
    const { data } = response.data;
    const { access_token } = data;
    console.log("Login successful, token:", access_token);
    localStorage.setItem('accessToken', access_token);
    localStorage.setItem('username', username);
    setIsLoggedIn(true);
  } catch (error) {
    console.error('Login error:', error);
  }
};

export const handleLogout = (setIsLoggedIn: (value: boolean) => void) => {
  localStorage.removeItem('accessToken');
  localStorage.removeItem('username');
  setIsLoggedIn(false);
};
