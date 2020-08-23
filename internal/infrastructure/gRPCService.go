package infrastructure

import (
	"context"
	"errors"
	"github.com/itohiro73/golang-grpc-ddd-server/pkg/pb/cat"
)

type CatService struct{}

func (s *CatService) FindCuteCat(ctx context.Context, message *cat.FindCuteCatMessage) (*cat.CuteCatResponse, error) {
	switch message.CatId {
	case "moko":
		// もこはチンチラシルバー
		return &cat.CuteCatResponse{
			Name: "Moko",
			Kind: "Chinchilla silver",
		}, nil
	case "mop":
		// もっぷはマンチカン
		return &cat.CuteCatResponse{
			Name: "Mop",
			Kind: "Munchkin",
		}, nil
	}

	return nil, errors.New("Not Found Cat")
}
