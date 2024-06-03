import axios from 'axios';

const instance = axios.create({
  baseURL: 'http://192.168.0.6:8080/api', // api endpoint env..
  timeout: 5000, // Set a timeout for requests (in milliseconds)
  headers: {
    'Content-Type': 'application/json'
  },
});

export default instance;