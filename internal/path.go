package internal

import (
	"fmt"
	"os"
	"path/filepath"
)

type WorkingDirectory struct {
	path string
}

func NewWorkingDirectory(path string) (*WorkingDirectory, error) {

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

type Path struct {
	WorkingDirectory *WorkingDirectory
	Path             string
}

func (p *Path) GetFullyQualifiedPath() string {

	if filepath.IsAbs(p.Path) {
		return p.Path
	}

	return filepath.Join(p.WorkingDirectory.path, p.Path)
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
