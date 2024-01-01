package repo

import cpb "github.com/matrix/microservice/go_post_service/genproto/post_service"

type (
	PostI[T any] interface {
		CreatePost(req *cpb.CreatePostReq) (response *cpb.PostResponse, err error)
	}
)