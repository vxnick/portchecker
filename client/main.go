package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("You must specify an IP, protocol and port to connect to")
		os.Exit(1)
	}

	ipaddr := os.Args[1]
	protocol := os.Args[2]
	port, _ := strconv.Atoi(os.Args[3])

	switch protocol {
	case "udp", "UDP":
		addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", ipaddr,
			port))

		sock, err := net.DialUDP("udp", nil, addr)
		if err != nil {
			fmt.Printf("Unable to connect to %s:%d (%s)\n", ipaddr, port,
				protocol)
			os.Exit(1)
		}
		defer sock.Close()

		for i := 1; i <= 5; i++ {
			fmt.Printf("Sending request %d of 5... ", i)

			_, err = sock.Write([]byte("Ping"))

			if err != nil {
				fmt.Println("failed")
			} else {
				fmt.Println("success")
			}

			time.Sleep(time.Second)
		}
	case "tcp", "TCP":
		for i := 1; i <= 5; i++ {
			addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", ipaddr,
				port))

			sock, err := net.DialTCP("tcp", nil, addr)
			if err != nil {
				fmt.Printf("Unable to connect to %s:%d (%s)\n", ipaddr, port,
					protocol)
				os.Exit(1)
			}
			defer sock.Close()

			fmt.Printf("Sending request %d of 5... ", i)

			_, err = sock.Write([]byte("Ping"))

			if err != nil {
				fmt.Println("failed")
			} else {
				fmt.Println("success")
			}
			sock.Close()

			time.Sleep(time.Second)
		}
	default:
		fmt.Println("Please specify either UDP or TCP for the protocol")
		os.Exit(1)
	}
}
