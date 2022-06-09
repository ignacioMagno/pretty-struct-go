import './App.css';
import JSONPretty from 'react-json-pretty';
import { useCallback, useEffect, useRef, useState } from 'react';

var backend = process.env.REACT_APP_BACKEND_IP
var ws = process.env.REACT_APP_PORT_WS
function App() {
  const [json, setJson] = useState<JSON>()
  const [connected, setConnected] = useState(false)
  const socket = useRef<WebSocket | undefined>()

  useEffect(() => {
    console.log("ws://" + backend + ":" + ws + "/reverse")
    socket.current = new WebSocket("ws://" + backend + ":" + ws + "/reverse")
    socket.current.onmessage = function (e: any) {
      setJson(e.data)
    }

    socket.current.onopen = function () {
      setConnected(true)
    }

    window.onbeforeunload = function (e) {
      console.log("cerrando desde window")
      socket.current?.close(-1)
    };

    return () => {

      socket.current?.send("-1")
      console.log("cerrando")
      socket.current?.close(-1)
    }
  }, [])

  return (
    <div className="">
      {connected ? <>Conectado</> : <>No conectado</>}
      {socket.current?.readyState}
      {json &&
        <JSONPretty data={json} ></JSONPretty>
      }
    </div>
  );
}

export default App;
