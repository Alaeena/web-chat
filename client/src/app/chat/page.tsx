/* eslint-disable react-hooks/exhaustive-deps */
'use client';

import React, { useContext, useEffect, useRef } from 'react';
import { useState } from 'react';
import { useRouter } from 'next/navigation';

import ChatBody from './chat_body';
import { WebsocketContext } from '@/modules/ws_provider';
import { AuthContext } from '@/modules/auth_provider';
import { Client, GetClients } from '@/services/wsService';

export type Message = {
   content: string;
   client_id: string;
   username: string;
   room_id: string;
   type: 'recv' | 'self';
};

const Chat = () => {
   const router = useRouter();
   const { user } = useContext(AuthContext);
   const { conn } = useContext(WebsocketContext);

   const [users, setUsers] = useState<Client[]>([]);
   const [messages, setMessage] = useState<Array<Message>>([]);
   const textarea = useRef<HTMLTextAreaElement>(null);

   const getUsers = async (roomId: string) => {
      const values = await GetClients(roomId);
      if (values) {
         setUsers(values);
      }
   };

   useEffect(() => {
      if (conn === null) {
         router.push('/');
         return;
      }

      //change this to url parameter
      const roomId = conn.url.split('/')[5];
      getUsers(roomId);
   }, []);

   useEffect(() => {
      if (conn === null) {
         router.push('/');
         return;
      }
      conn.onmessage = (res) => {
         const m: Message = JSON.parse(res.data);
         if (m.content == 'A new user has joined the room') {
            setUsers([...users, { username: m.username }]);
         }

         if (m.content == 'user left the chat') {
            const deleteUser = users.filter((user) => user.username != m.username);
            setUsers([...deleteUser]);
            setMessage([...messages, m]);
            return;
         }

         user?.username == m.username ? (m.type = 'self') : (m.type = 'recv');
         setMessage([...messages, m]);
      };
   }, [textarea, messages, conn, users]);

   const sendMessage = () => {
      if (!textarea.current?.value) return;
      if (conn === null) {
         router.push('/');
         return;
      }

      conn.send(textarea.current.value);
      textarea.current.value = '';
   };

   return (
      <div className="flex flex-col w-full">
         <div className="p-4 md:mx-6 mb-14">
            <ChatBody data={messages} />
         </div>
         <div className="fixed bottom-0 mt-4 w-full">
            <div className="flex md:flex-row px-4 py-2 bg-grey1 md:mx-4 rounded-md">
               <div className="flex w-full mr-4 rounded-md border border-blue">
                  <textarea
                     ref={textarea}
                     placeholder="type your message here"
                     className="w-full h-10 p-2 rounded-md focus:outline-none"
                     style={{ resize: 'none' }}
                  />
               </div>
               <div className="flex items-center">
                  <button className="p-2 rounded-md bg-blue text-white" onClick={sendMessage}>
                     Send
                  </button>
               </div>
            </div>
         </div>
      </div>
   );
};

export default Chat;
