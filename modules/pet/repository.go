package pet

import (
	"context"
	"fmt"
	"strings"

	"github.com/muhwyndhamhp/gotes-mx/utils/errs"
	"github.com/muhwyndhamhp/gotes-mx/utils/scopes"
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

// DeletePet implements PetRepository.
func (r *repo) DeletePet(ctx context.Context, pet *Pet) error {
	if err := r.db.Delete(pet).Error; err != nil {
		return errs.Wrap(err)
	}
	return nil
}

// FetchPets implements PetRepository.
func (r *repo) FetchPets(ctx context.Context, page int, pageSize int, petType PetType, keyword string) ([]Pet, error) {
	var result []Pet

	query := r.db.Debug().WithContext(ctx).
		Scopes(scopes.Paginate(page, pageSize)).
		Preload("Images")

	if keyword != "" {
		wrappedKeyword := fmt.Sprintf("%%%s%%", strings.ToLower(keyword))
		query = query.Where(
			query.Where("lower(name) like ?", wrappedKeyword).
				Or("lower(nickname) like ?", wrappedKeyword).
				Or("lower(address1) like ?", wrappedKeyword).
				Or("lower(address2) like ?", wrappedKeyword),
		)
	}

	if petType != All {
		query.Where("pet_type = ?", petType)
	}

	err := query.
		Find(&result).
		Error
	if err != nil {
		return nil, errs.Wrap(err)
	}

	return result, nil
}

// GetPetByID implements PetRepository.
func (r *repo) GetPetByID(ctx context.Context, id uint) (*Pet, error) {
	var result Pet
	err := r.db.WithContext(ctx).
		Preload("Images").
		First(&result, id).Error
	if err != nil {
		return nil, errs.Wrap(err)
	}
	return &result, nil
}

// InsertPet implements PetRepository.
func (r *repo) InsertPet(ctx context.Context, pet *Pet) error {
	err := r.db.WithContext(ctx).Save(pet).Error
	if err != nil {
		return errs.Wrap(err)
	}
	return nil
}

// UpdatePet implements PetRepository.
func (r *repo) UpdatePet(ctx context.Context, id uint, pet *Pet) error {
	pet.ID = id
	return r.InsertPet(ctx, pet)
}

func NewPetRepository(db *gorm.DB) PetRepository {
	return &repo{
		db: db,
	}
}
