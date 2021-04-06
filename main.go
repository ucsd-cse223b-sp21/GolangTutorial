package main

import (
	"fmt"
	"log"
	"os"
	"net"
	"net/http"
	"net/rpc"
)

var (
	client bool
	server bool
	serverAddr = "localhost"
	port = "1234"
)

const (
	RESPONSE_OK = 1
	RESPONSE_BAD = 2
	RESPONSE_MAGIC = 3
)

type ScoreServer int

type Message struct {
	User string
	Contents string
}

type MessageResponse struct {
	ResponseCode int
}

func (ss *ScoreServer) PutScore(m Message, mr *MessageResponse) error {
	log.Println("Made it to the Put Score")
	return nil
}

func main() {
	fmt.Println("Entry function main")

	if len(os.Args) != 2 {
		log.Fatal("Only pass one command line argument to this program")
	}

	val := os.Args[1]

	if val == "client" {
		client = true
	} else if val == "server" {
		server = true
	}

	if client {
		run_client()
	} else if server {
		run_server()
	}
}

func run_client(){
	log.Println("Starting the client")
	//addr = localhost:1234
	cli, err := rpc.DialHTTP("tcp", serverAddr+":"+port)

	if err != nil {
		log.Fatal("unable to dial: ", err)
	}

	var (
		mess Message
		messResp MessageResponse
	)

	mess.User = "stew"
	mess.Contents = "What hath god wraught"

	err = cli.Call("ScoreServer.PutScore", mess, &messResp)
	if err != nil {
		log.Fatal("Unable to call PutScore: ", err)
	}

	log.Printf("Got the response %d\n",messResp.ResponseCode)
}

func run_server() {
	log.Println("Starting the server")

	ss := new(ScoreServer)
	rpc.Register(ss)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("unable to listen on the server: ", err)
	}
	http.Serve(l, nil)
}