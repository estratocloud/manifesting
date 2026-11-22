package cli

import (
	"fmt"
	"strings"
)

type Variables map[string]any

func (v *Variables) String() string {
	return fmt.Sprintf("%v", *v)
}

func (v *Variables) Set(value string) error {
	key, val, ok := strings.Cut(value, "=")
	if !ok {
		return fmt.Errorf("invalid variable %q: expected key=value", value)
	}

	if *v == nil {
		*v = make(map[string]any)
	}

	(*v)[key] = val
	return nil
}

func (v *Variables) GetAll() map[string]any {
	return *v
}
