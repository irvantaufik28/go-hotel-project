package repository

import (
	"github.com/irvantaufik28/go-hotel-project/domain"
	"gorm.io/gorm"
)

type RoomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(database *gorm.DB) *RoomRepository {
	return &RoomRepository{
		db: database,
	}
}

func (roomRepo *RoomRepository) GetAllRoom() ([]*domain.Room, error) {
	var room []*domain.Room
	err := roomRepo.db.Find(&room).Error

	if err != nil {
		return nil, err
	}
	return room, nil
}
