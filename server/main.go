package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("You must specify a protocol and port to listen on")
		os.Exit(1)
	}

	var buf = make([]byte, 1024)
	protocol := os.Args[1]
	port, _ := strconv.Atoi(os.Args[2])

	switch protocol {
	case "udp", "UDP":
		addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", port))

		sock, err := net.ListenUDP("udp", addr)
		if err != nil {
			fmt.Printf("Unable to listen on port %d (%s)\n", port, protocol)
			os.Exit(1)
		}
		defer sock.Close()

		fmt.Printf("Listening on port %d (%s)...\n", port, protocol)

		for {
			_, _, err := sock.ReadFromUDP(buf)

			if err != nil {
				continue
			}

			fmt.Println("Received request")
		}
	case "tcp", "TCP":
		sock, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			fmt.Printf("Unable to listen on port %d (%s)\n", port, protocol)
			os.Exit(1)
		}
		defer sock.Close()

		fmt.Printf("Listening on port %d (%s)... \n", port, protocol)

		for {
			conn, err := sock.Accept()
			if err != nil {
				fmt.Println("Unable to accept: ", err.Error())
				os.Exit(1)
			}

			_, err = conn.Read(buf)
			if err != nil {
				fmt.Println("Error reading: ", err.Error())
			}

			conn.Write([]byte("Pong"))
			conn.Close()

			fmt.Println("Received request")
		}
	default:
		fmt.Println("Please specify either UDP or TCP for the protocol")
		os.Exit(1)
	}
}
