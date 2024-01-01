package service

import "github.com/matrix/microservice/go_post_service/config"

type ServiceManager interface{}

type grpClients struct{}

func NewGrpClients(cfg config.Config) (ServiceManager, error)