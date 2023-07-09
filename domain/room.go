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

type RoomRepository interface {
	GetAllRoom() ([]*Room, error)
}

type RoomService interface {
	GetAllRoom() ([]*RoomRes, error)
}
