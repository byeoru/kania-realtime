package grpcserver

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/byeoru/kania-realtime/grpc_server/updates"
	"github.com/byeoru/kania-realtime/types"
	wsserver "github.com/byeoru/kania-realtime/ws_server"
)

type updatesServer struct {
	pb.UnimplementedRealtimeUpdatesServer
}

func (s *updatesServer) UpdateSectorOwnership(_ context.Context, in *pb.UpdateSectorOwnershipRequest) (*emptypb.Empty, error) {
	title := "UpdateSector"
	body := &types.UpdateSector{
		Sector:     in.Sector,
		OldRealmID: in.OldRealmId,
		NewRealmID: in.NewRealmId,
		ActionType: in.ActionType,
		ActionId:   in.ActionId,
	}
	wsserver.BroadcastCh <- wsserver.MakeBroadcastMessage(title, body)
	return &emptypb.Empty{}, nil
}
