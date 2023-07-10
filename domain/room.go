package domain

import "time"

type Room struct {
	Id          int64     `gorm:"primaryKey" json:"id"`
	Code        string    `gorm:"type:varchar(40)" json:"code"`
	Type        string    `gorm:"type:varchar(40)" json:"type"`
	Price       int64     `gorm:"type:int" json:"price"`
	Description string    `gorm:"type:text" json:"description"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

type RoomRes struct {
	Id          int64     `gorm:"primaryKey" json:"id"`
	Code        string    `gorm:"type:varchar(40)" json:"code"`
	Type        string    `gorm:"type:varchar(40)" json:"type"`
	Price       int64     `gorm:"type:int" json:"price"`
	Description string    `gorm:"type:text" json:"description"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

type RoomReq struct {
	Code        string `gorm:"type:varchar(40)" json:"code"`
	Type        string `gorm:"type:varchar(40)" json:"type"`
	Price       int64  `gorm:"type:int" json:"price"`
	Description string `gorm:"type:text" json:"description"`
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
