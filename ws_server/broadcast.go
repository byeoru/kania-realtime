package wsserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/byeoru/kania-realtime/types"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin:     func(r *http.Request) bool { return true }, // Cross-Origin 모든 요청 허용, 나중에 꼭 수정해야 함
	ReadBufferSize:  512,                                        // 읽기 버퍼 크기: 512B
	WriteBufferSize: 512,                                        // 쓰기 버퍼 크기: 512B
}

var BroadcastCh = make(chan []byte, 10)

type WebSocketServer struct {
	clients sync.Map
}

type Client struct {
	conn   *websocket.Conn
	cancel context.CancelFunc
}

type auth struct {
	RmId int64 `json:"rm_id"`
}

func NewWebSocketServer() *WebSocketServer {
	return &WebSocketServer{
		clients: sync.Map{},
	}
}

func MakeBroadcastMessage(title string, body interface{}) (msgBytes []byte) {
	message := types.Response{
		Title: title,
		Body:  body,
	}

	msgBytes, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
	}
	return msgBytes
}

func (wss *WebSocketServer) handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}

	_, msg, err := conn.ReadMessage()
	if err != nil {
		conn.Close()
		return
	}

	var auth auth
	err = json.Unmarshal(msg, &auth)
	if err != nil {
		log.Println("json unmarshal failed:", err)
		conn.Close()
		return
	}

	if client, ok := wss.clients.Load(auth.RmId); ok {
		client.(*Client).cancel()     // 기존 고루틴 종료
		client.(*Client).conn.Close() // 기존 연결 닫기
	}

	ctx, cancel := context.WithCancel(context.Background())
	wss.clients.Store(auth.RmId, &Client{
		conn:   conn,
		cancel: cancel,
	})
	fmt.Println("New WebSocket connection established")

	go func() {
		ticker := time.NewTicker(15 * time.Second)
		defer func() {
			ticker.Stop()
			log.Println("🛑고루틴 종료")
		}()
		log.Println("🛑고루틴 생성")
		for {
			select {
			case <-ticker.C:
				err := conn.WriteMessage(websocket.PingMessage, nil) // 서버가 Ping 보냄
				if err != nil {
					// log.Println("PING 전송 실패, 연결 종료:", err)
					conn.Close()
					wss.clients.Delete(auth.RmId)
					return
				}
				log.Println("PING 전송")
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (wss *WebSocketServer) handleBroadcast() {
	for msg := range BroadcastCh {
		wss.clients.Range(func(rmId, client any) bool {
			c := client.(*Client)
			err := c.conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println("Error sending message:", err)
				// write error있을 경우 client 해제
				c.conn.Close()
				wss.clients.Delete(rmId)
			}
			return true
		})
	}
}
