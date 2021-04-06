package main

import (
	"fmt"
	"log"
	"os"
)

var (
	client bool
	server bool
)

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
}

func run_server() {
	log.Println("Starting the server")
}