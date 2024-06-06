import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/router';
import ChatFriendComponent from '@/components/rtc/chatFriendComponent';

export default function Page() {
    const [username, setUsername] = useState<string | null>(null);
    const { query } = useRouter();
    const friend_id = query.friend as string;

    useEffect(() => {
        const storedUsername = localStorage.getItem('username');
        if (storedUsername) {
            setUsername(storedUsername);
        }
    }, []);

    if (!username) {
        return <div>Loading...</div>;
    }

    return (
        <div>
            <ChatFriendComponent username={username} friendname={friend_id} />
            Your name is {username}, and Friend ID is {friend_id}.
        </div>
    );
}
