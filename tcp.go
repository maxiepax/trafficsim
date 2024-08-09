package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func tcpServer(port string) {
	port = ":" + port
	l, err := net.Listen("tcp", port)
	if err != nil {
			fmt.Println(err)
			return
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
			fmt.Println(err)
			return
	}

	for {
			data, err := bufio.NewReader(c).ReadString('\n')
			if err != nil {
					fmt.Println(err)
					return
			}

			fmt.Print("-> ", string(data))
			t := time.Now()
			time := t.Format(time.RFC3339)
			c.Write([]byte(time + " returned: " + data))
	}
}

func tcpClient(address string, port string, data string, rate int) {


	c, err := net.Dial("tcp", address+":"+port)
	if err != nil {
			fmt.Println(err)
			return
	}

	for {
		fmt.Fprintf(c, data+"\n")

		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: " + message)
		time.Sleep(time.Duration(rate * int(time.Millisecond)))
	}
}