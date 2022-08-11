## WebSocket implementation

#### Websocket Server is implemented in Golang

#### Websocket Client is implemented in HTML-Javascript

Just a simple Websocket implementations.

Usage:

- Clone the repo

- run ```go mod tidy```

- Spin up the Server by running ```go run server/server.go```

- Open the client

### Server implementation

Once the client requests the Server for the HTTP Upgrade(Websocket) by hitting the endpoint ```/ws``` then the Server will upgarde the connection to Websocket if the client supports.

Then after the successfull conenction establishment there will be bidirectional communication between the Client and Server.

Our Server is implemented to send the message every 10 secs once the client is upgraded to Websocket connection and also echos back the message that been sent by the client.

### Client implementation

Once the Client is upgraded to the Websocket connection it displays the messages from the Server and also we could send the message with the help of the input text field. Also it displays the Connection status with the Server.