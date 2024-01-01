package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/jmoiron/sqlx"
	"github.com/matrix/microservice/go_post_service/config"
	"github.com/matrix/microservice/go_post_service/pkg/logger"
	"github.com/matrix/microservice/go_post_service/service"
	"github.com/matrix/microservice/go_post_service/storage"
	"google.golang.org/grpc"
)

func main() {

	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "post_service")
	lis, err := net.Listen("tcp", cfg.ServiceHost+":"+strconv.Itoa(cfg.ServicePort))
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	conStr := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=%s", cfg.PostgresHost,
		cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB, "disable")

	db, err := sqlx.Connect("pgx", conStr)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(10)
	storageI := storage.NewStoragePG[any](db)

	grpClients, err := service.NewGrpClients(cfg)
	if err != nil {
		panic(err)
	}

	postService := service.NewPostService(cfg, storageI, grpClients, log)

	s := grpc.NewServer()
	// postService generation GRPC

	go func() {
		log.Info(`Post Service running...`)
		if err := s.Serve(lis); err != nil {
			log.Fatal("Error while listening: %v", logger.Error(err))
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	<-quit
	log.Info("Shutdown Server ...")

	s.GracefulStop()

}

// type StorageI interface {
// 	Post() PostI[any]
// }

// type storagePG struct {
// 	db       *sqlx.DB
// 	postRepo PostI[any]
// }

// func NewStoragePG[T any](db *sqlx.DB) StorageI {
// 	return &storagePG{
// 		db:       db,
// 		postRepo: NewPostRepo[T](db),
// 	}
// }

// func (s storagePG) Post() PostI[any] {
// 	return s.postRepo
// }

// type postRepo[T any] struct {
// 	DB *sqlx.DB
// }

// func NewPostRepo[T any](db *sqlx.DB) PostI[T] {
// 	return &postRepo[T]{DB: db}
// }

// func (p *postRepo[T]) Delete(id string) error

// type (
// 	PostI[T any] interface {
// 		Delete(id string) error
// 	}
// )
