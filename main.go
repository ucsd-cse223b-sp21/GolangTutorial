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
	serverAddr = "137.110.222.47"
	//serverAddr = "localhost"
	port = "1234"
)

var messageCounter map[string]int

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
	_ , ok := messageCounter[m.Contents]
	if !ok {
		messageCounter[m.Contents] = 1
	} else {
		messageCounter[m.Contents] = messageCounter[m.Contents] + 1
	}

	for key, value := range messageCounter {
		log.Println(key + ":" + fmt.Sprintf("%d",value))
	}
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

    //strange allocation, name exists
	messageCounter = make(map[string]int, 5)
	http.Serve(l, nil)
}