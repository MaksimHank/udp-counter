package udpcount

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

func Client(addr string) (int64, error) {
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
	n, err := conn.Read(buf)
	respStr := string(buf[:n])
	respStr = strings.TrimSpace(respStr)
	resp, err := strconv.ParseInt(respStr, 10, 64)
	if err != nil {
		fmt.Printf("Error reading from server: %v", err)
		return 0, err
	}

	fmt.Printf("Received: %d", resp)

	return resp, nil
}
