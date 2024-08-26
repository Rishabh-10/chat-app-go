# chat-app-go

a basic chat app made using websockets 

## Packages used

- github.com/gorilla/websocket v1.5.3

## Run Locally

Install dependencies

```bash
go mod tidy
```

Start the server

```bash
go run main.go
```

Start the UI using liveServer or directly open the register.html file on chrome
- Enter your username to register
- Now you can send messages to the hub
- You can open one more window to mimic some other user

## Additional Tools

To test the app you would require a websocket tool as curl is not sufficient for a full websocket connection

```bash
npm install -g wscat
```

Connect to the websocket using wscat

```bash
wscat -c ws://localhost:8001/ws -o ws://localhost
```
