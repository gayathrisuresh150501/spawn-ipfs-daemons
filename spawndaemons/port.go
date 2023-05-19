package spawndaemons

import (
	"fmt"
	"net"
	"os/exec"
	"strconv"
)

func FindOpenPort() int, err {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}

	tcpaddr, ok := listener.Addr().(*net.TCPAddr)
	if !ok {
		//log error
		fmt.Println("Invalid address")
	}

	return tcpaddr.Port
}

func AssignPort() {
	multiAddr := `/ip4/127.0.0.1/tcp/`
	API_port := multiAddr + strconv.Itoa(FindOpenPort())
	// fmt.Println(API_port)

	apiPort := exec.Command("ipfs", "config", "Addresses.API", API_port)
	apiPort.Run()

	Gateway_port := multiAddr + strconv.Itoa(FindOpenPort())
	// fmt.Println(Gateway_port)
	gatewayPort := exec.Command("ipfs", "config", "Addresses.Gateway", Gateway_port)
	gatewayPort.Run()
}

