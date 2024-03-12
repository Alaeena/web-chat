'use client';
import React, { useEffect, useState } from 'react';

import { CreateRoom, GetRooms } from '@/services/wsService';
import type { Room } from '@/services/wsService';
import Rooms from '@/components/Rooms';

function Home() {
  const [value, setValue] = useState('');
  const [rooms, setRooms] = useState<Room[]>([]);

  async function getRooms() {
    const rooms = await GetRooms();
    if (rooms) {
      setRooms(rooms);
    }
  }
  useEffect(() => {
    getRooms();
  }, []);

  async function handleSubmit(e: React.SyntheticEvent) {
    e.preventDefault();
    if(!value){
      return
    }
    const success = await CreateRoom(value);
    if (success) {
      getRooms();
    }
    setValue('');
  }
  return (
    <div className="my-8 px-4 w-full h-full">
      <div className="flex justify-center mt-3 p-5">
        <input
          className="border border-grey2 p-2 rounded-md focus:outline-none focus:border-blue"
          placeholder="tên phòng"
          value={value}
          onChange={(e) => setValue(e.target.value)}
        ></input>
        <button onClick={handleSubmit} className="bg-blue border text-white rounded-md p-2 ml-1">
          Tạo phòng
        </button>
      </div>
      <div className="mt-6">
        <Rooms rooms={rooms} />
      </div>
    </div>
  );
}

export default Home;
