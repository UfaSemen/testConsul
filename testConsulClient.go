package main

import (
	"fmt"
	"net"
	"strconv"

	"github.com/hashicorp/consul/api"
)

func clientStart() error {
	services, err := serviceDiscovery(serviceName)
	if err != nil {
		return err
	}
	serVal := services[0].Service.Meta[key]
	fmt.Println(serVal)

	serAddr := services[0].Node.Address
	if services[0].Service.Address != "" {
		serAddr = services[0].Service.Address
	}
	conn, err := net.Dial("tcp", serAddr+":"+strconv.Itoa(services[0].Service.Port))
	if err != nil {
		return fmt.Errorf("error connecting to server: %s", err.Error())
	}

	_, err = fmt.Fprintf(conn, key)
	if err != nil {
		return fmt.Errorf("error writing to server: %s", err.Error())
	}
	buf := make([]byte, 1024)
	reqLen, err := conn.Read(buf)
	if err != nil {
		return fmt.Errorf("error reading from server: %s", err.Error())
	}
	tcpVal := string(buf[0:reqLen])
	fmt.Println(tcpVal)
	if serVal == tcpVal {
		fmt.Println("success")
	}
	return nil
}

func serviceDiscovery(serName string) ([]*api.ServiceEntry, error) {
	cli, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return nil, fmt.Errorf("error creating client: %s", err.Error())
	}

	cat := cli.Health()
	services, _, err := cat.Service(serName, "", true, nil)
	if err != nil {
		return nil, fmt.Errorf("error discovering service: %s", err.Error())
	}
	if len(services) == 0 {
		return nil, fmt.Errorf("error discovering service: no service %s in system", serName)
	}
	/*kv := cli.KV()
	pair, _, err := kv.Get(key, nil)
	if err != nil {
		return nil, fmt.Errorf("error getting from kv", err.Error())
	}*/
	return services, nil
}
