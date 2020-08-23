package dto

import (
	"github.com/itohiro73/golang-grpc-ddd-server/internal/domain/model"
	"github.com/jinzhu/gorm"
)

type Dog struct {
	gorm.Model
	Name string
	Kind string
}

func ConvertCuteDog(d *model.CuteDog) *Dog {
	return &Dog{
		Name: d.Name,
		Kind: d.Kind,
	}
}

func UpdateCuteDog(target *Dog, source *model.CuteDog) *Dog {
	target.Name = source.Name
	target.Kind = source.Kind
	return target
}

func AdaptDog(d *Dog) *model.CuteDog {
	return &model.CuteDog{
		Id:   model.DogID(d.ID),
		Name: d.Name,
		Kind: d.Kind,
	}
}

func WrapDogResult(result *gorm.DB) *Dog {
	return result.Value.(*Dog)
}
