package main

import (
	"log"

	"github.com/Cozzytree/comtroller/internal/server"
)

func main() {
	s := server.InitServer()

	log.Println("Server started on port", s.Addr)
	log.Fatal(s.ListenAndServe())
}
