package ws

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type msg struct {
	Type string `json:"type"`
	Msg  string `json:"msg"`
}

type WS_Client struct {
	Client   *websocket.Conn
	Hub      *Hub
	ClientId string
	SendChan chan []byte
}

func (c *WS_Client) Close(reason string) {
	fmt.Println(reason)
}

func (c *WS_Client) ReadPump() {
	defer func() {
		c.Client.Close()
		c.Close("Read closed")
	}()

	for {
		_, b, err := c.Client.ReadMessage()

		fmt.Println("message :", string(b))

		if err != nil {
			fmt.Println("Read error:", err)
			break
		}

		// Check if BroadCastChan is full and log if messages might be blocked
		select {
		case c.Hub.BroadCastChan <- b:
			fmt.Println("Message sent to broadcast channel successfully")
		default:
			fmt.Println("WARNING: BroadCastChan is full, message might be blocked")
			// Try to send anyway, this will block if channel is full
			c.Hub.BroadCastChan <- b
		}
	}
}

func (c *WS_Client) WritePump() {
	defer func() {
		c.Client.Close()
		c.Close("Write pump close")
	}()

	for e := range c.SendChan {
		var err error
		fmt.Println("sending message to client", string(e))
		switch string(e) {
		case "BLINK-ON":
			err = c.Client.WriteJSON(msg{Type: "blink", Msg: "on"})
		case "BLINK-OFF":
			err = c.Client.WriteJSON(msg{Type: "blink", Msg: "off"})
		default:
			fmt.Println("Unknown command")
		}

		if err != nil {
			fmt.Println("Write error:", err)
			break
		}
	}
}

func (c *WS_Client) ID() string {
	return c.ClientId
}

func (c *WS_Client) SendMessage(b []byte) {
	c.SendChan <- b
}

func InitNewClient(w http.ResponseWriter, r *http.Request, hub *Hub) {
	upgrader := websocket.Upgrader{}
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUpgradeRequired)
		return
	}

	newClient := WS_Client{
		Client:   conn,
		ClientId: r.PathValue("client"),
		Hub:      hub,
		SendChan: make(chan []byte, 100),
	}

	// read and write concurrently
	go newClient.ReadPump()
	go newClient.WritePump()

	hub.RegisterChan <- &newClient
}
