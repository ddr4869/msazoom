import React, { useState } from 'react';
import axios from 'axios';
import { useRouter } from "next/router";

export default function LoginForm() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const router = useRouter();
  const handleSubmit = async (e:any) => {
    e.preventDefault();
    try {
      const response = await axios.post('http://localhost:8080/api/user/login', {
        username,
        password
      });
      console.log("resp -> ", response)
      if (response && response.status === 200) {
        const { data } = response.data;
        const { access_token } = data;
        localStorage.setItem('accessToken', access_token);
        localStorage.setItem('username', username);
        router.push("/");
      }
      console.log("test")
    } catch (error) {
      console.error('로그인 에러:', error);
      alert('로그인 실패!');
    }
  };

  return (
    <div>
      <h1>Login Page</h1>
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="username">Username:</label>
          <input
            type="text"
            id="username"
            name="username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            required
          />
        </div>
        <div>
          <label htmlFor="password">Password:</label>
          <input
            type="password"
            id="password"
            name="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </div>
        <button type="submit">Login</button>
      </form>
    </div>
  );
}
