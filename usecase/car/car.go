package car

import (
	"base_grpc/domain/model"
	"base_grpc/domain/repository"
	"context"
)

type carUsecase struct {
}

func NewCar() repository.Car {
	return &carUsecase{}
}

func (c carUsecase) GetCar(ctx context.Context) *model.Car {
	return &model.Car{
		Name: "BMW",
		Type: "BMW",
	}
}