package repository

import (
	"log"

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

func (roomRepo *RoomRepository) GetById(id int64) (*domain.Room, error) {
	var room domain.Room
	err := roomRepo.db.Preload("TypeRoom").Where("id = ?", id).Take(&room).Error
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (repo *RoomRepository) Create(room *domain.Room) error {
	err := repo.db.Create(room).Error
	if err != nil {
		return err
	}
	return nil
}

func (roomRepo *RoomRepository) Update(id int64, updatedRoom *domain.Room) error {
	room := &domain.Room{}
	err := roomRepo.db.First(room, id).Error
	if err != nil {
		return err
	}

	room.Code = updatedRoom.Code
	room.TypeRoomID = updatedRoom.TypeRoomID
	room.Status = updatedRoom.Status

	err = roomRepo.db.Save(room).Error
	if err != nil {
		return err
	}

	return nil
}

func (roomRepo *RoomRepository) Delete(id int64) error {
	log.Printf("Deleting room with ID: %d", id)

	var room domain.Room
	err := roomRepo.db.First(&room, id).Error
	if err != nil {
		log.Printf("Failed to find room: %v", err)
		return err
	}

	err = roomRepo.db.Delete(&room).Error
	if err != nil {
		log.Printf("Failed to delete room: %v", err)
		return err
	}

	log.Printf("Room with ID %d deleted successfully", id)
	return nil
}
