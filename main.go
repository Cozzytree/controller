package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"runtime"

	"github.com/Cozzytree/comtroller/internal/server"
	"github.com/Cozzytree/comtroller/internal/server/ws"
)

func main() {
	hub := ws.NewHub()

	go hub.Run()

	s := server.InitServer(hub)

	go func() {
		if err := s.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("TCP_PORT")))
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	log.Println("TCP server listeninig on port", os.Getenv("TCP_PORT"))

	fmt.Println(runtime.NumGoroutine())

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go func(conn net.Conn) {
			defer func() {
				conn.Close()
				hub.CUnregister <- ws.CStruct{
					Id:   "control",
					Conn: conn,
				}
			}()

			hub.CRegister <- ws.CStruct{
				Id:   "control",
				Conn: conn,
			}

			conn.Write([]byte("PING"))

			buffer := make([]byte, 1024)

			for {
				err := readMessage(conn, hub.ClientBroadCastChan, buffer)
				if err != nil {
					log.Println("connection read error:", err)
					break
				}
			}
		}(conn)
	}
}

func readMessage(c net.Conn, sendChan chan<- []byte, buffer []byte) error {
	n, err := c.Read(buffer)
	if err != nil {
		fmt.Println("error while reading", err.Error())
		return err
	}

	fmt.Println("recieved message from client", string(buffer[:n]))
	sendChan <- buffer[:n]

	return nil
}
