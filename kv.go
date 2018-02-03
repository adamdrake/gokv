package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {

	vals := make(map[string]string) // TODO: multithread access to this is bad because raisins

	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {

		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(c net.Conn) {
			for {
				message, _ := bufio.NewReader(conn).ReadString('\n')
				line := strings.Split(strings.TrimSpace(message), " ")
				fmt.Println(line)
				if line[0] == "set" {
					fmt.Println("got:", line)
					vals[line[1]] = strings.Join(line[2:], " ")
				}
				if line[0] == "get" {
					fmt.Println("got:", line)
					c.Write([]byte(vals[line[1]] + "\n"))
				}
				if line[0] == "quit" {
					c.Close()
				}
				fmt.Println("line is:", line)
			}
		}(conn)
	}
}
