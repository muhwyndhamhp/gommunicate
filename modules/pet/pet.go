package pet

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	Name     string
	Nickname string
	Birthday *time.Time
	Race     string
	Reasons  string
	PetType  PetType
	Address1 string
	Address2 string
	Images   []Image
}

type Image struct {
	gorm.Model
	PetID uint
	URL   string
}

type PetTypeSet struct {
	PetType PetType
	URL     string
}

type PetType string

const (
	All  PetType = ""
	Cat  PetType = "cat"
	Dog  PetType = "dog"
	Bird PetType = "bird"
)

type PetRepository interface {
	FetchPets(ctx context.Context, page, pageSize int, petType PetType, keyword string) ([]Pet, error)
	InsertPet(ctx context.Context, pet *Pet) error
	UpdatePet(ctx context.Context, id uint, pet *Pet) error
	DeletePet(ctx context.Context, pet *Pet) error
	GetPetByID(ctx context.Context, id uint) (*Pet, error)
}
