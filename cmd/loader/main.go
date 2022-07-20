package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"postservice/delivery"
	deliveryGrpc "postservice/delivery/grpc"
	"postservice/internal/config"
	"postservice/internal/repository"
	"postservice/internal/usecase"
	"postservice/pkg/logger"
)

var (
	configPath string
	logPath    string
)

func init() {
	flag.StringVar(&configPath, `config-path`, `./configs`, `Configuration file path`)
	flag.StringVar(&logPath, `log-path`, ``, `Configuration log path`)
	flag.Parse()
}

func main() {
	// Initial logger
	lg := newLogger()
	defer func() { _ = lg.Sync() }()

	lg.Info(`Start`)
	defer lg.Info(`Stop`)

	// Initial config
	cfg := newConfig(lg)

	stop := make(chan os.Signal, 1)
	signal.Notify(
		stop,
		os.Kill,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	// Start listen port
	address := cfg.Loader().GetHost() + ":" + strconv.Itoa(cfg.Loader().GetPort())
	listen, err := net.Listen("tcp", address)
	if err != nil {
		lg.Fatal(
			`TcpListener.Start.Error`,
			zap.String(`address`, address),
			zap.Error(err),
		)
	}
	defer listen.Close()

	srv := grpc.NewServer()

	postgresRepository := newPostgresRepository(cfg.Database().Postgres(), lg)

	deliveryGrpc.RegisterLoaderServiceServer(
		srv,
		delivery.NewLoaderServiceServer(
			usecase.New(
				cfg,
				lg,
				postgresRepository,
			),
		),
	)

	// Start server
	go func() {
		lg.Info(`Main.Listening`, zap.String(`address`, address))
		if err := srv.Serve(listen); err != nil {
			log.Fatal(err)
		}
	}()

	// Waiting stop signal
	stopSignal := <-stop
	lg.Info(
		`Main.Signal`,
		zap.String(`signal`, stopSignal.String()),
	)

	// Safe shutdown server
	shutdownServer(srv, lg)
}

// Initial config
func newConfig(lg logger.ILogger) config.IConfig {
	config, err := config.New(configPath)
	if err != nil {
		lg.Fatal(`Config.Error`, zap.Error(err))
	}

	return config
}

// Initial logger
func newLogger() logger.ILogger {
	lg, err := logger.New(logPath)
	if err != nil {
		log.Fatalf(`Start logger: %s\n`, err)
	}

	return lg
}

func newPostgresRepository(cfg config.IPostgres, lg logger.ILogger) repository.IRepository {
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.GetUserName(), cfg.GetPassword(), cfg.GetHost(), cfg.GetPort(), cfg.GetDbName())

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		lg.Fatal(
			`Postgres.Connect.Error`,
			zap.Error(err),
		)
	}

	return repository.NewPostgresRepository(lg, db)
}

// Safe shutdown server
func shutdownServer(srv *grpc.Server, lg logger.ILogger) {
	lg.Info("Main.Shutdowning")
	srv.GracefulStop()
}
