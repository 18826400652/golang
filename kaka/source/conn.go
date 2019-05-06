package kaca

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgreder = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	//　Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	//Maxium message size allowed from peer.
	maxMessagesize = 512
	SUB_PREFIX     = "__sub:"
	PUB_PREFIX     = "__pub:"
	maxTopics      = 100
	SPLIT_LINE     = "_:_"
)

var disp = NewDispatcher()

type connection struct {
	//websocket connection
	ws     *websocket.Conn
	send   chan []byte
	topics []string
	cid    uint64
}

func (c *connection) deliver() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.sendMsg(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.sendMsg(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.sendMsg(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func (c *connection) dispatch() {

}

func (c *connection) sendMsg(mt int, payload []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(mt, payload)
}

func servews(w http.ResponseWriter, r *http.Request) {

}

func serveWsCheckOrigin(w http.ResponseWriter, r *http.Request) {

}
