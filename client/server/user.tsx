import axios from './axios'

export const LoginAxios = (username, password) => {
    //noStore()
    try {
      return new Promise<any>((resolve, reject) => {
        const reqUrl = '/user/login';
        axios.post(reqUrl,  {
            username: username,
            password: password
        })
        .then(res => {
          resolve(res.data.data);
          const { data } = res.data;
          const { access_token } = data;
          localStorage.setItem('accessToken', access_token);
          localStorage.setItem('username', username);
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

// signupAxios
export const SignupAxios = (username, password) => {
    //noStore()
    try {
      return new Promise<any>((resolve, reject) => {
        const reqUrl = '/user';
        axios.post(reqUrl,  {
            username: username,
            password: password
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

export const CheckFriendAxios = (token:string, friend_name: string) => {
    //noStore()
    try {
      return new Promise<any>((resolve, reject) => {
        const reqUrl = '/user/friend/check?friend=' + friend_name;
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

export const AddFriendAxios = (token:string, friend_name: string) => {
    //noStore()
    try {
      return new Promise<any>((resolve, reject) => {
        const reqUrl = '/user/friend';
        axios.post(reqUrl, {
          friend: friend_name
        }, {
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

export const GetFriendsAxios = (token:string) => {
    //noStore()
    try {
      return new Promise<any>((resolve, reject) => {
        const reqUrl = '/user/friend';
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

export const RemoveFriendAxios = (token:string, friend_id: string) => {
    //noStore()
    try {
      return new Promise<any>((resolve, reject) => {
        const reqUrl = '/user/friend?friend=' + friend_id;
        axios.delete(reqUrl,  {
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