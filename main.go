package main

import (
	"github.com/Cozzytree/comtroller/internal/server"
)

func main() {
	s := server.InitServer()

	s.ListenAndServe()
}
