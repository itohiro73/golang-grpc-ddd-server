package handler

import (
	"context"
	"errors"
	"github.com/itohiro73/golang-grpc-ddd-server/internal/domain/model"
	"github.com/itohiro73/golang-grpc-ddd-server/internal/usecase"
	"github.com/itohiro73/golang-grpc-ddd-server/pkg/pb/dog"
	"strconv"
)

type dogHandler struct {
	dogUseCase usecase.DogUseCase
}

func NewDogHandler(du usecase.DogUseCase) dog.DogServer {
	return &dogHandler{
		du,
	}
}

func (dh *dogHandler) AddCuteDog(ctx context.Context, newDog *dog.CuteDog) (*dog.CuteDogID, error) {
	cuteDog := model.CuteDog{
		model.DogID(newDog.Id),
		newDog.Name,
		newDog.Kind,
	}
	newDogId, err := dh.dogUseCase.AddCuteDog(ctx, &cuteDog)
	if err != nil {
		return &dog.CuteDogID{Id: int64(newDogId)}, errors.New("AddCuteDog failed for dog: " + newDog.Name + "/" + newDog.Kind)
	}
	return &dog.CuteDogID{Id: int64(newDogId)}, err
}

func (dh *dogHandler) FindCuteDog(ctx context.Context, dogID *dog.CuteDogID) (*dog.CuteDog, error) {
	aDog, err := dh.dogUseCase.FindCuteDog(ctx, model.DogID(dogID.Id))
	if err != nil {
		return nil, errors.New("FindCuteDog failed for id: " + strconv.FormatInt(dogID.Id, 10))
	}
	pbDog := &dog.CuteDog{Id: int64(aDog.Id), Name: aDog.Name, Kind: aDog.Kind}
	return pbDog, err
}

func (dh *dogHandler) UpdateCuteDog(ctx context.Context, updateDog *dog.CuteDog) (*dog.CuteDog, error) {
	cuteDog := model.CuteDog{
		model.DogID(updateDog.Id),
		updateDog.Name,
		updateDog.Kind,
	}
	updatedDog, err := dh.dogUseCase.UpdateCuteDog(ctx, &cuteDog)
	if err != nil {
		return &dog.CuteDog{Id: int64(updatedDog.Id), Name: updatedDog.Name, Kind: updatedDog.Kind}, errors.New("UpdateCuteDog failed for updateDog: " + updateDog.Name + "/" + updateDog.Kind)
	}
	return &dog.CuteDog{Id: int64(updatedDog.Id), Name: updatedDog.Name, Kind: updatedDog.Kind}, err
}

func (dh *dogHandler) DeleteCuteDog(ctx context.Context, dogID *dog.CuteDogID) (*dog.CuteDogID, error) {
	deletedDogId, err := dh.dogUseCase.DeleteCuteDog(ctx, model.DogID((dogID.Id)))
	if err != nil {
		return &dog.CuteDogID{Id: int64(deletedDogId)}, errors.New("DeleteCuteDog failed for id: " + string(dogID.Id))
	}
	return &dog.CuteDogID{Id: int64(deletedDogId)}, err
}
