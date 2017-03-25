package repository

import "github.com/dpb587/metalink"

type File struct {
	Repository Repository        `json:"repository"`
	Metalink   metalink.Metalink `json:"metalink"`
	File       metalink.File     `json:"file"`
}
