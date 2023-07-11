package service

import (
	"log"

	"github.com/irvantaufik28/go-hotel-project/domain"
)

type roomServiceImpl struct {
	RoomRepository domain.RoomRepository
}

func NewRoomService(roomRepository domain.RoomRepository) domain.RoomService {
	return &roomServiceImpl{
		RoomRepository: roomRepository,
	}
}

func (service *roomServiceImpl) GetAllRoom() ([]*domain.RoomRes, error) {
	room, err := service.RoomRepository.GetAllRoom()

	if err != nil {
		return nil, err
	}
	var rooms []*domain.RoomRes

	for _, v := range room {
		var room = domain.RoomRes{
			Id:         v.Id,
			Code:       v.Code,
			TypeRoomID: v.TypeRoomID,
			Status:     v.Status,
			CreatedAt:  v.CreatedAt,
			UpdatedAt:  v.UpdatedAt,
		}
		rooms = append(rooms, &room)
	}

	return rooms, nil
}

func (service *roomServiceImpl) GetById(id int64) (*domain.RoomRes, error) {
	room, err := service.RoomRepository.GetById(id)

	if err != nil {
		return nil, err
	}
	res := &domain.RoomRes{
		Id:         room.Id,
		Code:       room.Code,
		TypeRoomID: room.TypeRoomID,
		Status:     room.Status,
		CreatedAt:  room.CreatedAt,
		UpdatedAt:  room.UpdatedAt,
	}
	return res, nil
}

func (service *roomServiceImpl) Create(payload *domain.RoomReq) error {
	room := &domain.Room{
		Code:       payload.Code,
		TypeRoomID: payload.TypeRoomID,
		Status:     payload.Status,
	}

	err := service.RoomRepository.Create(room)
	if err != nil {
		return err
	}

	return nil
}

func (service *roomServiceImpl) Update(id int64, payload *domain.RoomReq) error {
	room := &domain.Room{
		Code:       payload.Code,
		TypeRoomID: payload.TypeRoomID,
		Status:     payload.Status,
	}

	err := service.RoomRepository.Update(id, room)
	if err != nil {
		return err
	}
	return nil
}

func (service *roomServiceImpl) Delete(id int64) error {
	log.Printf("Deleting room with ID: %d", id)

	err := service.RoomRepository.Delete(id)
	if err != nil {
		log.Printf("Failed to delete room: %v", err)
		return err
	}

	log.Printf("Room with ID %d deleted successfully", id)
	return nil
}
