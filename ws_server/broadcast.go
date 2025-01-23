package wsserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/byeoru/kania-realtime/types"
	"github.com/gorilla/websocket"
)

// WebSocket 연결을 위한 타입 정의
type Client struct {
	conn *websocket.Conn
	lock sync.Mutex
}

var upgrader = websocket.Upgrader{
	CheckOrigin:     func(r *http.Request) bool { return true }, // Cross-Origin 모든 요청 허용, 나중에 꼭 수정해야 함
	ReadBufferSize:  256,                                        // 읽기 버퍼 크기: 256B
	WriteBufferSize: 256,                                        // 쓰기 버퍼 크기: 256B
}

var clients = make(map[*websocket.Conn]*Client) // 모든 연결을 관리
var functionMap = map[string]interface{}{}

func makeMessage(title string, id int64, body string) (msgBytes []byte) {
	// 메시지 생성
	// fmt.Println("body", body)
	message := types.Response{
		Title: title,
		ID:    id,
		Body:  []byte(body),
	}

	// msg를 JSON으로 직렬화
	msgBytes, _ = json.Marshal(message)
	return
}

func handleBroadcast(w http.ResponseWriter, r *http.Request) {
	var err error
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}
	defer conn.Close()

	client := &Client{conn: conn}
	clients[conn] = client
	fmt.Println("New WebSocket connection established")

	for {
		// 클라이언트로부터 메시지 받기
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			delete(clients, conn) // 연결 종료 시 클라이언트 맵에서 제거
			return
		}

		// 클라이언트로 메시지 보내기
		client.lock.Lock() // 특정 클라이언트에 대한 동기화 처리

		var request types.Request
		err = json.Unmarshal(p, &request)
		if err != nil {
			fmt.Println("Message format Error:", err)
			return
		}

		if fn, exists := functionMap[request.Title]; exists {
			if getWorldTimeFn, ok := fn.(func() time.Time); ok {
				response := makeMessage(request.Title, request.ID, getWorldTimeFn().Format(time.RFC3339))
				err = conn.WriteMessage(messageType, []byte(response))
				if err != nil {
					fmt.Println("Error sending message:", err)
					return
				}
			}
		}
		client.lock.Unlock()
	}
}
