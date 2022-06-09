use websocket between server (go) and frontend (react), react render pretty json if recieve this message

the server open other simple webhook, where recieve structs and notify to frontend

use POST

PORT_WS is port web_socket
PORT_SERVER is webhook

```go
var structWithTags  str
s, _ := json.Marshal(structWithTags)
_, err = http.Post("http://localhost:"+PORT_SERVER, "application/json", bytes.NewReader(js))
```


```yaml
services:
  backend:
    image: ignaciomagno/pretty-struct-golang-backend:v0.0
    environment:
      - PORT_SERVER=8090
      - PORT_WS=5050
    ports:
      - 8090:8090
      - 5050:5050
    container_name: pretty-struct-golang-backend
  frontend:
    image: ignaciomagno/pretty-struct-golang-frontend:v0.0
    environment:
      REACT_APP_BACKEND_IP: 192.168.1.99
      REACT_APP_PORT_WS: 5050
    ports:
      - "3001:3000"
    container_name: pretty-struct-golang-frontend
```