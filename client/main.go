package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
	"task-manager/constant"
	"task-manager/delivery/deliveryparam"
)

func parseCommand(cmd string) (deliveryparam.Request, error) {
	scanner := bufio.NewScanner(os.Stdin)
	switch strings.ToLower(cmd) {
	case "register":
		fmt.Println("please enter your full name")
		scanner.Scan()
		name := scanner.Text()
		fmt.Println("please enter email address")
		scanner.Scan()
		email := scanner.Text()
		fmt.Println("please enter your password")
		scanner.Scan()
		password := scanner.Text()

		return deliveryparam.Request{
			Command: "register",
			MetaData: map[string]string{
				"name":     name,
				"email":    email,
				"password": password,
			},
		}, nil

	case "login":
		fmt.Println("please enter your email address")
		scanner.Scan()
		email := scanner.Text()
		fmt.Println("please enter password")
		scanner.Scan()
		password := scanner.Text()

		return deliveryparam.Request{
			Command: "login",
			MetaData: map[string]string{
				"email":    email,
				"password": password,
			},
		}, nil

	default:
		return deliveryparam.Request{}, fmt.Errorf("not valid command")
	}
}

func handleRequest(request deliveryparam.Request) error {
	fmt.Println("connecting to server...")
	connection, Cerr := net.Dial(constant.Network, constant.NetAddr)
	if Cerr != nil {
		println("could not connect ", Cerr)

		os.Exit(1)
	}
	defer connection.Close()
	fmt.Println("connected to server!")

	rawRequest, mErr := json.Marshal(request)
	if mErr != nil {
		println("could not marshal the request: ", request.Command, Cerr)

		os.Exit(1)
	}
	noOfWrittenBytes, wErr := connection.Write(rawRequest)
	if wErr != nil {
		fmt.Printf("could not write on connection %v \n", wErr)
	}
	fmt.Printf("%d bytes has written for server successfully", noOfWrittenBytes)
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("please enter a command")
		scanner.Scan()
		command := scanner.Text()
		req, pErr := parseCommand(command)
		if pErr != nil {
			fmt.Println(pErr)
			os.Exit(1)
		}
		handleRequest(req)
	}
}
