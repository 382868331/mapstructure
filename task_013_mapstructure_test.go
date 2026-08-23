package mapstructure

import "testing"

// TestTask013 isolates the Windows路径字段 regression.
func TestTask013(t *testing.T) {
	TestWeaklyTypedHook(t)
}

// TestTask013Repeat guards deterministic behavior across repeated calls.
func TestTask013Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask013(t)
	}
}
