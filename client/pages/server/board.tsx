import axios from './axios'

export const getBoardsAxios = (token:string) => {
    //noStore()
    try {
      return new Promise<any>((resolve, reject) => {
        const reqUrl = '/board';
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

// create board
export const createBoardAxios = (token:string, board_name:string, board_password:string) => {
    try {
      return new Promise<any>((resolve, reject) => {
        const reqUrl = '/board';
        axios.post(reqUrl, {
            board_name: board_name,
            board_password: board_password
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


// recommend board
export const recommendBoardAxios = (board_id:number) => {
    //noStore()
    try {
      return new Promise<any>((resolve, reject) => {
        const reqUrl = `/board/recommend`;
        axios.post(reqUrl, {
           board_id: board_id
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
  