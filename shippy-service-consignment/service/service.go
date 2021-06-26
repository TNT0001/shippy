package service

import (
	"context"
	pb "tungnguyen.shippy/shippy-service-consignment/proto/consignment"
	"tungnguyen.shippy/shippy-service-consignment/repository"
)

type Service interface {
	CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error)
	GetConsignment(context.Context, *pb.GetRequest) (*pb.Response, error)
	mustEmbedUnimplementedShippingServiceServer()
}

type service struct {
	repo repository.Repository
	pb.UnimplementedShippingServiceServer
}

type ServiceConfig struct {
	Repo repository.Repository
}

func NewService(serviceConfig ServiceConfig) *service {
	return &service{
		repo:                               serviceConfig.Repo,
		UnimplementedShippingServiceServer: pb.UnimplementedShippingServiceServer{},
	}
}

func (srv *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	consignment, err := srv.repo.Create(req)
	if err != nil {
		return nil, err
	}

	return &pb.Response{Create: true, Consignment: consignment}, nil
}

func (srv *service) GetConsignment(ctx context.Context, getRequest *pb.GetRequest) (*pb.Response, error) {
	consignments, err := srv.repo.GetConsignment(getRequest)
	if err != nil {
		return nil, err
	}

	return &pb.Response{Create: true, Consignments: consignments}, nil
}
