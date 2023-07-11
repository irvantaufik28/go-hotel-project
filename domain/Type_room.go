package domain

import "time"

type TypeRoom struct {
	ID          int64     `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(40)" json:"name"`
	Price       int64     `gorm:"type:int" json:"price"`
	Description string    `gorm:"type:text" json:"description"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
	Room        []Room    `gorm:"foreignKey:TypeRoomID"`
}

type TypeRoomRepository interface {
	GetAllTypeRoom() ([]*TypeRoom, error)
}

type TypeRoomRes struct {
	Id          int64     `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(40)" json:"name"`
	Price       int64     `gorm:"type:int" json:"price"`
	Description string    `gorm:"text" json:"description"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
	Room        []Room    `gorm:"foreignKey:TypeRoomID"`
}

type TypeRoomService interface {
	GetAllTypeRoom() ([]*TypeRoomRes, error)
}
