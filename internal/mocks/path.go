package mocks

import (
	"os"
)

type Path struct {
	Path                      string
	GetFullyQualifiedPathFunc func() string
	ExistsFunc                func() (bool, error)
	ExistsOrErrorFunc         func(string) error
	OpenFunc                  func() (*os.File, error)
	ReadFileFunc              func() ([]byte, error)
	WriteFileFunc             func([]byte) error
}

func (p *Path) GetFullyQualifiedPath() string {
	if p.GetFullyQualifiedPathFunc != nil {
		return p.GetFullyQualifiedPathFunc()
	}
	return p.Path
}

func (p *Path) Exists() (bool, error) {
	if p.ExistsFunc != nil {
		return p.ExistsFunc()
	}
	return false, nil
}

func (p *Path) ExistsOrError(message string) error {
	if p.ExistsOrErrorFunc != nil {
		return p.ExistsOrErrorFunc(message)
	}
	return nil
}

func (p *Path) Open() (*os.File, error) {
	if p.OpenFunc != nil {
		return p.OpenFunc()
	}
	return nil, nil
}

func (p *Path) ReadFile() ([]byte, error) {
	if p.ReadFileFunc != nil {
		return p.ReadFileFunc()
	}
	return nil, nil
}

func (p *Path) WriteFile(data []byte) error {
	if p.WriteFileFunc != nil {
		return p.WriteFileFunc(data)
	}
	return nil
}
