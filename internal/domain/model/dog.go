package model

type CuteDog struct {
	Id DogID
	Name string
	Kind string
}

type DogID int64