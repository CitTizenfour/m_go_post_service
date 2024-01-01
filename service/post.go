package service

import (
	"github.com/matrix/microservice/go_post_service/config"
	"github.com/matrix/microservice/go_post_service/pkg/logger"
	"github.com/matrix/microservice/go_post_service/storage"
)

type postService struct {
	*config.Config
	storage.StorageI
	ServiceManager
	logger.Logger
}

func NewPostService(cfg config.Config, storage storage.StorageI, service ServiceManager, log logger.Logger) *postService {
	return &postService{
		Config:         &cfg,
		StorageI:       storage,
		ServiceManager: service,
		Logger:         log,
	}
}
