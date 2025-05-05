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
		defer func() {
			conn.Close()
			hub.CUnregister <- ws.CStruct{
				Id:   "control",
				Conn: conn,
			}
		}()

		conn.Write([]byte("Hello world"))

		hub.CRegister <- ws.CStruct{
			Id:   "control",
			Conn: conn,
		}
	}
}
