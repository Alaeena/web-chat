'use client';

import React, { ReactNode, createContext, useEffect, useState } from 'react';
import { Storage } from '@/utils';
import { useRouter } from 'next/navigation';

export type UserInfo = {
  username: string;
  id: string;
};

export const AuthContext = createContext<{
  login: boolean;
  setLogin: (auth: boolean) => void;
  user: UserInfo;
  setUser: (user: UserInfo) => void;
}>({
  login: false,
  setLogin: () => {},
  user: { username: '', id: '' },
  setUser: () => {},
});
const AuthProvider = ({ children }: { children: ReactNode }) => {
  const router = useRouter();
  const [login, setLogin] = useState(false);
  const [user, setUser] = useState<UserInfo>({ username: '', id: '' });

  useEffect(() => {
    const userInfo = Storage.get();

    if (userInfo.username) {
      setUser(userInfo);
      setLogin(true);
    } else {
      const path = window.location.pathname;
      if (path != '/login' && path != '/register') {
        router.push('/login');
        return;
      }
    }
  }, []);

  return (
    <AuthContext.Provider
      value={{
        login: login,
        setLogin: setLogin,
        user: user,
        setUser: setUser,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
};

export default AuthProvider;
