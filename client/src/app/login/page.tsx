'use client';

import React, { useContext, useEffect } from 'react';
import { useState } from 'react';
import { useRouter } from 'next/navigation';

import { LoginUser } from '@/services/authService';
import { AuthContext } from '@/modules/auth_provider';
import Link from 'next/link';

const Login = () => {
   const router = useRouter();
   const { login } = useContext(AuthContext);

   const [email, setEmail] = useState('');
   const [password, setPassword] = useState('');

   useEffect(() => {
      if (login) {
         router.push('/');
      }
   }, [login]);

   async function handleSubmit(e: React.SyntheticEvent) {
      e.preventDefault();

      const res = await LoginUser(email, password);
      if (res) {
         router.push('/');
      }
   }
   return (
      <div className="flex items-center justify-center flex-col min-w-full min-h-screen">
         <form className="flex flex-col md:w-2/5">
            <div className="text-3xl font-bold text-center">
               <span className="text-blue">Xin chào!!!</span>
            </div>
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
               đăng nhập
            </button>
         </form>
         <div className="mt-4">
            <span>Bạn chưa có tài khoản?</span>
            <Link href="/register" className="mx-2 text-blue font-bold">
               tạo tài khoản
            </Link>
         </div>
      </div>
   );
};

export default Login;
