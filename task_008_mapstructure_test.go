package mapstructure

import "testing"

// TestTask008 isolates the 目标容量 regression.
func TestTask008(t *testing.T) {
	TestStringToSliceHookFunc(t)
}

// TestTask008Repeat guards deterministic behavior across repeated calls.
func TestTask008Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask008(t)
	}
}
