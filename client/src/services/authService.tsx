'use client';

import { HOST_URL } from '@/constants';
import { Storage } from '@/utils';

export const LoginUser = async (email: string, password: string): Promise<boolean> => {
   const res = await fetch(`${HOST_URL}/auth/login`, {
      method: 'POST',
      credentials: 'include',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email, password }),
   });
   const data = await res.json();

   if (res.ok) {
      Storage.set({
         id: data.id,
         username: data.username,
      });
   }
   return res.ok;
};
export const RegisterUser = async (email: string, username: string, password: string): Promise<boolean> => {
   const res = await fetch(`${HOST_URL}/auth/register`, {
      method: 'POST',
      credentials: 'include',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email, password, username }),
   });
   const data = await res.json();

   if (res.ok) {
      Storage.set({
         id: data.id,
         username: data.username,
      });
   }
   return res.ok;
};
export const LogoutUser = async () => {
   await fetch(`${HOST_URL}/auth/logout`, {
      credentials: 'include',
      headers: { 'Content-Type': 'application/json' },
      method: 'POST',
   });
};
