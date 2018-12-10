# Tcp Server

**Description**

1. TCP server takes in any request text per line and send a query to an external API, until client send 'quit' or timed out. 
2. TCP server can accept multiple connections at the same time

**Tcpserver**

| Field       | Type   | Description                                     |
| ----------- | ------ | ----------------------------------------------- |
| Port        | int    | *Required*. the server listen port              |
| IdleTimeout | int32  | *Optional*. default = 30 sec, .                 |
| Endpoint    | string | *Optional*, default = https://www.honestbee.tw. |

**Start Server**

*tcpserver.Start()

**Stop Server**

*tcpserver.Stop()

**Show Statistic**

*tcpserver.Statistic()



