package ws

import "net"

type ClientInterface interface {
	Close(reason string)

	ReadPump()

	WritePump()

	SendMessage(m []byte)

	ID() string
}

type CStruct struct {
	Id   string
	Conn net.Conn
}

type Hub struct {
	RegisterChan chan ClientInterface

	CRegister   chan CStruct
	CUnregister chan CStruct

	UnregisterChan chan ClientInterface

	BroadCastChan chan []byte

	ClientBroadCastChan chan []byte

	Clients    map[string]ClientInterface
	Controller map[string]net.Conn
}

func NewHub() *Hub {
	return &Hub{
		RegisterChan:   make(chan ClientInterface, 10),
		UnregisterChan: make(chan ClientInterface, 10),

		CRegister:   make(chan CStruct, 10),
		CUnregister: make(chan CStruct, 10),

		Clients: make(map[string]ClientInterface),

		BroadCastChan:       make(chan []byte),
		ClientBroadCastChan: make(chan []byte),

		Controller: make(map[string]net.Conn),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.RegisterChan:
			h.Clients[client.ID()] = client
		case client := <-h.UnregisterChan:
			delete(h.Clients, client.ID())
		case client := <-h.CRegister:
			h.Controller[client.Id] = client.Conn
		case client := <-h.CUnregister:
			delete(h.Controller, client.Id)
		case b := <-h.BroadCastChan:
			for _, c := range h.Controller {
				c.Write(b)
			}
		case b := <-h.ClientBroadCastChan:
			for _, ci := range h.Clients {
				ci.SendMessage(b)
			}
		}
	}
}
