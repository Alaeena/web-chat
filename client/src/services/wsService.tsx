'use client';
import { v4 as uuidv4 } from 'uuid';
import { HOST_URL, WS_URL } from '@/constants';
import { UserInfo } from '@/modules/auth_provider';
export type Room = {
   id: string;
   name: string;
   clients: [];
};
export type Client = {
   username: string;
};

export const GetRooms = async () => {
   const res = await fetch(`${HOST_URL}/websocket/`, {
      method: 'GET',
   });
   return await res.json();
};

export const CreateRoom = async (name: string): Promise<boolean> => {
   const res = await fetch(`${HOST_URL}/websocket/`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ id: uuidv4(), name }),
   });
   return res.ok;
};
export const GetClients = async (id: string): Promise<Client[]> => {
   const res = await fetch(`${HOST_URL}/websocket/${id}/clients`, {
      method: 'GET',
   });
   return await res.json();
};
export const JoinRoom = (roomId: string, user: UserInfo): WebSocket => {
   const connStr = `${WS_URL}/websocket/${roomId}?userId=${user.id}&username=${user.username}`;
   return new WebSocket(connStr);
};
