package mapstructure

import "testing"

// TestTask019 isolates the 空结构 regression.
func TestTask019(t *testing.T) {
	TestNestedTypeWithDefaults(t)
}

// TestTask019Repeat guards deterministic behavior across repeated calls.
func TestTask019Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask019(t)
	}
}
