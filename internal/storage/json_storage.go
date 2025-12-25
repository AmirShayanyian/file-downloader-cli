package storage

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/AmirShayanyian/file-downloader-cli/internal/queue"
)

type JSONStore struct {
	path string
}

func NewJSONStore(path string) *JSONStore {
	return &JSONStore{path: path}
}

func (s *JSONStore) List() ([]queue.Queue, error) {
	if _, err := os.Stat(s.path); os.IsNotExist(err) {
		return []queue.Queue{}, nil
	}

	data, err := os.ReadFile(s.path)
	if err != nil {
		return nil, err
	}

	var queues []queue.Queue
	if err := json.Unmarshal(data, &queues); err != nil {
		return nil, err
	}

	return queues, nil
}

func (s *JSONStore) Save(q queue.Queue) error {
	queues, _ := s.List()

	queues = append(queues, q)

	data, err := json.MarshalIndent(queues, "", "  ")
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(s.path), 0755); err != nil {
		return err
	}

	return os.WriteFile(s.path, data, 0644)
}
