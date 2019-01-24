package websocket

/*
########websocket

websocket是HTML5重要特性，它实现了基于浏览器的远程Socket。使用浏览器和服务器可以进行全双工通信

说明：websocket出现前，为实现即时通信，采用技术为"轮询"(浏览器不间断向服务器发送http request请求)

websocket采用了特殊报头，使得浏览器和服务器只需要做一个握手动作，即可以浏览器和服务器间建立一条连接通道

websocket解决了web实时化问题，相比传统http好处：
1。一个Web客户端建立一个tcp连接
2。websocket服务端可以推送数据到web客户端
3。有更加轻量级的头，减少数据传送量
*/