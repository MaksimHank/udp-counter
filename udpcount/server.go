package udpcount

import (
	"fmt"
	"net"
)

func StartServer(addr string) {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		fmt.Printf("func StartServer. \nAddress error")
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Printf("func StartServer. \nStarting error server")
	}
	defer conn.Close()

	fmt.Printf("func StartServer. Server has successfully started on addr %s", addr)

	//TODO: реализовать счётчик

	buffer := make([]byte, 1024)

	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("func StartServer. Reading error!")
			continue
		}
		message := string(buffer[:n])

		response := []byte(message)
		_, err = conn.WriteToUDP(response, clientAddr)
		if err != nil {
			fmt.Printf("func StartServer. Sending error!")
		}

	}
}
