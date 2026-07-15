package mocks

import (
	"github.com/estratocloud/manifesting/internal"
)

type WorkingDirectory struct {
	GetPathFunc func() string
	NewPathFunc func(path string) internal.PathInterface
}

func (wd *WorkingDirectory) GetPath() string {
	if wd.GetPathFunc != nil {
		return wd.GetPathFunc()
	}
	return ""
}

func (wd *WorkingDirectory) NewPath(path string) internal.PathInterface {
	if wd.NewPathFunc != nil {
		return wd.NewPathFunc(path)
	}
	return nil
}
