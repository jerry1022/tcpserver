package tcpserver

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

var mutex = &sync.Mutex{}

func handleConn(conn net.Conn, server *Tcpserver) {
	respChan := make(chan interface{}, 30)
	errChan := make(chan interface{}, 30)
	defer func() {
		fmt.Println("Close Connection")
		atomic.AddInt32(&server.ConnectionCount, -1)
		close(respChan)
		close(errChan)
		conn.Close()
	}()

	atomic.AddInt32(&server.ConnectionCount, 1)

	input := bufio.NewScanner(conn)
	for input.Scan() {
		conn.SetDeadline(time.Now().Add(time.Duration(server.IdleTimeout) * time.Second))
		lineRecv := input.Text()
		if lineRecv == "quit" {
			return
		}
		atomic.AddInt32(&server.RequestLimits, -1)
		if CheckRequestLimits(server.RequestLimits) {
			go CallExternalAPI(server, lineRecv, respChan, errChan)
		}
		select {
		case resp, ok := <-respChan:
			if !ok {
				continue
			}
			atomic.AddInt64(&server.ProcessedRequest, 1)
			result := fmt.Sprintf("select channel work: %v", resp)
			fmt.Fprintln(conn, result)
		case err := <-errChan:
			fmt.Fprintln(conn, err)
		default:
			fmt.Println("nothing")
		}

	}

}

func CheckRequestLimits(requestLimits int32) bool {
	return requestLimits > 0
}

func CallExternalAPI(server *Tcpserver, query string, respChan chan interface{}, errChan chan interface{}) {
	if server.Endpoint != "" {
		// TODO: added the query string
		resp, err := http.Get(server.Endpoint)
		if err != nil {
			errChan <- err
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		respChan <- string(body)
	} else {
		respChan <- query
	}
}
