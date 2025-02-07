package providers

import (
	"net/http"
	"sync"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type WebSocketService struct {
	upgrader websocket.Upgrader
	clients  map[*websocket.Conn]bool
	mu       sync.Mutex
	logger   *LoggerService
}

func NewWebSocketProvider(
	cfg *config.AppConfig,
	logger *LoggerService,
) *WebSocketService {
	return &WebSocketService{
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		clients: make(map[*websocket.Conn]bool),
		logger:  logger,
	}
}

func (ws *WebSocketService) HandleWebSocket(ctx *gin.Context) {
	conn, err := ws.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ws.logger.Error("Failed to upgrade connection to WebSocket", zap.Error(err))
		return
	}

	ws.mu.Lock()
	ws.clients[conn] = true
	ws.mu.Unlock()

	ws.logger.Info("WebSocket connection established", zap.String("remoteAddr", conn.RemoteAddr().String()))
	go ws.readMessages(conn)
}

func (ws *WebSocketService) readMessages(conn *websocket.Conn) {
	defer func() {
		ws.logger.Info("WebSocket connection closed", zap.String("remoteAddr", conn.RemoteAddr().String()))
		ws.mu.Lock()
		delete(ws.clients, conn)
		ws.mu.Unlock()
		conn.Close()
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			ws.logger.Warn("Error reading message from WebSocket", zap.Error(err))
			break
		}

		ws.logger.Info("Message received from WebSocket", zap.ByteString("message", message))
		ws.broadcastMessage(message, conn)
	}
}

func (ws *WebSocketService) broadcastMessage(message []byte, sender *websocket.Conn) {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	for client := range ws.clients {
		if client != sender {
			err := client.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				ws.logger.Warn("Failed to send message to client", zap.Error(err))
				client.Close()
				delete(ws.clients, client)
			}
		}
	}
}

func (ws *WebSocketService) SendMessageToAll(message []byte) {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	for client := range ws.clients {
		err := client.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			ws.logger.Warn("Failed to send message to client", zap.Error(err))
			client.Close()
			delete(ws.clients, client)
		}
	}
}

func (ws *WebSocketService) CloseAllConnections() {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	for client := range ws.clients {
		client.Close()
		delete(ws.clients, client)
	}

	ws.logger.Info("All WebSocket connections closed")
}
