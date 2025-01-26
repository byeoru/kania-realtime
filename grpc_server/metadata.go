package grpcserver

import (
	"context"
	"fmt"
	"math"

	pb "github.com/byeoru/kania-realtime/grpc_server/metadata"
	"github.com/byeoru/kania-realtime/util"
)

type metadataServer struct {
	pb.UnimplementedMapDataServer
}

type PointData struct {
	X float64
	Y float64
}

func (s *metadataServer) GetDistance(_ context.Context, in *pb.GetDistanceRequest) (*pb.GetDistanceReply, error) {
	originCenter := findCellCenter(getPackPolygon(in.GetOrigin()))
	targetCenter := findCellCenter(getPackPolygon(in.GetTarget()))
	distance := calculateDistance(originCenter, targetCenter)
	return &pb.GetDistanceReply{Distance: distance}, nil
}

func findCellCenter(points []PointData) PointData {
	// Sum all x and y coordinates
	var sumX, sumY float64
	for _, point := range points {
		sumX += point.X
		sumY += point.Y
	}

	// Calculate the averages
	xCenter := float64(sumX) / float64(len(points))
	yCenter := float64(sumY) / float64(len(points))

	return PointData{X: xCenter, Y: yCenter}
}

func getPackPolygon(i int32) []PointData {
	return util.Map(util.Pack.Cells.Cells.V[i], func(v int) PointData {
		point := util.Pack.Cells.Vertices.P[v]
		return PointData{X: float64(point[0]), Y: float64(point[1])}
	})
}

func calculateDistance(originCenter, targetCenter PointData) float64 {
	distance := math.Sqrt(math.Pow(targetCenter.X-originCenter.X, 2) + math.Pow(targetCenter.Y-originCenter.Y, 2))
	return distance
}

func (s *metadataServer) GetSectorInfo(_ context.Context, in *pb.GetSectorInfoRequest) (*pb.GetSectorInfoReply, error) {
	province := util.Pack.Cells.Cells.Province[fmt.Sprint(in.Sector)]
	population := util.Pack.Cells.Cells.Pop[fmt.Sprint(in.Sector)]
	return &pb.GetSectorInfoReply{Province: province, Population: population}, nil
}
