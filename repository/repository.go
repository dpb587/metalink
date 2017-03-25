package repository

type Repository struct {
	URI     string `json:"uri"`
	Version string `json:"version"`
	Path    string `json:"path"`
}
