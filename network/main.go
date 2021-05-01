package main

import (
	"fmt"
	"net"
	"strconv"
)

func readConnFromServer(port int) (net.Conn, error) {
	l, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		fmt.Println("Cannot listen:", err)
		return nil, err
	}
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Cannot accept:", err)
		return nil, err
	}
	return conn, nil
}

func readConnFromClient(host string, port int) (net.Conn, error) {
	conn, err := net.Dial("tcp", host+":"+strconv.Itoa(port))
	if err != nil {
		fmt.Println("Dial connect error:", err)
		return nil, err
	}
	return conn, nil
}

func writeDataByServer(conn net.Conn, s string) {
	data := []byte(s)
	_, err := conn.Write(data)
	if err != nil {
		fmt.Println("Cannot write", err)
	}
}

func readDataByClient(conn net.Conn) (int, error) {
	data := make([]byte, 1024)
	count, err := conn.Read(data)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to read:", conn.RemoteAddr())
	} else {
		fmt.Println(count, "bytes from", conn.RemoteAddr())
		fmt.Println("content:", string(data[:count]))
	}
	return count, err
}

func main() {
	// TODO
	connServer, _ := readConnFromServer(8080)
	fmt.Println(connServer.LocalAddr().String())
	defer connServer.Close()
	connClient, _ := readConnFromClient("localhost", 8080)
	defer connClient.Close()
	fmt.Println(connClient.LocalAddr().String())

	writeDataByServer(connServer, "Hello!")
	readDataByClient(connClient)
}
