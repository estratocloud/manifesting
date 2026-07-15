package internal

import (
	"testing"
)

func overrideFS(t *testing.T) {
	old := fs
	t.Cleanup(func() {
		fs = old
	})
}
