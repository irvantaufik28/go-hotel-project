package domain

import "time"

type Room struct {
	Id         int64     `gorm:"primaryKey" json:"id"`
	Code       string    `gorm:"type:varchar(40)" json:"code"`
	TypeRoomID int64     `gorm:"type:int" json:"type_room_id"`
	Status     string    `gorm:"type:varchar(40)" json:"status"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
	TypeRoom   TypeRoom  `gorm:"foreignKey:TypeRoomID" json:"type_room"`
}

type RoomRes struct {
	Id         int64     `gorm:"primaryKey" json:"id"`
	Code       string    `gorm:"type:varchar(40)" json:"code"`
	TypeRoomID int64     `gorm:"type:int" json:"type_room_id"`
	Status     string    `gorm:"type:varchar(40)" json:"status"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

type RoomReq struct {
	Code       string `gorm:"type:varchar(40)" json:"code"`
	TypeRoomID int64  `gorm:"type:int" json:"type_room_id"`
	Status     string `gorm:"type:varchar(40)" json:"status"`
}

type RoomRepository interface {
	GetAllRoom() ([]*Room, error)
	GetById(id int64) (*Room, error)
	Create(room *Room) error
	Update(id int64, room *Room) error
	Delete(id int64) error
}

type RoomService interface {
	GetAllRoom() ([]*RoomRes, error)
	GetById(id int64) (*RoomRes, error)
	Create(payload *RoomReq) error
	Update(id int64, room *RoomReq) error
	Delete(id int64) error
}
