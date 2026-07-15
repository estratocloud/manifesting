package internal

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func assertMerge(t *testing.T, baseInput string, overrideInput string, expectedOutput string) {
	base := []byte(strings.TrimSpace(baseInput))
	override := []byte(strings.TrimSpace(overrideInput))
	expected := strings.TrimSpace(expectedOutput)

	result, err := MergeYAML(base, override)
	got := strings.TrimSpace(string(result))

	require.NoError(t, err)
	assert.Equal(t, expected, got)
}

// MergeYAML Ensure we can merge two documents
func Test_MergeYAML1(t *testing.T) {
	assertMerge(t, `
name: base
base: stuff
`, `
name: override
additional: from-override
`, `
additional: from-override
base: stuff
name: override
`)
}

// MergeYAML Ensure we can merge nested objects/maps
func Test_MergeYAML2(t *testing.T) {
	assertMerge(t, `
map:
    layer:
        base: base
        override: base
spec:
    - name: one
      value: base
`, `
map:
    layer:
        extra: override
        override: override
spec:
    - name: two
      value: override
`, `
map:
    layer:
        base: base
        extra: override
        override: override
spec:
    - name: one
      value: base
    - name: two
      value: override
`)
}
