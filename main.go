package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Traffic struct {
	Listen	[]string
	Talk	[]Talk
}

type Talk struct {
	Proto	string
	Address string
	Port	int
	Rate    int
	Data	string
}


func main(){

	jsonData := []byte(`
	{
		"listen": ["8080","8081"],
		"talk": [
		{
			"proto": "tcp",
			"address": "127.0.0.1",
			"port": 8081,
			"rate": 1000,
			"data": "hello world"
		},
		{
			"proto": "tcp",
			"address": "127.0.0.1",
			"port": 8080,
			"rate": 1000,
			"data": "hello world"
		}
		]
  	}
	`)

	var traffic Traffic
	err := json.Unmarshal([]byte(jsonData), &traffic)
	if err != nil {
        log.Fatalf("Unable to marshal JSON due to %s", err)
    }

	for _, v := range traffic.Listen {
		fmt.Printf("starting server on port: %s\n", v)
		go tcpServer(v)
	}

	for _, v := range traffic.Talk {
		fmt.Printf("talking on port: %d\n", v.Port)
		go tcpClient("127.0.0.1", "8080", "testdata", 1000)
	}

	select{}

}
/* package main

import (
 "io"
 "net"
 "os"
)

const (
 HOST = "localhost"
 PORT = "3333"
 TYPE = "tcp"
)

func main() {
 tcpServer, err := net.ResolveTCPAddr(TYPE, HOST+":"+PORT)
 if err != nil {
  println("ResolveTCPAddr failed:", err.Error())
  os.Exit(1)
 }

 conn, err := net.DialTCP(TYPE, nil, tcpServer)
 if err != nil {
  println("Dial failed:", err.Error())
  os.Exit(1)
 }

 defer conn.Close()

 _, err = conn.Write([]byte("Ground Control To Major Tom"))
 if err != nil {
  println("Write data failed:", err.Error())
  os.Exit(1)
 }

 //new line of code
 conn.CloseWrite()

 received := make([]byte, 4096)
 for {
  println("Reading data...")
  temp := make([]byte, 4096)
  _, err = conn.Read(temp)
  if err != nil {
   if err == io.EOF {
    break
   }
   println("Read data failed:", err.Error())
   os.Exit(1)
  }
  received = append(received, temp...)
 }

 println("Received message:", string(received))

} */