
import axios from './axios'
// create board
export const createChatAxios = (username:string, chat_title:string, password:string) => {
  try {
    return new Promise<any>((resolve, reject) => {
      const reqUrl = `/chat/create?title=${chat_title}&username=${username}&password=${password}`;
      const token = localStorage.getItem('accessToken')
      axios.get(reqUrl,  {
          headers: {
            'Authorization': `Bearer ${token}`
          }
      })
      .then(res => {
        resolve(res.data.data);
      })
      .catch(err => {
          console.log(err)
        reject(err.message);
      })
    })
  } catch (error) {
    console.error('Server Error:', error);
    throw new Error('Failed to connect server.');
  }
}

export const getChatsAxios = () => {
  //noStore()
  try {
    return new Promise<any>((resolve, reject) => {
      const reqUrl = '/chat';
      const token = localStorage.getItem('accessToken')
      axios.get(reqUrl,  {
          headers: {
            'Authorization': `Bearer ${token}`
          }
        })
      .then(res => {
        resolve(res.data.data);
      })
      .catch(err => {
          console.log(err)
        reject(err.message);
      })
    })
  } catch (error) {
    console.error('Server Error:', error);
    throw new Error('Failed to connect server.');
  }
}

export const getChatAxios = (chat_id:string) => {
  //noStore()
  try {
    return new Promise<any>((resolve, reject) => {
      const reqUrl = '/chat/' + chat_id;
      const token = localStorage.getItem('accessToken')
      axios.get(reqUrl,  {
          headers: {
            // Bearer 토큰을 Authorization 헤더에 추가
            'Authorization': `Bearer ${token}`
          }
        })
      .then(res => {
        resolve(res.data.data);
      })
      .catch(err => {
          console.log(err)
        reject(err.message);
      })
    })
  } catch (error) {
    console.error('Server Error:', error);
    throw new Error('Failed to connect server.');
  }
}

export const getRandomChatIdAxios = () => {
  try {
    return new Promise<any>((resolve, reject) => {
      const reqUrl = '/chat/random';
      const token = localStorage.getItem('accessToken')
      axios.get(reqUrl, {
        headers: {
          'Authorization': `Bearer ${token}`
        }
      })
      .then(res => {
        resolve(res.data.data);
      })
      .catch(err => {
          console.log(err)
        reject(err.message);
      })
    })
  } catch (error) {
    console.error('Server Error:', error);
    throw new Error('Failed to connect server.');
  }
}