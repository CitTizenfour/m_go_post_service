package service

import (
	"context"

	"github.com/matrix/microservice/go_post_service/config"
	cpb "github.com/matrix/microservice/go_post_service/genproto/post_service"
	"github.com/matrix/microservice/go_post_service/pkg/logger"
	"github.com/matrix/microservice/go_post_service/storage"
	"go.uber.org/zap"
)

type postService struct {
	*config.Config
	storage.StorageI
	ServiceManager
	logger.Logger
	cpb.UnimplementedPostServiceServer
}

func NewPostService(cfg config.Config, storage storage.StorageI, service ServiceManager, log logger.Logger) *postService {
	return &postService{
		Config:         &cfg,
		StorageI:       storage,
		ServiceManager: service,
		Logger:         log,
	}
}

func (p *postService) CreatePost(ctx context.Context, request *cpb.CreatePostReq) (resp *cpb.PostResponse, err error) {
	p.Logger.Info(`Create Post req:`, logger.Any("req", request))
	resp, err = p.StorageI.Post().CreatePost(request)
	if err != nil {
		p.Logger.Error(`Error while creating Post`, zap.Error(err))
		return nil, err
	}
	return
}	