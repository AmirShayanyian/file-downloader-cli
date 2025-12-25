package storage

type JSONStore struct {
	path string
}

func NewJSONStore(path string) *JSONStore {
	return &JSONStore{path: path}
}
