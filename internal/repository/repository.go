package repository

import (
	"database/sql"
	postRepo "postservice/internal/repository/post"
	postRepoPostgres "postservice/internal/repository/post/postgres"
	"postservice/pkg/logger"
)

type IRepository interface {
	PostRepository() postRepo.IPostRepository
}

type repository struct {
	postRepository postRepo.IPostRepository
}

func (r *repository) PostRepository() postRepo.IPostRepository {
	return r.postRepository
}

func NewPostgresRepository(lg logger.ILogger, adapter *sql.DB) IRepository {
	return &repository{
		postRepository: postRepoPostgres.NewPostRepository(lg, adapter),
	}
}
