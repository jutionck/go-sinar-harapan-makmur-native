package repository

type BaseRepository[T any] interface {
	Create(newData T) error
	List() ([]T, error)
	Get(id string) (T, error)
	Update(newData T) error
	Delete(id string) error
}
