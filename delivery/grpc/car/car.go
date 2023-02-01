package car

import (
	"base_grpc/delivery/proto"
	"base_grpc/domain/repository"
	"context"
)

type carService struct {
	usecase repository.Car
}

func (c carService) GetCar(ctx context.Context, request *proto.CarRequest) (*proto.Car, error) {
	res := c.usecase.GetCar(ctx)
	return &proto.Car{
		Car:  res.Name,
		Type: res.Type,
	}, nil
}

func NewCarSer(usecase repository.Car) proto.CarObjectServer {
	return &carService{
		usecase: usecase,
	}
}
