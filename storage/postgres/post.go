package postgres

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	cpb "github.com/matrix/microservice/go_post_service/genproto/post_service"
	"github.com/matrix/microservice/go_post_service/storage/repo"
)

var (
	updatedAt, createdAt time.Time
)

type postRepo[T any] struct {
	DB *sqlx.DB
}

func NewPostRepo[T any](db *sqlx.DB) repo.PostI[T] {
	return &postRepo[T]{DB: db}
}

func (p *postRepo[T]) CreatePost(req *cpb.CreatePostReq) (response *cpb.PostResponse, err error) {

	query := `
		INSERT INTO post (
			text,
			content,
			publication_date
		) VALUES ($1, $2, $3) 
	`
	_, err = p.DB.Exec(query, req.Text, req.Content, req.PublicationDate)
	fmt.Println(`CHECK`, err)
	if err != nil {
		return nil, err
	}

	response = &cpb.PostResponse{
		Id:              req.Id,
		Text:            req.Text,
		Content:         req.Content,
		PublicationDate: req.PublicationDate,
		CreatedAt:       createdAt.Format(time.RFC3339),
		UpdatedAt:       updatedAt.Format(time.RFC3339),
	}

	return
}
