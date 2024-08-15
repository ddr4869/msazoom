import axios from 'axios';

const instance = axios.create({
  baseURL: 'https://'+process.env.NEXT_PUBLIC_HOST,
  timeout: 5000, // Set a timeout for requests (in milliseconds)
  headers: {
    'Content-Type': 'application/json'
  },
});

export default instance;