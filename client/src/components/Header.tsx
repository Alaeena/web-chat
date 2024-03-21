'use client';
import React from 'react';
import { useContext } from 'react';
import { useRouter } from 'next/navigation';

import { AuthContext } from '@/modules/auth_provider';
import { WebsocketContext } from '@/modules/ws_provider';
import { LogoutUser } from '@/services/authService';
import { Storage } from '@/utils';

const Header = () => {
   const router = useRouter();
   const { user, login, setUser, setLogin } = useContext(AuthContext);
   const { conn, setConn } = useContext(WebsocketContext);

   if (!login) {
      return;
   }
   const handleClick = async () => {
      LogoutUser();
      setUser({ username: '', id: '' });
      setLogin(false);
      Storage.set({ username: '', id: '' });
      router.push('/login');
   };
   const handleReturn = async () => {
      setConn(null);
      router.push('/');
   };
   return (
      <div className="flex justify-between px-4 py-2 bg-grey2 w-100">
         {(conn && (
            <button onClick={handleReturn} className="cursor-point text-lg">
               {'< Trở lại'}
            </button>
         )) || <ul></ul>}
         <div>
            <span>{user.username},</span>
            <strong onClick={handleClick} className="ml-2 text-blue cursor-pointer">
               Đăng xuất
            </strong>
         </div>
      </div>
   );
};

export default Header;
