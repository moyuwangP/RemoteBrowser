package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"runtime"
	"time"
)

const (
	heartbeatPort = 11451
	clientPort    = 11450
)

func main() {
	args := os.Args
	argc := len(args)

	switch argc {
	case 0:
	case 1:
		fmt.Printf("usage: %s address \n", args[0])
	default:
		go heartbeat(args[1])
		listen(args[1])
	}

}

func listen(ip string) {
	udpAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("0.0.0.0:%d", clientPort))
	panicIfNeeded(err)
	udpConn, err := net.ListenUDP("udp", udpAddr)
	panicIfNeeded(err)

	defer udpConn.Close()
	for {
		lenBuf := make([]byte, 4)
		n, err := io.ReadAtLeast(udpConn, lenBuf, 4)
		panicIfNeeded(err)
		if n != 4 {
			panic("length not correct")
		}
		fmt.Println(lenBuf)
		size := int(binary.LittleEndian.Uint32(lenBuf))

		fmt.Printf("size: %d\n", size)
		contentBuf := make([]byte, size)
		n, err = io.ReadAtLeast(udpConn, contentBuf, size)
		panicIfNeeded(err)
		if n != size {
			panic("read exits before finished")
		}

		openInBrowser(string(contentBuf))
	}
}

func heartbeat(dest string) {
	conn, err := net.Dial("udp", fmt.Sprintf("%s:%d", dest, heartbeatPort))
	panicIfNeeded(err)
	defer conn.Close()

	for {
		bytes := []byte{0}
		_, _ = conn.Write(bytes)
		time.Sleep(60 * time.Second)
	}
}

func openInBrowser(url string) {
	fmt.Printf("open: %s\n", url)
	var command *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		command = exec.Command("xdg-open", url)
	case "windows":
		command = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		command = exec.Command("open", url)
	default:
		fmt.Errorf("OS not supported")
		return
	}
	err := command.Start()
	if err != nil {
		fmt.Errorf(err.Error())
	}
}

func panicIfNeeded(err error) {
	if err != nil {
		panic(err)
	}
}
