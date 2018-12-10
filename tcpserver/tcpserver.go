package tcpserver

import (
	"fmt"
	"net"
	"strconv"
	"sync/atomic"
	"time"
)

type Tcpserver struct {
	ConnectionCount  int32
	Endpoint         string
	IdleTimeout      int
	Listener         net.Listener
	Port             int
	ProcessedRequest int64
	RequestLimits    int32
	RequestRates     int
}

func (server *Tcpserver) Start() {
	fmt.Printf("Start Tcp Server at port: %d\n", server.Port)
	fmt.Printf("Idle Timeout  %d sec\n", server.IdleTimeout)
	fmt.Printf("HTTP Endpoint %s \n", server.Endpoint)

	go server.RestRequestLimits()

	listener, err := net.Listen("tcp", ":"+strconv.Itoa(server.Port))
	if err != nil {
		fmt.Println(err)
	}
	defer listener.Close() 

	server.Listener = listener
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			break
		}
		go handleConn(conn, server)
	}
}

func (server *Tcpserver) Statistic() {
	fmt.Println("=============================================Tcp Server Statistic=====================================")
	fmt.Println("time\t\t\t\t\t\t\t\tCurrent Connection\tProcessed Request\tRequest Rate")
	fmt.Printf("%s\t\t%d\t\t\t%d\t\t\t%d\n", time.Now(), server.ConnectionCount, server.ProcessedRequest, server.RequestRates)
	fmt.Printf("RequestLimits %d\n", server.RequestLimits)
	fmt.Println("======================================================================================================")
}

func (server *Tcpserver) RestRequestLimits() {
	requestLimits := server.RequestLimits
	for {
		time.Sleep(10 * time.Second)
		atomic.StoreInt32(&server.RequestLimits, requestLimits)
	}
}

func (server *Tcpserver) Stop() {
	fmt.Println("Stop Tcp Server")
	if server.Listener != nil {
		server.Listener.Close()
		server.Listener = nil
	}
}
