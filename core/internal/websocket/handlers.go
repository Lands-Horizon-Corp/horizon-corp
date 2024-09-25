package websocket

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Adjust this in production to check origins
		return true
	},
}

// ServeWs handles WebSocket requests from clients
func ServeWs(hub *Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		clientID := uuid.NewString()
		client := &Client{
			ID:   clientID,
			Conn: conn,
			Send: make(chan []byte, 256),
		}

		hub.Register <- client

		// Allow collection of memory referenced by the caller by doing all work in new goroutines
		go client.writePump()
		go client.readPump(hub)
	}
}

func (c *Client) readPump(hub *Hub) {
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()
	c.Conn.SetReadLimit(512)
	for {
		_, _, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}
		// We don't handle incoming messages in this example
	}
}

func (c *Client) writePump() {
	defer c.Conn.Close()
	for {
		message, ok := <-c.Send
		if !ok {
			// The hub closed the channel
			c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}
		c.Conn.WriteMessage(websocket.TextMessage, message)
	}
}
