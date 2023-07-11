package repository

import (
	"github.com/irvantaufik28/go-hotel-project/domain"
	"gorm.io/gorm"
)

type TypeRoomRepository struct {
	db *gorm.DB
}

func NewTypeRoomRepository(database *gorm.DB) *TypeRoomRepository {
	return &TypeRoomRepository{
		db: database,
	}
}

func (typeRoomRepo *TypeRoomRepository) GetAllTypeRoom() ([]*domain.TypeRoom, error) {
	var typeRoom []*domain.TypeRoom
	err := typeRoomRepo.db.Preload("Room").Find(&typeRoom).Error

	if err != nil {
		return nil, err
	}
	return typeRoom, nil
}
