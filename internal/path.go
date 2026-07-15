package internal

import (
	"fmt"
	"os"
	"path/filepath"
)

type PathInterface interface {
	GetFullyQualifiedPath() string
	Exists() (bool, error)
	ExistsOrError(message string) error
	Open() (*os.File, error)
	ReadFile() ([]byte, error)
	WriteFile(data []byte) error
}

type Path struct {
	wd   WorkingDirectoryInterface
	path string
}

func (p *Path) GetFullyQualifiedPath() string {

	if filepath.IsAbs(p.path) {
		return p.path
	}

	return filepath.Join(p.wd.GetPath(), p.path)
}

func (p *Path) Exists() (bool, error) {
	_, err := fs.Stat(p.GetFullyQualifiedPath())

	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func (p *Path) ExistsOrError(message string) error {
	exists, err := p.Exists()
	if err != nil {
		return err
	}

	if !exists {
		if message == "" {
			message = "file '%s' does not exist"
		}
		return fmt.Errorf(message, p.GetFullyQualifiedPath())
	}

	return nil
}

func (p *Path) Open() (*os.File, error) {
	return fs.Open(p.GetFullyQualifiedPath())
}

func (p *Path) ReadFile() ([]byte, error) {
	return fs.ReadFile(p.GetFullyQualifiedPath())
}

func (p *Path) WriteFile(data []byte) error {
	return fs.WriteFile(p.GetFullyQualifiedPath(), data, 0644)
}
