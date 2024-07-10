
import axios from './axios'
// create board
export const createChatAxios = (username:string, chat_title:string, password:string) => {
  try {
    return new Promise<any>((resolve, reject) => {
      const reqUrl = process.env.NEXT_PUBLIC_CHAT_SERVICE+`/chat/create?title=${chat_title}&username=${username}&password=${password}`;
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
      const reqUrl = process.env.NEXT_PUBLIC_CHAT_SERVICE+'/chat';
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
      const reqUrl = process.env.NEXT_PUBLIC_CHAT_SERVICE+'/chat/' + chat_id;
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
      const reqUrl = process.env.NEXT_PUBLIC_CHAT_SERVICE+'/chat/random';
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

export const checkChatPasswordAxios = (chat_id:string, password:string) => {
  try {
    return new Promise<any>((resolve, reject) => {
      const reqUrl = process.env.NEXT_PUBLIC_CHAT_SERVICE+'/chat/check_password';
      const token = localStorage.getItem('accessToken')
      axios.post(reqUrl, {
        chat_id: chat_id,
        password: password
      }, {
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