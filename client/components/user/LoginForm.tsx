interface LoginFormProps {
  username: string;
  setUsername: (username: string) => void;
  password: string;
  setPassword: (password: string) => void;
  handleLogin: (e: React.FormEvent) => void;
}

const LoginForm = ({ username, setUsername, password, setPassword, handleLogin }:LoginFormProps) => (
  <form onSubmit={handleLogin}>
    <input type="text" value={username} onChange={(e) => setUsername(e.target.value)} placeholder="Username" required />
    <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} placeholder="Password" required />
    <button type="submit">Login</button>
  </form>
);

export default LoginForm;