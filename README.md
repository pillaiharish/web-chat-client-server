# web-chat-client-server
Here the IP address is my mac IP on which the golang binary is running. The clients can connect to the internal IP with port 8000 to start connection. Connecting to http://192.168.0.148:8000 will create a new websocket connection. I have implemented a counter which will update the connected client on FCFS basis. The client app will broadcast messages and will be received by other client only if the websocket connection is established prior to messaging. This is a connection-oriented, stateful protocol that operates over TCP (Transmission Control Protocol). WebSocket is a protocol designed for full-duplex communication over a long-lived TCP connection. It is stateful, enabling ongoing interaction and data exchange between a client and a server.

## Client 1
![Screeshot for client1](https://github.com/pillaiharish/web-chat-client-server/blob/main/client1.png)

## Client 2
![Screeshot for client1](https://github.com/pillaiharish/web-chat-client-server/blob/main/client2.jpg)

## Client 3
![Screeshot for client1](https://github.com/pillaiharish/web-chat-client-server/blob/main/client3.jpg)