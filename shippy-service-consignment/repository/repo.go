package repository

import (
	"sync"

	pb "tungnguyen.shippy/shippy-service-consignment/proto/consignment"
)

type Repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetConsignment(*pb.GetRequest) ([]*pb.Consignment, error)
}

type repo struct {
	mu           sync.RWMutex
	consignments []*pb.Consignment
}

func NewRepo() Repository {
	return &repo{mu: sync.RWMutex{}, consignments: []*pb.Consignment{}}
}

func (repo *repo) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	update := append(repo.consignments, consignment)
	repo.consignments = update
	return consignment, nil
}

func (repo *repo) GetConsignment(*pb.GetRequest) ([]*pb.Consignment, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	return repo.consignments, nil
}
