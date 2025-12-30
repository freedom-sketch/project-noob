package database

import "github.com/freedom-sketch/sub2go/internal/database/models"

func AutoMigrate() error {
	return DB.AutoMigrate(
		&models.User{},
		&models.Admin{},
		&models.Subscription{},
		&models.Country{},
		&models.Server{},
	)
}
