import axios from './axios'

export const getPodsAxios = () => {
    //noStore()
    try {
      return new Promise<any>((resolve, reject) => {
        const reqUrl = '/pods';
        axios.get(reqUrl,  {})
        .then(res => {
          resolve(res.data);
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

export const getPodLogsAxios = (name: string, namespace: string) => {
    //noStore()
    try {
      return new Promise<any>((resolve, reject) => {
        const reqUrl = `/pod/${namespace}/log`;
        axios.get(reqUrl,  { params: { name } })
        .then(res => {
          resolve(res.data);
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
