import axios from './axios'
export const LoginAxios = (username:string, password:string) => {
    //noStore()
    try {
      return new Promise<any>((resolve, reject) => {
        const reqUrl = process.env.NEXT_PUBLIC_USER_SERVICE+'/user/login';
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

export const GuestLoginAxios = () => {
  try {
    return new Promise<any>((resolve, reject) => {
      const reqUrl = process.env.NEXT_PUBLIC_USER_SERVICE+'/user/login/non-member';
      axios.post(reqUrl)
        .then(res => {
          resolve(res.data.data);
          const { data } = res.data;
          const { id, access_token } = data;
          localStorage.setItem('accessToken', access_token);
          localStorage.setItem('username', id);  // 비회원 로그인 시 생성된 ID
        })
        .catch(err => {
          console.log(err);
          reject(err.message);
        });
    });
  } catch (error) {
    console.error('Server Error:', error);
    throw new Error('Failed to connect to the server.');
  }
};

// signupAxios
export const SignupAxios = (username:string, password:string) => {
    //noStore()
    try {
      return new Promise<any>((resolve, reject) => {
        const reqUrl = process.env.NEXT_PUBLIC_USER_SERVICE+'/user';
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

export const CheckFriendAxios = (friend_name: string) => {
    //noStore()
    try {
      return new Promise<any>((resolve, reject) => {
        const reqUrl = process.env.NEXT_PUBLIC_USER_SERVICE+'/user/friend/check?friend=' + friend_name;
        const token = localStorage.getItem('accessToken') || '';
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

export const AddFriendAxios = (friend_name: string) => {
    //noStore()
    try {
      return new Promise<any>((resolve, reject) => {
        const reqUrl = process.env.NEXT_PUBLIC_USER_SERVICE+'/user/friend';
        const token = localStorage.getItem('accessToken') || '';
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

export const GetFriendsAxios = () => {
    //noStore()
    try {
      return new Promise<any>((resolve, reject) => {
        const reqUrl = process.env.NEXT_PUBLIC_USER_SERVICE+'/user/friend';
        const token = localStorage.getItem('accessToken') || '';
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

export const RemoveFriendAxios = (friend_id: string) => {
    //noStore()
    try {
      return new Promise<any>((resolve, reject) => {
        const reqUrl = process.env.NEXT_PUBLIC_USER_SERVICE+'/user/friend?friend=' + friend_id;
        const token = localStorage.getItem('accessToken') || '';
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

export const GetFollowerAxios = () => {
    //noStore()
    try {
      return new Promise<any>((resolve, reject) => {
        const reqUrl = process.env.NEXT_PUBLIC_USER_SERVICE+'/user/friend/follower';
        const token = localStorage.getItem('accessToken') || '';
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