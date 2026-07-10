package internal

import "os"

var fs = struct {
	Getwd     func() (string, error)
	Stat      func(string) (os.FileInfo, error)
	Open      func(string) (*os.File, error)
	ReadFile  func(string) ([]byte, error)
	WriteFile func(string, []byte, os.FileMode) error
}{
	Getwd:     os.Getwd,
	Stat:      os.Stat,
	Open:      os.Open,
	ReadFile:  os.ReadFile,
	WriteFile: os.WriteFile,
}
