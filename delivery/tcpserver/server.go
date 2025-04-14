package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"task-manager/constant"
	"task-manager/delivery/deliveryparam"
	"task-manager/repository"
	"task-manager/service"
)

func handleConnectedClient(client net.Conn) error {
	fmt.Printf("a client has connected\n")
	defer client.Close()
	var rawReader = make([]byte, 1024)

	noReadBytes, rErr := client.Read(rawReader)
	if rErr != nil {
		return fmt.Errorf("failed to read client request %v", rErr)
	}
	fmt.Printf("client has sent %d bytes and its read\n", noReadBytes)
	// my request should be json format and Request format
	var request = &deliveryparam.Request{}
	umErr := json.Unmarshal(rawReader[:noReadBytes], request)
	if umErr != nil {
		return fmt.Errorf("failed to parse request %v base message is %v",
			umErr, rawReader)
	}

	handleClientRequest(request)
	fmt.Printf("command is  %v \n", request.Command)
	fmt.Printf("command's meta data is   %v \n", request.MetaData)
	return nil
}

func handleClientRequest(req *deliveryparam.Request) error {
	switch req.Command {
	case "register":
		name, nameOk := req.MetaData["name"]
		email, emailOk := req.MetaData["email"]
		password, passwordOk := req.MetaData["password"]
		if !(passwordOk && nameOk && emailOk) {
			return fmt.Errorf("not enough meta data passed")
		}
		rUser, rErr := userService.Register(service.CreateUserRequest{
			Email:    email,
			Name:     name,
			Password: password,
		})
		if rErr != nil {
			fmt.Println("an error", rErr)
		} else {
			fmt.Printf("created %+v\n", rUser)
		}
	case "login":
		email, emailOk := req.MetaData["email"]
		password, passwordOk := req.MetaData["password"]
		if !(passwordOk && emailOk) {
			return fmt.Errorf("not enough meta data passed")
		}
		userID := userService.Login(service.ValidateUserRequest{
			Email:    email,
			Password: password,
		})
		if userID.ValidatedID == 0 {
			return fmt.Errorf("login failed!")
		}
		validatedUserID = userID.ValidatedID
	}
	return nil
}

var userRepo = repository.NewUserStorage()
var userService = service.NewUserService(&userRepo)
var validatedUserID = 0

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
