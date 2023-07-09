package service

import (
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
			Id:          v.Id,
			Code:        v.Code,
			Type:        v.Type,
			Price:       v.Price,
			Description: v.Description,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
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
		Id:          room.Id,
		Code:        room.Code,
		Type:        room.Type,
		Price:       room.Price,
		Description: room.Description,
		CreatedAt:   room.CreatedAt,
		UpdatedAt:   room.UpdatedAt,
	}
	return res, nil
}
