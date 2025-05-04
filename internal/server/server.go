package server

import "net/http"

type ServerStruct struct {
}

func InitServer() *http.Server {
	s := http.Server{}
	return &s
}
