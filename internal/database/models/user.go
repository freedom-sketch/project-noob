package models

import "time"

type User struct {
	UUID       string    `gorm:"primaryKey;column:uuid;type:varchar(36);unique;not null;index"`
	TelegramID *int64    `gorm:"column:tg_id;uniqueIndex"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`

	Subscriptions []Subscription `gorm:"foreignKey:UserUUID;references:UUID;constraint:OnDelete:CASCADE"`
	Admin         *Admin         `gorm:"foreignKey:UUID;references:UserUUID"`
}

func (User) TableName() string {
	return "users"
}
