package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Cozzytree/comtroller/internal/server/ws"
)

type ServerStruct struct {
	Hub *ws.Hub
}

func InitServer(hub *ws.Hub) *http.Server {
	ss := &ServerStruct{
		Hub: hub,
	}

	s := http.Server{
		Addr:    fmt.Sprintf(":%v", os.Getenv("PORT")),
		Handler: ss.RegisterRoutes(),
	}
	return &s
}
