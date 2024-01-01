package repo

type (
	PostI[T any] interface {
		Delete(id string) error
	}
)