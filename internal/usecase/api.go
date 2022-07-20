package usecase

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	deliveryGrpc "postservice/delivery/grpc"
	"postservice/internal/config"
	postRepo "postservice/internal/repository/post"
	"postservice/pkg/logger"
	"strconv"
)

type IApi interface {
	LoaderStart(ctx context.Context) (res *deliveryGrpc.LoaderResponse, err error)
	LoaderCheck(ctx context.Context) (res *deliveryGrpc.LoaderResponse, err error)
	GetAll(ctx context.Context) (res *deliveryGrpc.GetAllResponse, err error)
	Get(ctx context.Context, id int64) (res *deliveryGrpc.GetResponse, err error)
	Update(ctx context.Context, post postRepo.IPost) (res *deliveryGrpc.Status, err error)
	Delete(ctx context.Context, id int64) (res *deliveryGrpc.Status, err error)
}

type api struct {
	cfg config.IConfig
	lg  logger.ILogger
}

func newApi(cfg config.IConfig, lg logger.ILogger) IApi {
	return &api{
		cfg: cfg,
		lg:  lg,
	}
}

func (a *api) getLoaderConnection() (conn *grpc.ClientConn, err error) {
	loaderAddress := a.cfg.Loader().GetHost() + ":" + strconv.Itoa(a.cfg.Loader().GetPort())
	conn, err = grpc.Dial(loaderAddress, grpc.WithInsecure())
	return
}

func (a *api) LoaderStart(ctx context.Context) (res *deliveryGrpc.LoaderResponse, err error) {
	conn, err := a.getLoaderConnection()
	if err != nil {
		a.lg.Error("UseCase.IApi.LoaderStart.getLoaderConnection.Error", zap.Error(err))
		return
	}

	client := deliveryGrpc.NewLoaderServiceClient(conn)

	startRequest := &empty.Empty{}
	res, err = client.Start(ctx, startRequest)
	if err != nil {
		a.lg.Error("UseCase.IApi.LoaderStart.gRpcClient.Start.Error", zap.Error(err))
		return
	}

	return
}

func (a *api) LoaderCheck(ctx context.Context) (res *deliveryGrpc.LoaderResponse, err error) {
	conn, err := a.getLoaderConnection()
	if err != nil {
		a.lg.Error("UseCase.IApi.LoaderCheck.getLoaderConnection.Error", zap.Error(err))
		return
	}

	client := deliveryGrpc.NewLoaderServiceClient(conn)

	startRequest := &empty.Empty{}
	res, err = client.Check(context.Background(), startRequest)
	if err != nil {
		a.lg.Error("UseCase.IApi.LoaderCheck.gRpcClient.Check.Error", zap.Error(err))
		return
	}
	return
}

func (a *api) getGrudConnection() (conn *grpc.ClientConn, err error) {
	grudAddress := a.cfg.Grud().GetHost() + ":" + strconv.Itoa(a.cfg.Grud().GetPort())
	conn, err = grpc.Dial(grudAddress, grpc.WithInsecure())
	return
}

func (a *api) GetAll(ctx context.Context) (res *deliveryGrpc.GetAllResponse, err error) {
	conn, err := a.getGrudConnection()
	if err != nil {
		a.lg.Error("UseCase.IApi.GetAll.getGrudConnection.Error", zap.Error(err))
		return
	}

	client := deliveryGrpc.NewGrudServiceClient(conn)

	startRequest := &empty.Empty{}
	res, err = client.GetAll(ctx, startRequest)
	if err != nil {
		a.lg.Error("UseCase.IApi.GetAll.gRpcClient.GetAll.Error", zap.Error(err))
		return
	}

	return
}

func (a *api) Get(ctx context.Context, id int64) (res *deliveryGrpc.GetResponse, err error) {
	conn, err := a.getGrudConnection()
	if err != nil {
		a.lg.Error("UseCase.IApi.Get.getGrudConnection.Error", zap.Error(err))
		return
	}

	client := deliveryGrpc.NewGrudServiceClient(conn)

	grudRequest := &deliveryGrpc.GrudRequest{Id: id}
	res, err = client.Get(ctx, grudRequest)
	if err != nil {
		a.lg.Error("UseCase.IApi.Get.gRpcClient.Get.Error", zap.Error(err))
		return
	}

	return
}

func (a *api) Update(ctx context.Context, post postRepo.IPost) (res *deliveryGrpc.Status, err error) {
	conn, err := a.getGrudConnection()
	if err != nil {
		a.lg.Error("UseCase.IApi.Update.getGrudConnection.Error", zap.Error(err))
		return
	}

	client := deliveryGrpc.NewGrudServiceClient(conn)

	grudRequest := &deliveryGrpc.GrudRequest{
		Id:     post.GetId(),
		UserId: post.GetUserId(),
		Title:  post.GetTitle(),
		Body:   post.GetBody(),
	}
	res, err = client.Update(ctx, grudRequest)
	if err != nil {
		a.lg.Error("UseCase.IApi.Update.gRpcClient.Update.Error", zap.Error(err))
		return
	}

	return
}

func (a *api) Delete(ctx context.Context, id int64) (res *deliveryGrpc.Status, err error) {
	conn, err := a.getGrudConnection()
	if err != nil {
		a.lg.Error("UseCase.IApi.Delete.getGrudConnection.Error", zap.Error(err))
		return
	}

	client := deliveryGrpc.NewGrudServiceClient(conn)

	grudRequest := &deliveryGrpc.GrudRequest{Id: id}
	res, err = client.Delete(ctx, grudRequest)
	if err != nil {
		a.lg.Error("UseCase.IApi.Delete.gRpcClient.Delete.Error", zap.Error(err))
		return
	}

	return
}
