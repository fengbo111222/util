package main


import (
"bufio"
"fmt"
"net"
"os"
"strconv"

)
const (
	SERVER_IP       = "192.168.1.200"
	SERVER_PORT     = 6001
	SERVER_RECV_LEN = 50
)


func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	serverAddr := SERVER_IP + ":" + strconv.Itoa(SERVER_PORT)
	conn, err := net.Dial("udp", serverAddr)
	checkError(err)

	defer conn.Close()

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()//12345

		lineLen := len(line)

		n := 0
		for written := 0; written < lineLen; written += n {
			var toWrite string
			if lineLen-written > SERVER_RECV_LEN {
				toWrite = line[written : written+SERVER_RECV_LEN]
			} else {
				toWrite = line[written:]
			}

			n, err = conn.Write([]byte(toWrite))
			checkError(err)


			//msg := make([]byte, SERVER_RECV_LEN)
			//n, err = conn.Read(msg)
			//checkError(err)

			//fmt.Println("Response:", string(msg))
		}
	}
}
