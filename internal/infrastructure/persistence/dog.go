package persistence

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/keitakn/golang-grpc-server/internal/domain/model"
	"github.com/keitakn/golang-grpc-server/internal/domain/repository"
	"github.com/keitakn/golang-grpc-server/internal/infrastructure/dto"
)

type dogPersistence struct{
	db *gorm.DB
}

func NewDogPersistence(db *gorm.DB) repository.CuteDogRepository {
	db.AutoMigrate(&dto.Dog{})
	return &dogPersistence{
		db,
	}
}

func (dp dogPersistence) AddCuteDog(ctx context.Context, newDog *model.CuteDog) (model.DogID, error) {
	dbDog := dto.ConvertCuteDog(newDog)
	result := dp.db.Create(dbDog)
	return model.DogID(dto.WrapDogResult(result).ID), result.Error
}

func (dp dogPersistence) FindCuteDog(ctx context.Context, id model.DogID) (*model.CuteDog, error) {
	var foundDog dto.Dog
	result := dp.db.First(&foundDog, int64(id))
	return dto.AdaptDog(dto.WrapDogResult(result)), result.Error
}

func (dp dogPersistence) UpdateCuteDog(ctx context.Context, updateDog *model.CuteDog) (*model.CuteDog, error) {
	var foundDog dto.Dog
	dp.db.First(&foundDog, int64(updateDog.Id))
	dto.UpdateCuteDog(&foundDog, updateDog)
	result := dp.db.Save(&foundDog)
	return dto.AdaptDog(dto.WrapDogResult(result)), result.Error
}

func (dp dogPersistence) DeleteCuteDog(ctx context.Context, id model.DogID) (model.DogID, error) {
	var foundDog dto.Dog
	dp.db.First(&foundDog, int64(id))
	result := dp.db.Delete(&foundDog)
	return model.DogID(dto.WrapDogResult(result).ID), result.Error
}
