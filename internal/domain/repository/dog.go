package repository

import (
	"context"
	"github.com/keitakn/golang-grpc-server/internal/domain/model"
)

type CuteDogRepository interface {
	AddCuteDog(ctx context.Context, newDog *model.CuteDog) (model.DogID, error)
	FindCuteDog(ctx context.Context, dogId model.DogID) (*model.CuteDog, error)
	UpdateCuteDog(ctx context.Context, updateDog *model.CuteDog) (*model.CuteDog, error)
	DeleteCuteDog(ctx context.Context, newDog model.DogID)(model.DogID, error)
}