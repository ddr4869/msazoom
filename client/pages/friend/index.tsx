import React, { useRef } from 'react'
import { useRouter } from 'next/router'

export default function Page() {
    // get friend id from url
    const { query } = useRouter()
    const friend_id = query.friend
    const webSocketRef = useRef();
    const username = localStorage.getItem('username')
    console.log('friend:', friend_id)
    return (
        <div>
            Your name is {username}, and Friend ID is {friend_id}.
        </div>
    );
}