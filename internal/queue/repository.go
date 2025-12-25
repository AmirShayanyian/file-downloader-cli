package queue

type Repository interface {
	Save(queue Queue) error
	List() ([]Queue, error)
}
