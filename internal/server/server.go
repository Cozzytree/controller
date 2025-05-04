package server

import (
	"fmt"
	"net/http"
	"os"
)

type ServerStruct struct {
}

func InitServer() *http.Server {
	ss := &ServerStruct{}

	s := http.Server{
		Addr:    fmt.Sprintf(":%v", os.Getenv("PORT")),
		Handler: ss.RegisterRoutes(),
	}
	return &s
}
