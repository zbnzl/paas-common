package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pod "github.com/zbnzl/paas-pod/api/pod/v1"
	pb "github.com/zbnzl/paas-pod/api/podapi/v1"
)

type PodApiService struct {
	pb.UnimplementedPodApiServer
	podClient pod.PodClient
	log       *log.Helper
}

func NewPodApiService(podClient pod.PodClient, logger log.Logger) *PodApiService {
	return &PodApiService{podClient: podClient, log: log.NewHelper(logger)}
}

// FindPodById podApi.FindPodById 通过API向外暴露为/podApi/findPodById，接收http请求
// 即：/podApi/FindPodById 请求会调用go.micro.api.podApi 服务的podApi.FindPodById 方法
func (s *PodApiService) FindPodById(ctx context.Context, req *pod.PodId) (*pod.PodInfo, error) {
	return s.podClient.FindPodByID(ctx, req)
}

// AddPod podApi.AddPod 通过API向外暴露为/podApi/addPod，接收http请求
// 即：/podApi/AddPod 请求会调用go.micro.api.podApi 服务的podApi.AddPod 方法
func (s *PodApiService) AddPod(ctx context.Context, req *pod.PodInfo) (*pod.Response, error) {
	return &pod.Response{}, nil
}

// DeletePodById podApi.DeletePodById 通过API向外暴露为/podApi/deletePodById，接收http请求
// 即：/podApi/DeletePodById 请求会调用go.micro.api.podApi 服务的 podApi.DeletePodById 方法
func (s *PodApiService) DeletePodById(ctx context.Context, req *pod.PodId) (*pod.Response, error) {
	return &pod.Response{}, nil
}

// UpdatePod podApi.UpdatePod 通过API向外暴露为/podApi/updatePod，接收http请求
// 即：/podApi/UpdatePod 请求会调用go.micro.api.podApi 服务的podApi.UpdatePod 方法
func (s *PodApiService) UpdatePod(ctx context.Context, req *pod.PodInfo) (*pod.Response, error) {
	return &pod.Response{}, nil
}

// Call 默认的方法podApi.Call 通过API向外暴露为/podApi/call，接收http请求
// 即：/podApi/call或/podApi/ 请求会调用go.micro.api.podApi 服务的podApi.Call 方法
func (s *PodApiService) Call(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{}, nil
}
