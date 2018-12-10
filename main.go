package main

import (
	"./tcpserver"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var server *tcpserver.Tcpserver = &tcpserver.Tcpserver{Port: 1234, IdleTimeout: 30, RequestLimits: 30, Endpoint: "https://www.honestbee.tw/"}

var desc string = "+---+-----------------+\n| 1 | Config Server   |\n+---+-----------------+\n| 2 | Start Server    |\n+---+-----------------+\n| 3 | Show Statistic  |\n+---+-----------------+\n| 4 | Stop Server    |\n+---+-----------------+\n| 5 | Exit            |\n+---+-----------------+\n"

func main() {
	fmt.Println(desc)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		command := input.Text()
		switch command {
		case "1":
			fmt.Println("Config Server")
			if server.Listener != nil {
				fmt.Println("Please Stop Server first")
				fmt.Println(desc)
				break
			}
			fmt.Println("type ? to get Help")
			for input.Scan() {
				config := input.Text()
				if config == "q" {
					break
				}
				switch config {
				case "1":
					fmt.Println("Config Port\n==============\nPort: ")
					input.Scan()
					portStr := input.Text()
					port, _ := strconv.Atoi(portStr)
					fmt.Printf("New Port is %+v\n", port)
					server.Port = port
					fmt.Println(" 1: config port\n 2: config Idle Timeout\n 3: config Endpoint\n q: Exist config\n")
					break
				case "2":
					fmt.Println("Config IdelTimeout\n===========\nTimeout (sec): ")
					idleTimeoutStr := input.Text()
					idleTimeout, _ := strconv.Atoi(idleTimeoutStr)
					fmt.Printf("New Idle Timeout is %+v\n", idleTimeout)
					server.IdleTimeout = idleTimeout
					fmt.Println(" 1: config port\n 2: config Idle Timeout\n 3: config Endpoint\n q: Exist config\n")
					break
				case "3":
					fmt.Println("Config Endpoint\n=============\nEndpoint: ")
					endpoint := input.Text()
					fmt.Printf("New Endpoint is %+v\n", endpoint)
					server.Endpoint = endpoint
					fmt.Println(" 1: config port\n 2: config Idle Timeout\n 3: config Endpoint\n q: Exist config\n")
					break
				case "?":
					fmt.Println(" 1: config port\n 2: config Idle Timeout\n 3: config Endpoint\n q: Exist config\n")
				}
			}
			fmt.Println(desc)
			break
		case "2":
			fmt.Println("Start Server")
			if server.Listener != nil {
				fmt.Println("Tcpserver Already Started")
				break
			}
			go server.Start()
			fmt.Println(desc)
			break
		case "3":
			fmt.Println("Show Statistic")
			server.Statistic()
			fmt.Println(desc)
			break
		case "4":
			fmt.Println("Stop Server")
			server.Stop()
			fmt.Println(desc)
			break
		case "5":
			fmt.Println("Exit")
			break
		default:
			fmt.Println("Unexpect")
			fmt.Println(desc)
		}

	}

}
