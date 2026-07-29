package internal

import (
	"path/filepath"
)

type GeneratedDirectoryInterface interface {
	GetPath() string
	NewPath(path string) PathInterface
}

type GeneratedDirectory struct {
	path string
}

func NewGeneratedDirectory(path string) (GeneratedDirectoryInterface, error) {

	if path == "" {
		path = ".generated"
	}

	if !filepath.IsAbs(path) {
		cwd, err := fs.Getwd()
		if err != nil {
			return nil, err
		}
		path = filepath.Join(cwd, path)
	}

	return &GeneratedDirectory{
		path: path,
	}, nil
}

func (wd *GeneratedDirectory) GetPath() string {
	return wd.path
}

func (wd *GeneratedDirectory) NewPath(path string) PathInterface {
	return &Path{
		wd:   wd,
		path: path,
	}
}
