package repository

import (
	"base_grpc/domain/model"
	"context"
)

type Car interface {
	GetCar(ctx context.Context) *model.Car
}
