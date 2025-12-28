package queue

import "github.com/AmirShayanyian/file-downloader-cli/internal/download"

type Queue struct {
	Name  string           `json:"name"`
	Tasks []*download.Task `json:"tasks"`
}
