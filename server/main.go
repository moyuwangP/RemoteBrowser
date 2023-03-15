package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"server/utils"
	"strings"
)

const (
	daemonListenAddress = "127.0.0.1:9876"
	heartbeatPort       = 11451
	clientPort          = 11450
)

func main() {
	args := os.Args
	argc := len(args)
	switch argc {
	case 0:
	case 1:
		fmt.Printf("usage: %s \n"+
			"        -d to start a daemon\n"+
			"        -r to register this app as default browser\n"+
			"        -s link to send url to remote computer", args[0])
	default:
		option := args[1]
		switch option {
		case "-r":
			utils.Reg()
		case "-d":
			daemon(args[2:])
		case "-s":
			send(args[2:], daemonListenAddress)
		}
	}

}

type ip struct {
	ip string
}

func daemon(args []string) {
	udpAddr, err := net.ResolveUDPAddr("udp", daemonListenAddress)
	utils.PanicIfNeeded(err)
	udpConn, err := net.ListenUDP("udp", udpAddr)
	utils.PanicIfNeeded(err)
	defer udpConn.Close()
	fmt.Println("daemon created")

	clientIp := ip{""}
	go listenClientIpChanges(&clientIp)

	for {
		lenBuf := make([]byte, 4)
		_, err := io.ReadAtLeast(udpConn, lenBuf, 4)
		utils.PanicIfNeeded(err)
		urlSize := int(binary.LittleEndian.Uint32(lenBuf))
		buf := make([]byte, urlSize)
		_, err = io.ReadAtLeast(udpConn, buf, urlSize)
		utils.PanicIfNeeded(err)

		if clientIp.ip != "" {
			send([]string{string(buf)}, fmt.Sprintf("%s:%d", clientIp.ip, clientPort))
		}
	}

}

func listenClientIpChanges(ip *ip) {
	conn, err := net.ListenPacket("udp", fmt.Sprintf(":%d", heartbeatPort))
	utils.PanicIfNeeded(err)
	defer conn.Close()

	buffer := make([]byte, 1)
	for {
		_, addr, err := conn.ReadFrom(buffer)
		utils.PanicIfNeeded(err)
		incomingIp := extractIp(addr.String())
		if incomingIp != ip.ip {
			ip.ip = incomingIp
			fmt.Printf("client ip changed to %s\n", ip.ip)
		}
	}
}

func extractIp(address string) string {
	segments := strings.Split(address, ":")
	ip := ""
	for i := 0; i < len(segments)-1; i++ {
		ip += segments[i]
	}
	return ip
}

func send(args []string, dest string) {
	if len(args) == 0 {
		fmt.Printf("url missing")
	}
	url := args[0]

	conn, err := net.Dial("udp", dest)
	utils.PanicIfNeeded(err)
	bytes := []byte(url)
	defer conn.Close()
	conn.Write(intToBytes(len(bytes)))
	conn.Write(bytes)
}

func intToBytes(i int) []byte {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(i))
	return bs
}
