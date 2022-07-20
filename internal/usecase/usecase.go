package usecase

import (
	"postservice/internal/config"
	"postservice/internal/repository"
	"postservice/pkg/logger"
)

type IUseCase interface {
	Api() IApi
	Grud() IGrud
	Loader() ILoader
}

type useCase struct {
	api    IApi
	grud   IGrud
	loader ILoader
}

func (u *useCase) Api() IApi {
	return u.api
}

func (u *useCase) Grud() IGrud {
	return u.grud
}
func (u *useCase) Loader() ILoader {
	return u.loader
}

func New(cfg config.IConfig, lg logger.ILogger, repo repository.IRepository) IUseCase {
	return &useCase{
		api:    newApi(cfg, lg),
		grud:   newGrud(cfg, lg, repo),
		loader: newLoader(cfg, lg, repo),
	}
}
