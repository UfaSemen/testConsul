package main

import (
	"fmt"
	"net"
	"strconv"

	"github.com/hashicorp/consul/api"
)

func serverStart() error {
	portInt, err := strconv.Atoi(port)
	if err != nil {
		return fmt.Errorf("wrong port value: %s", err.Error())
	}
	err = serviceRegister(portInt, serviceName, map[string]string{key: value})
	if err != nil {
		return err
	}
	fmt.Println("servise registered")

	l, err := net.Listen("tcp", address+":"+port)
	if err != nil {
		return fmt.Errorf("error listening: %s", err.Error())
	}
	defer l.Close()
	fmt.Println("tcp server started")

	for {
		conn, err := l.Accept()
		if err != nil {
			return fmt.Errorf("error accepting: %s", err.Error())
		}
		go handleRequest(conn)
	}
}

func serviceRegister(p int, serName string, meta map[string]string) error {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return fmt.Errorf("error creating client: %s", err.Error())
	}

	agent := client.Agent()
	serviceDef := api.AgentServiceRegistration{
		Name: serName,
		Port: p,
		Meta: meta,
	}
	err = agent.ServiceRegister(&serviceDef)
	if err != nil {
		return fmt.Errorf("error regestring service: %s", err.Error())
	}

	/*kv := client.KV()
	p := &api.KVPair{Key: key, Value: []byte(value)}
	_, err = kv.Put(p, nil)
	if err != nil {
		return fmt.Errorf("error putting in kv", err.Error())
	} */
	return nil
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("error reading:", err.Error())
		return
	}
	if string(buf[0:reqLen]) == key {
		_, err = conn.Write([]byte(value))
		if err != nil {
			fmt.Println("error writing:", err.Error())
			return
		}
		fmt.Println("value sent")
	}

}
