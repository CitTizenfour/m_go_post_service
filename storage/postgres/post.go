package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/matrix/microservice/go_post_service/storage/repo"
)

type postRepo[T any] struct {
	DB *sqlx.DB
}

func NewPostRepo[T any](db *sqlx.DB) repo.PostI[T] {
	return &postRepo[T]{DB: db}
}

func (p *postRepo[T]) Delete(id string) error
