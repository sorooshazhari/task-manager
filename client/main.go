package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"task-manager/constant"
	"task-manager/delivery/deliveryparam"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("please enter a test command")
	scanner.Scan()
	test_command := scanner.Text()
	fmt.Printf("connecting to server...\n")
	connection, Cerr := net.Dial(constant.Network, constant.NetAddr)
	if Cerr != nil {
		println("could not connect ", Cerr)
		os.Exit(1)
	}
	fmt.Printf("connected to server\n")
	newRequest := deliveryparam.Request{
		Command: test_command,
	}
	rawRequest, mErr := json.Marshal(newRequest)

	if mErr != nil {
		println("could not marshal the request: ", newRequest.Command, Cerr)
		os.Exit(1)
	}
	fmt.Println(rawRequest)
	noOfWrittenBytes, wErr := connection.Write(rawRequest)
	if wErr != nil {
		fmt.Printf("could not write on connection %v \n", wErr)
	}

	fmt.Printf("%d bytes has written for server successfully", noOfWrittenBytes)
	connection.Close()

}
