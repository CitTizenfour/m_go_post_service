package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/matrix/microservice/go_post_service/storage/postgres"
	"github.com/matrix/microservice/go_post_service/storage/repo"
)

type StorageI interface {
	Post() repo.PostI[any]
}

type storagePG struct {
	db       *sqlx.DB
	postRepo repo.PostI[any]
}

func NewStoragePG[T any](db *sqlx.DB) StorageI {
	return &storagePG{
		db:       db,
		postRepo: postgres.NewPostRepo[T](db),
	}
}

func (s storagePG) Post() repo.PostI[any] {
	return s.postRepo
}
