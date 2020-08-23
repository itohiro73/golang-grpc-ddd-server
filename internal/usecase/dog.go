package usecase

import (
	"context"
	"github.com/itohiro73/golang-grpc-ddd-server/internal/domain/model"
	"github.com/itohiro73/golang-grpc-ddd-server/internal/domain/repository"
)

type DogUseCase interface {
	AddCuteDog(ctx context.Context, newDog *model.CuteDog) (model.DogID, error)
	FindCuteDog(ctx context.Context, dogId model.DogID) (*model.CuteDog, error)
	UpdateCuteDog(ctx context.Context, updateDog *model.CuteDog) (*model.CuteDog, error)
	DeleteCuteDog(ctx context.Context, dogId model.DogID) (model.DogID, error)
}

type dogUseCase struct {
	dogRepository repository.CuteDogRepository
}

func NewDogUseCase(dr repository.CuteDogRepository) DogUseCase {
	return &dogUseCase{
		dogRepository: dr,
	}
}

func (d dogUseCase) AddCuteDog(ctx context.Context, newDog *model.CuteDog) (model.DogID, error) {
	return d.dogRepository.AddCuteDog(ctx, newDog)
}

func (d dogUseCase) FindCuteDog(ctx context.Context, dogId model.DogID) (*model.CuteDog, error) {
	return d.dogRepository.FindCuteDog(ctx, dogId)
}

func (d dogUseCase) UpdateCuteDog(ctx context.Context, updateDog *model.CuteDog) (*model.CuteDog, error) {
	return d.dogRepository.UpdateCuteDog(ctx, updateDog)
}

func (d dogUseCase) DeleteCuteDog(ctx context.Context, dogId model.DogID) (model.DogID, error) {
	return d.dogRepository.DeleteCuteDog(ctx, dogId)
}
