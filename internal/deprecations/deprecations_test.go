package deprecations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// New Ensure new works with false
func Test_New1(t *testing.T) {
	got := New(false)
	assert.Equal(t, &Checker{
		fail:   false,
		raised: make(map[string]struct{}),
	}, got)
}

// New Ensure new works with true
func Test_New2(t *testing.T) {
	got := New(true)
	assert.Equal(t, &Checker{
		fail:   true,
		raised: make(map[string]struct{}),
	}, got)
}

// WillFail Ensure WillFail returns false
func Test_WillFail1(t *testing.T) {
	got := New(false).WillFail()
	assert.Equal(t, false, got)
}

// WillFail Ensure WillFail returns true
func Test_WillFail2(t *testing.T) {
	got := New(true).WillFail()
	assert.Equal(t, true, got)
}

// IsError Doesn't return an error in non-failure more
func Test_IsError1(t *testing.T) {
	got := New(false).IsError("example")
	assert.NoError(t, got)
}

// IsError Does return an error in failure more
func Test_IsError2(t *testing.T) {
	got := New(true).IsError("example")
	assert.EqualError(t, got, "example")
}

// IsError Always returns an error, even if it's the same message
func Test_IsError3(t *testing.T) {
	d := New(true)
	got := d.IsError("example")
	assert.EqualError(t, got, "example")
	got = d.IsError("example")
	assert.EqualError(t, got, "example")
}
