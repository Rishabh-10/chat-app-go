# chat-app-go-hello

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

## Additional Tools

To test the app you would require a websocket tool as curl is not sufficient for a full websocket connection

```bash
npm install -g wscat
```

Connect

```bash
wscat -c ws://localhost:8001/ws -o ws://localhost
```
