package deprecations

import (
	"errors"
	"fmt"
	"os"
)

type Checker struct {
	fail   bool
	raised map[string]struct{}
}

func New(fail bool) *Checker {
	return &Checker{
		fail:   fail,
		raised: make(map[string]struct{}),
	}
}

func (d *Checker) WillFail() bool {
	return d.fail
}

func (d *Checker) IsError(message string) error {

	if d.fail {
		return errors.New(message)
	}

	if _, ok := d.raised[message]; ok {
		return nil
	}
	d.raised[message] = struct{}{}

	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "⚠️ ", "\033[33mDEPRECATED:", message, "\033[0m")

	return nil
}
