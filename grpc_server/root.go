package grpcserver

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"

	metadataPb "github.com/byeoru/kania-realtime/grpc_server/metadata"
	updatesPb "github.com/byeoru/kania-realtime/grpc_server/updates"
	"github.com/byeoru/kania-realtime/util"
	"google.golang.org/grpc"
)

var (
	grpcInit sync.Once
	port     = flag.Int("port", 50051, "The server port")
)

func NewServer() {
	grpcInit.Do(func() {
		initMetadata()
		flag.Parse()
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		metadataPb.RegisterMapDataServer(s, &metadataServer{})
		updatesPb.RegisterRealtimeUpdatesServer(s, &updatesServer{})

		log.Printf("server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	})
}

func initMetadata() {
	// 1. 파일 열기
	file, err := os.Open("data/pack.json")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// 2. 파일 내용 읽기
	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// 3. JSON 데이터 파싱
	err = json.Unmarshal(byteValue, &util.Pack)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}
}
