'use client';

import React, { useContext, useEffect } from 'react';
import { useState } from 'react';
import { useRouter } from 'next/navigation';

import { RegisterUser } from '@/services/authService';
import { AuthContext } from '@/modules/auth_provider';
import Link from 'next/link';

const Login = () => {
   const router = useRouter();
   const { login } = useContext(AuthContext);

   const [email, setEmail] = useState('');
   const [password, setPassword] = useState('');
   const [username, setUsername] = useState('');

   useEffect(() => {
      if (login) {
         router.push('/');
      }
   }, [login]);

   async function handleSubmit(e: React.SyntheticEvent) {
      e.preventDefault();

      const res = await RegisterUser(email, username, password);
      if (res) {
         router.push('/');
      }
   }
   return (
      <div className="flex items-center justify-center flex-col min-w-full min-h-screen">
         <form className="flex flex-col md:w-2/5">
            <div className="text-3xl font-bold text-center">
               <span className="text-blue">Tạo tài khoản mới!</span>
            </div>
            <input
               placeholder="tên người dùng"
               className="p-3 mt-8 rounded-md bg-grey3 border-2 border-grey1 focus:outline-none focus:border-blue"
               onChange={(e) => setUsername(e.target.value)}
            ></input>
            <input
               placeholder="email"
               className="p-3 mt-8 rounded-md bg-grey3 border-2 border-grey1 focus:outline-none focus:border-blue"
               onChange={(e) => setEmail(e.target.value)}
            ></input>
            <input
               type="password"
               placeholder="mật khẩu"
               className="p-3 mt-8 rounded-md bg-grey3 border-2 border-grey1 focus:outline-none focus:border-blue"
               onChange={(e) => setPassword(e.target.value)}
            ></input>
            <button type="submit" onClick={handleSubmit} className="p-3 mt-6 rounded-md bg-blue font-bold text-white">
               đăng ký
            </button>
         </form>
         <div className="mt-4">
            <span>Bạn đã có tài khoản?</span>
            <Link href="/login" className="mx-2 text-blue font-bold">
               đăng nhập
            </Link>
         </div>
      </div>
   );
};

export default Login;
