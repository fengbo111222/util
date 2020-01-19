package main

import (
	"fmt"
	"net"
	"os"
)

const (
	SERVER_IP       = "192.168.1.200"
	SERVER_PORT     = "6001"

	//SERVER_IP       = "127.0.0.1"
	//SERVER_PORT     = "10006"
	SERVER_RECV_LEN = 50
)
func main() {
	address := SERVER_IP + ":" + SERVER_PORT
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer conn.Close()

	for {
		// Here must use make and give the lenth of buffer
		data := make([]byte, SERVER_RECV_LEN)
		_, _, err := conn.ReadFromUDP(data)
		//_, rAddr, err := conn.ReadFromUDP(data)
		if err != nil {
			fmt.Println(err)
			continue
		}

		strData := string(data)
		fmt.Println("Received:", strData)

		//_, err = conn.WriteToUDP([]byte(strData), rAddr)
		//if err != nil {
		//	fmt.Println(err)
		//	continue
		//}
		//
		//fmt.Println("Send:", strData)
	}
}
