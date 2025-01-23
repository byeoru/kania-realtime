package main

import (
	grpcserver "github.com/byeoru/kania-realtime/grpc_server"
	wsserver "github.com/byeoru/kania-realtime/ws_server"
)

func main() {
	// WebSocket 서버 시작
	go wsserver.NewServer()
	// gRPC 서버 시작
	grpcserver.NewServer()
}
