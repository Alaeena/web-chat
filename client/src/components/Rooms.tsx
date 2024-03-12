import React, { useContext } from 'react';
import { useRouter } from 'next/navigation';

import { JoinRoom, type Room } from '@/services/wsService';
import { WebsocketContext } from '@/modules/ws_provider';
import { AuthContext } from '@/modules/auth_provider';

const Rooms = ({ rooms }: { rooms: Room[] }) => {
  const { setConn } = useContext(WebsocketContext);
  const { user } = useContext(AuthContext);
  const router = useRouter();

  const handleJoin = async (value: string) => {
    const ws = await JoinRoom(value, user);

    if (ws.OPEN) {
      setConn(ws);
      router.push('/chat');
    }
  };

  return (
    <div>
      {rooms.map((room, index) => (
        <div key={index} className="border border-blue mt-2 p-4 flex items-center rounded-md w-full">
          <div className="w-full">
            <div className="text-sm">Phòng chat</div>
            <div className="text-blue font-bold text-lg">{room.name}</div>
          </div>
          <div>
            <button
              className="py-1 w-32 block text-md text-white bg-blue rounded-md"
              onClick={() => handleJoin(room.id)}
            >
              tham gia
            </button>
          </div>
        </div>
      ))}
    </div>
  );
};

export default Rooms;
