import axios from './axios'

export const getFriendMessageAxios = async (friend_name: string) => {
    try {
        return new Promise<any>((resolve, reject) => {
            const reqUrl = `/message?friend_name=` + friend_name;
            const token = localStorage.getItem('accessToken') || '';
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
        }).catch((error) => {
            console.error('Server Error:', error);
            throw new Error('Failed to connect server.');
        }
        )
    } catch (error) {
        console.error('Server Error:', error);
        throw new Error('Failed to connect server.');
    }
}