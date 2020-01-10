package infrastructure

import (
	"context"
	"errors"
	pb "github.com/keitakn/golang-grpc-server/pb"
)

type CatService struct{}

func (s *CatService) FindCuteCat(ctx context.Context, message *pb.FindCuteCatMessage) (*pb.CuteCatResponse, error) {
	switch message.CatId {
	case "moko":
		// もこはチンチラシルバー
		return &pb.CuteCatResponse{
			Name: "Moko",
			Kind: "Chinchilla silver",
		}, nil
	case "mop":
		// もっぷはマンチカン
		return &pb.CuteCatResponse{
			Name: "Mop",
			Kind: "Munchkin",
		}, nil
	}

	return nil, errors.New("Not Found Cat")
}
