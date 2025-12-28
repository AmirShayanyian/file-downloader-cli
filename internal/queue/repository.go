package queue

type Repository interface {
	Save(queue Queue) error
	Update(queue Queue) error
	Find(name string) (*Queue, error)
	List() ([]Queue, error)
}
