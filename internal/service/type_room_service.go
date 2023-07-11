package service

import (
	"github.com/irvantaufik28/go-hotel-project/domain"
)

type typeRoomServiceImp struct {
	TypeRoomRepository domain.TypeRoomRepository
}

func NewTypeRoomService(typeRoomRespository domain.TypeRoomRepository) domain.TypeRoomService {
	return &typeRoomServiceImp{
		TypeRoomRepository: typeRoomRespository,
	}
}

func (service *typeRoomServiceImp) GetAllTypeRoom() ([]*domain.TypeRoomRes, error) {
	typeRoom, err := service.TypeRoomRepository.GetAllTypeRoom()

	if err != nil {
		return nil, err
	}

	var typeRooms []*domain.TypeRoomRes

	for _, v := range typeRoom {
		var typeRoom = domain.TypeRoomRes{
			Id:          v.ID,
			Name:        v.Name,
			Price:       v.Price,
			Description: v.Description,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
			Room:        v.Room,
		}
		typeRooms = append(typeRooms, &typeRoom)
	}
	return typeRooms, nil
}
