package usecase

import (
	"context"
	"go.uber.org/zap"
	"postservice/internal/config"
	"postservice/internal/repository"
	postRepo "postservice/internal/repository/post"
	"postservice/pkg/logger"
)

type IGrud interface {
	GetAll(context.Context) (posts []postRepo.Post, err error)
	Get(context.Context, int64) (post postRepo.Post, err error)
	Update(context.Context, postRepo.IPost) (rowsAffected int64, err error)
	Delete(context.Context, int64) (rowsAffected int64, err error)
}

type grud struct {
	cfg  config.IConfig
	repo repository.IRepository
	lg   logger.ILogger
}

func newGrud(cfg config.IConfig, lg logger.ILogger, repo repository.IRepository) IGrud {
	return &grud{
		cfg:  cfg,
		repo: repo,
		lg:   lg,
	}
}

func (g *grud) GetAll(ctx context.Context) (posts []postRepo.Post, err error) {
	posts, err = g.repo.PostRepository().GetAll()
	if err != nil {
		g.lg.Error("UseCase.Grud.GetAll.Error", zap.Error(err))
		return
	}
	return
}

func (g *grud) Get(ctx context.Context, id int64) (post postRepo.Post, err error) {
	post, err = g.repo.PostRepository().Get(id)
	if err != nil {
		g.lg.Error("UseCase.Grud.Get.Error", zap.Error(err))
		return
	}
	return
}

func (g *grud) Update(ctx context.Context, post postRepo.IPost) (rowsAffected int64, err error) {
	rowsAffected, err = g.repo.PostRepository().Update(post)
	if err != nil {
		g.lg.Error("UseCase.Grud.Update.Error", zap.Error(err))
		return
	}
	return
}

func (g *grud) Delete(ctx context.Context, id int64) (rowsAffected int64, err error) {
	rowsAffected, err = g.repo.PostRepository().Delete(id)
	if err != nil {
		g.lg.Error("UseCase.Grud.Delete.Error", zap.Error(err))
		return
	}
	return
}
