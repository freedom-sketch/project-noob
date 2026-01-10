package queries

import (
	"context"
	"testing"
	"time"

	"github.com/freedom-sketch/sub2go/internal/database"
	"github.com/freedom-sketch/sub2go/internal/database/models"
	"github.com/google/uuid"
)

func TestCreateAndGetSubscription(t *testing.T) {
	db, err := database.ConnectInMemory()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}

	ctx := context.Background()

	userUUID := uuid.New().String()

	user := models.User{
		UUID: userUUID,
	}
	if err := db.WithContext(ctx).Create(&user).Error; err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	now := time.Now().Truncate(time.Second)
	endDate := now.Add(30 * 24 * time.Hour)

	sub := &models.Subscription{
		UserUUID:  user.UUID,
		StartDate: now,
		EndDate:   endDate,
		Email:     "test@example.com",
	}

	if err := CreateSubscription(ctx, db, sub); err != nil {
		t.Fatalf("Error creating subscription: %v", err)
	}

	got, err := GetSubscriptionByUserUUID(ctx, db, user.UUID)
	if err != nil {
		t.Fatalf("Failed to get subscription: %v", err)
	}

	t.Logf("Subscription created and received: %+v", got)

	if !got.StartDate.Equal(now) {
		t.Errorf("StartDate does not match: expected %v, got %v", now, got.StartDate)
	}

	if !got.EndDate.Equal(endDate) {
		t.Errorf("EndDate does not match: expected %v, got %v", endDate, got.EndDate)
	}

	if got.UserUUID != user.UUID {
		t.Errorf("UserUUID does not match: expected %s, got %s", user.UUID, got.UserUUID)
	}
}

func TestExtendSubscription(t *testing.T) {
	db, err := database.ConnectInMemory()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}

	ctx := context.Background()

	userUUID := uuid.New().String()

	user := models.User{
		UUID: userUUID,
	}
	if err := db.WithContext(ctx).Create(&user).Error; err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	now := time.Now().Truncate(time.Second)
	initialEnd := now.Add(30 * 24 * time.Hour)

	sub := &models.Subscription{
		UserUUID:  user.UUID,
		StartDate: now,
		EndDate:   initialEnd,
	}

	if err := CreateSubscription(ctx, db, sub); err != nil {
		t.Fatalf("Error creating subscription: %v", err)
	}

	t.Logf("Initial end date: %v", initialEnd)

	err = ExtendSubscription(ctx, db, user.UUID, 7)
	if err != nil {
		t.Fatalf("Error extending subscription: %v", err)
	}

	got, err := GetSubscriptionByUserUUID(ctx, db, user.UUID)
	if err != nil {
		t.Fatalf("Failed to get subscription after extension: %v", err)
	}

	t.Logf("After extension, subscription valid until: %v", got.EndDate)
}
