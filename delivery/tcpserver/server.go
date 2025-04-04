package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"task-manager/constant"
	"task-manager/delivery/deliveryparam"
)

func handleConnectedClient(client net.Conn) error {
	fmt.Printf("a client has connected\n")
	defer client.Close()
	var rawReader = make([]byte, 1024)

	noReadBytes, rErr := client.Read(rawReader)
	if rErr != nil {
		return fmt.Errorf("failed to read client request %v", rErr)
	}
	fmt.Printf("client has sent %d bytes and its read", noReadBytes)
	// my request should be json format and Request format
	var request = &deliveryparam.Request{}
	umErr := json.Unmarshal(rawReader[:noReadBytes], request)
	if umErr != nil {
		return fmt.Errorf("failed to parse request %v base message is %v",
			umErr, rawReader)
	}

	fmt.Printf("command is  %v \n", request.Command)
	fmt.Printf("command's meta data is   %v \n", request)
	return nil
}

func main() {
	listener, lErr := net.Listen(constant.Network, constant.NetAddr)
	if lErr != nil {
		fmt.Printf("failed to listen address %v \n", lErr)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Printf("connection established ...\n")
	for {
		connection, cErr := listener.Accept()
		if cErr != nil {
			fmt.Printf("failed to accept client %v \n", cErr)
			continue
		}
		if cErr := handleConnectedClient(connection); cErr != nil {
			fmt.Println("An error occured while handling connection", cErr)
		}
	}
}
