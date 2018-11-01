package main

import (
	"flag"
	"fmt"
)

const (
	address     = "192.168.0.1"
	port        = "3333"
	serviceName = "testService"
	key         = "testKey"
	value       = "testValue"
)

func main() {
	mode := flag.String("m", "server", "")
	flag.Parse()
	switch *mode {
	case "server":
		err := serverStart()
		if err != nil {
			fmt.Println(err.Error())
		}
	case "client":
		err := clientStart()
		if err != nil {
			fmt.Println(err.Error())
		}
	default:
		fmt.Println("wrong mode")
	}
}
