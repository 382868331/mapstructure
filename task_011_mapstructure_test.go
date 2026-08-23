package mapstructure

import "testing"

// TestTask011 isolates the 并发元数据 regression.
func TestTask011(t *testing.T) {
	TestStringToIPHookFunc(t)
}

// TestTask011Repeat guards deterministic behavior across repeated calls.
func TestTask011Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask011(t)
	}
}
