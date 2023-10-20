package clients

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/circuitbreaker"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	pod "github.com/zbnzl/paas-pod/api/pod/v1"
	"github.com/zbnzl/paas-podApi/internal/conf"
)

func NewPodGrpcClient(data *conf.Data, logger log.Logger) pod.PodClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(data.GetClient().GetPod()),
		grpc.WithMiddleware(
			logging.Client(logger),
			circuitbreaker.Client(),
		),
	)
	if err != nil {
		panic("conn pod failed")
	}

	return pod.NewPodClient(conn)
}

func NewPodHttpClient(client *conf.Data_Client, logger log.Logger) pod.PodHTTPClient {
	conn, err := http.NewClient(
		context.Background(),
		http.WithMiddleware(
			logging.Client(logger),
			circuitbreaker.Client(),
		),
		http.WithEndpoint(client.GetPod()),
	)
	if err != nil {
		panic("conn pod failed")
	}
	return pod.NewPodHTTPClient(conn)
}
