package usecase

import (
	"context"
	"go.uber.org/zap"
	"postservice/internal/config"
	"postservice/internal/repository"
	"postservice/pkg/logger"
	"sync"
)

const (
	runningStatus = `running`
	doneStatus    = `done`
)

type ILoader interface {
	Start(context.Context) (status string, err error)
	Check(context.Context) (status string, err error)
}

type loader struct {
	status string
	mx     *sync.Mutex

	cfg  config.IConfig
	repo repository.IRepository
	lg   logger.ILogger
}

func newLoader(cfg config.IConfig, lg logger.ILogger, repo repository.IRepository) ILoader {
	return &loader{
		status: doneStatus,
		mx:     &sync.Mutex{},
		cfg:    cfg,
		repo:   repo,
		lg:     lg,
	}
}

func (l *loader) Start(ctx context.Context) (status string, err error) {
	l.mx.Lock()
	defer l.mx.Unlock()

	if l.status == runningStatus {
		return l.status, nil
	}

	l.status = runningStatus
	go l.startLoader()

	return l.status, nil
}

func (l *loader) Check(ctx context.Context) (status string, err error) {
	return l.status, nil
}

func (l *loader) startLoader() {
	urls := make(chan string, 50)
	wg := sync.WaitGroup{}
	urls <- l.cfg.Loader().GetLoadUrl()

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			getUrl := <-urls
			resp, err := l.repo.PostRepository().Load(getUrl)
			if err != nil {
				l.lg.Error("UseCase.Loader.startLoader.Error", zap.Error(err))
				return
			}

			_, err = l.repo.PostRepository().SaveBatch(resp.GetPosts())
			if err != nil {
				l.lg.Error("UseCase.Loader.startLoader.SaveBatch.Error", zap.Error(err))
				return
			}

			urls <- resp.GetNextLink()

			wg.Done()
		}()
	}

	wg.Wait()

	l.status = doneStatus
}
