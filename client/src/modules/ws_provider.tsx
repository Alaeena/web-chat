'use client';

import React, { ReactNode, useState, createContext } from 'react';
export type Conn = WebSocket | null;
export const WebsocketContext = createContext<{
  conn: Conn;
  setConn: (c: Conn) => void;
}>({
  conn: null,
  setConn: () => {},
});

const WebSocketProvider = ({ children }: { children: ReactNode }) => {
  const [conn, setConn] = useState<Conn>(null);

  return (
    <WebsocketContext.Provider
      value={{
        conn: conn,
        setConn: setConn,
      }}
    >
      {children}
    </WebsocketContext.Provider>
  );
};

export default WebSocketProvider;
