package mapstructure

import "testing"

// TestTask017 isolates the 标签匹配 regression.
func TestTask017(t *testing.T) {
	TestNestedTypePointerWithDefaults(t)
}

// TestTask017Repeat guards deterministic behavior across repeated calls.
func TestTask017Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask017(t)
	}
}
