package delivery

import (
	"context"
	"database/sql"
	postRepo "postservice/internal/repository/post"
	"postservice/internal/usecase"

	deliveryGrpc "postservice/delivery/grpc"

	"github.com/golang/protobuf/ptypes/empty"
)

func NewGrudServiceServer(grudUseCase usecase.IUseCase) deliveryGrpc.GrudServiceServer {
	return &grudServiceServer{
		grudUseCase: grudUseCase,
	}
}

type grudServiceServer struct {
	grudUseCase usecase.IUseCase
}

func (g *grudServiceServer) GetAll(ctx context.Context, e *empty.Empty) (*deliveryGrpc.GetAllResponse, error) {
	posts, err := g.grudUseCase.Grud().GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var data []*deliveryGrpc.Post
	for _, post := range posts {
		data = append(data, &deliveryGrpc.Post{
			Id:     post.GetId(),
			UserId: post.GetUserId(),
			Title:  post.GetTitle(),
			Body:   post.GetBody(),
		})
	}

	return &deliveryGrpc.GetAllResponse{
		Status: &deliveryGrpc.Status{
			Name:    OkStatus,
			Message: "",
		},
		Data: data,
	}, nil
}

func (g *grudServiceServer) Get(ctx context.Context, request *deliveryGrpc.GrudRequest) (*deliveryGrpc.GetResponse, error) {
	post, err := g.grudUseCase.Grud().Get(ctx, request.Id)
	if err == sql.ErrNoRows {
		return &deliveryGrpc.GetResponse{
			Status: &deliveryGrpc.Status{
				Name:    ErrorStatus,
				Message: "Not found",
			},
			Data: nil,
		}, nil
	}
	if err != nil {
		return nil, err
	}

	data := &deliveryGrpc.Post{
		Id:     post.GetId(),
		UserId: post.GetUserId(),
		Title:  post.GetTitle(),
		Body:   post.GetBody(),
	}

	return &deliveryGrpc.GetResponse{
		Status: &deliveryGrpc.Status{
			Name:    OkStatus,
			Message: "",
		},
		Data: data,
	}, nil
}

func (g *grudServiceServer) Update(ctx context.Context, request *deliveryGrpc.GrudRequest) (*deliveryGrpc.Status, error) {
	post := postRepo.Post{
		Id:     request.GetId(),
		UserId: request.GetUserId(),
		Title:  request.GetTitle(),
		Body:   request.GetBody(),
	}
	rowsAffected, err := g.grudUseCase.Grud().Update(ctx, &post)
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return &deliveryGrpc.Status{
			Name:    ErrorStatus,
			Message: "Not found",
		}, err
	}

	status := &deliveryGrpc.Status{
		Name:    OkStatus,
		Message: "Update completed successfully",
	}

	return status, nil
}

func (g *grudServiceServer) Delete(ctx context.Context, request *deliveryGrpc.GrudRequest) (*deliveryGrpc.Status, error) {
	rowsAffected, err := g.grudUseCase.Grud().Delete(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return &deliveryGrpc.Status{
			Name:    ErrorStatus,
			Message: "Not found",
		}, err
	}

	status := &deliveryGrpc.Status{
		Name:    OkStatus,
		Message: "Delete completed successfully",
	}

	return status, nil
}
