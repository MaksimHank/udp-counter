package udpcount

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

func StartServer(addr string) {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		fmt.Printf("func StartServer. \nAddress error")
		return
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Printf("func StartServer. \nStarting error server")
		return
	}
	defer conn.Close()

	var counter int
	mu := sync.Mutex{}

	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			mu.Lock()
			counter++
			mu.Unlock()
		}
	}()

	fmt.Printf("func StartServer. Server has successfully started on addr %s", addr)

	buffer := make([]byte, 1024)

	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("func StartServer. Reading error!")
			continue
		}

		message := string(buffer[:n])
		mu.Lock()
		message = strconv.Itoa(counter)
		mu.Unlock()

		response := []byte(message)
		_, err = conn.WriteToUDP(response, clientAddr)
		if err != nil {
			fmt.Printf("func StartServer. Sending error!")
		}

	}
}
