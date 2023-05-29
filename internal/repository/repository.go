package repository

import (
	"context"
	"time"

	"github.com/AbdulwahabNour/booking/internal/models"
)


type DatabaseRepo interface {
	InsertReservation(ctx context.Context, res *models.Reservation)(int ,error)
	InsertRoomRestrictions(ctx context.Context, res *models.RoomRestriction)error
	CheckAvailabilityByDateAndRoom(ctx context.Context,roomId int, start, end time.Time) (bool, error)
	SearchAvailabilityForRooms(ctx context.Context, pageSize int, offset int, start, end time.Time) ([]models.Room, error)
	GetRoomById(ctx context.Context, id int) (*models.Room, error)
	CountRooms(ctx context.Context) (int, error)
	GetReservationById(ctx context.Context, id int) (*models.Reservation, error)
	GetRoomsByOffset(ctx context.Context, pageSize int, offset int) ([]models.Room, error)
 
} 