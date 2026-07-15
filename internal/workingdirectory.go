package internal

import (
	"path/filepath"
)

type WorkingDirectoryInterface interface {
	GetPath() string
	NewPath(path string) PathInterface
}

type WorkingDirectory struct {
	path string
}

func NewWorkingDirectory(path string) (WorkingDirectoryInterface, error) {

	if !filepath.IsAbs(path) {
		cwd, err := fs.Getwd()
		if err != nil {
			return nil, err
		}
		path = filepath.Join(cwd, path)
	}

	_, err := fs.Stat(path)
	if err != nil {
		return nil, err
	}

	return &WorkingDirectory{
		path: path,
	}, nil
}

func (wd *WorkingDirectory) GetPath() string {
	return wd.path
}

func (wd *WorkingDirectory) NewPath(path string) PathInterface {
	return &Path{
		wd:   wd,
		path: path,
	}
}
