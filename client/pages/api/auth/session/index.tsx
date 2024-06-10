import { getSession } from 'next-auth/react';

export default async (req:any, res:any) => {
  const session = await getSession({ req });

  if (session) {
    // 사용자 세션이 있을 경우
    res.status(200).json(session);
  } else {
    // 사용자 세션이 없을 경우
    res.status(401).json({ message: 'Not authenticated' });
  }
};
