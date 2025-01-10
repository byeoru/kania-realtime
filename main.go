package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin:     func(r *http.Request) bool { return true }, // Cross-Origin 모든 요청 허용, 나중에 꼭 수정해야 함
	ReadBufferSize:  256,                                        // 읽기 버퍼 크기: 256B
	WriteBufferSize: 256,                                        // 쓰기 버퍼 크기: 256B
}

var kaniaWorldTime = time.Date(312, 5, 2, 1, 20, 0, 0, time.UTC) // kanias world 시작 시간

// WebSocket 연결을 위한 타입 정의
type Client struct {
	conn *websocket.Conn
	lock sync.Mutex
}

var clients = make(map[*websocket.Conn]*Client) // 모든 연결을 관리

func handleConnection(w http.ResponseWriter, r *http.Request) {
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

		fmt.Println("Received:", string(p))

		// 클라이언트로 메시지 보내기
		client.lock.Lock() // 특정 클라이언트에 대한 동기화 처리
		err = conn.WriteMessage(messageType, []byte("Message received"))
		client.lock.Unlock()

		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}
	}
}

func runKaniaWorldTimer(ticker *time.Ticker) {
	for range ticker.C {
		// Ticker 신호가 오면 작업 수행
		kaniaWorldTime = kaniaWorldTime.Add(1 * time.Minute)

		// 모든 클라이언트에게 시간 정보 보내기
		for _, client := range clients {
			client.lock.Lock() // 개별 클라이언트에 대한 동기화
			// 메시지 생성
			message := map[string]interface{}{
				"title": "worldTime",
				"body":  []byte(kaniaWorldTime.Format(time.RFC3339)),
			}

			// msg를 JSON으로 직렬화
			msgBytes, _ := json.Marshal(message)
			err := client.conn.WriteMessage(websocket.TextMessage, msgBytes)
			client.lock.Unlock()

			if err != nil {
				fmt.Println("Error sending message:", err)
			}
		}
	}
}

func main() {
	http.HandleFunc("/", handleConnection)
	fmt.Println("WebSocket server is listening on :8081")
	// 6초마다 작업을 수행하는 Ticker 생성 : 1초마다 kania world 1분이 지남
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	go runKaniaWorldTimer(ticker)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Error starting WebSocket server:", err)
	}
}
