import type { AppProps } from 'next/app'
import { SessionProvider } from "next-auth/react"
import { useEffect } from 'react'

import '@/styles/global-styles.css'
export default function App({ Component, pageProps }: AppProps) {

  useEffect(() => {
    const handleVisibilityChange = () => {
      if (document.visibilityState === 'hidden') {
        if (window.confirm("Are you sure you want to close the browser?")) {
          localStorage.clear();
        }
      }
    };

    document.addEventListener('visibilitychange', handleVisibilityChange);

    return () => {
      document.removeEventListener('visibilitychange', handleVisibilityChange);
    };
  }, []);

  return (
      <SessionProvider session={pageProps.session}>
          <Component {...pageProps} />
      </SessionProvider>

  )
}