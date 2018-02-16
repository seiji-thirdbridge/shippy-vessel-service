// shippy-vessel-service/handler.go

package main

import (
	"context"

	"gopkg.in/mgo.v2"

	pb "github.com/seiji-thirdbridge/shippy-vessel-service/proto/vessel"
)

type service struct {
	session *mgo.Session
}

func (s *service) getRepo() repository {
	return &VesselRepository{s.session.Clone()}
}

func (s *service) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	repo := s.getRepo()
	defer repo.Close()

	vessel, err := repo.FindAvailable(req)
	if err != nil {
		return err
	}

	res.Vessel = vessel
	return nil
}

func (s *service) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	repo := s.getRepo()
	defer repo.Close()

	err := repo.Create(req)
	if err != nil {
		return err
	}

	res.Vessel = req
	res.Created = true
	return nil
}
