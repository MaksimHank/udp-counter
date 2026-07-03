package udpcount

import (
	"fmt"
	"net"
	"time"
)

func Client(addr string) (int, error) {
	conn, err := net.Dial("udp", addr)
	if err != nil {
		fmt.Printf("func Client. Connection error: %v", err)
		return 0, err
	}
	defer conn.Close()

	message := "Send request to udp-server"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Printf("func Client. Error sending message: %v", err)
		return 0, err
	}
	fmt.Println("Sent: ", message)

	conn.SetReadDeadline(time.Now().Add(5 * time.Second))

	buf := make([]byte, 1024)
	resp, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("Error reading from server: %v", err)
		return 0, err
	}

	fmt.Printf("Received: %s", string(buf[:resp]))

	return resp, nil
}
