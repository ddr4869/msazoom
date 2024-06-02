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
