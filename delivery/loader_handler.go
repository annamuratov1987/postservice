package delivery

import (
	"context"
	"postservice/internal/usecase"

	deliveryGrpc "postservice/delivery/grpc"

	"github.com/golang/protobuf/ptypes/empty"
)

func NewLoaderServiceServer(loaderUseCase usecase.IUseCase) deliveryGrpc.LoaderServiceServer {
	return &loaderServiceServer{
		loaderUseCase: loaderUseCase,
	}
}

type loaderServiceServer struct {
	loaderUseCase usecase.IUseCase
}

func (s loaderServiceServer) Start(ctx context.Context, _ *empty.Empty) (res *deliveryGrpc.LoaderResponse, err error) {
	status, err := s.loaderUseCase.Loader().Start(ctx)
	if err != nil {
		return
	}

	return &deliveryGrpc.LoaderResponse{
		Status: status,
	}, nil
}

func (s loaderServiceServer) Check(ctx context.Context, _ *empty.Empty) (res *deliveryGrpc.LoaderResponse, err error) {
	status, err := s.loaderUseCase.Loader().Check(ctx)
	if err != nil {
		return
	}

	return &deliveryGrpc.LoaderResponse{
		Status: status,
	}, nil
}
