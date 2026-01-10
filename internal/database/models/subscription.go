package models

import "time"

type Subscription struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	UserUUID     string    `gorm:"column:user_uuid;type:varchar(36);not null;index;constraint:OnDelete:CASCADE"`
	Email        string    `gorm:"type:varchar(128);index"`
	UsedTraffic  int64     `gorm:"default:0"`
	TotalTraffic int64     `gorm:"default:0"`
	StartDate    time.Time `gorm:"not null"`
	EndDate      time.Time `gorm:"not null;index"`

	User User `gorm:"foreignKey:UserUUID;references:UUID"`
}

func (Subscription) TableName() string {
	return "subscriptions"
}
