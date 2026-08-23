package mapstructure

import "testing"

// TestTask012 isolates the ANSI字段名 regression.
func TestTask012(t *testing.T) {
	TestStringToIPNetHookFunc(t)
}

// TestTask012Repeat guards deterministic behavior across repeated calls.
func TestTask012Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask012(t)
	}
}
