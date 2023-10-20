package service

import (
	"context"

	pb "github.com/zbnzl/paas-podApi/api/podApi/v1"
)

type PodApiService struct {
	pb.UnimplementedPodApiServer
}

func NewPodApiService() *PodApiService {
	return &PodApiService{}
}

func (s *PodApiService) FindPodById(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{}, nil
}
func (s *PodApiService) AddPod(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{}, nil
}
func (s *PodApiService) DeletePodById(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{}, nil
}
func (s *PodApiService) UpdatePod(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{}, nil
}
func (s *PodApiService) Call(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{}, nil
}
