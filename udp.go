package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func udpServer() {
	arguments := os.Args
	if len(arguments) == 1 {
			fmt.Println("Please provide a port number!")
			return
	}
	PORT := ":" + arguments[1]

	s, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
			fmt.Println(err)
			return
	}

	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
			fmt.Println(err)
			return
	}

	defer connection.Close()
	buffer := make([]byte, 1024)
	rand.Seed(time.Now().Unix())

	for {
			n, addr, err := connection.ReadFromUDP(buffer)
			fmt.Print("-> ", string(buffer[0:n-1]))

			if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
					fmt.Println("Exiting UDP server!")
					return
			}

			data := []byte(strconv.Itoa(random(1, 1001)))
			fmt.Printf("data: %s\n", string(data))
			_, err = connection.WriteToUDP(data, addr)
			if err != nil {
					fmt.Println(err)
					return
			}
	}
}

func udpClient() {
        arguments := os.Args
        if len(arguments) == 1 {
                fmt.Println("Please provide a host:port string")
                return
        }
        CONNECT := arguments[1]

        s, err := net.ResolveUDPAddr("udp4", CONNECT)
        c, err := net.DialUDP("udp4", nil, s)
        if err != nil {
                fmt.Println(err)
                return
        }

        fmt.Printf("The UDP server is %s\n", c.RemoteAddr().String())
        defer c.Close()

        for {
                reader := bufio.NewReader(os.Stdin)
                fmt.Print(">> ")
                text, _ := reader.ReadString('\n')
                data := []byte(text + "\n")
                _, err = c.Write(data)
                if strings.TrimSpace(string(data)) == "STOP" {
                        fmt.Println("Exiting UDP client!")
                        return
                }

                if err != nil {
                        fmt.Println(err)
                        return
                }

                buffer := make([]byte, 1024)
                n, _, err := c.ReadFromUDP(buffer)
                if err != nil {
                        fmt.Println(err)
                        return
                }
                fmt.Printf("Reply: %s\n", string(buffer[0:n]))
        }
}