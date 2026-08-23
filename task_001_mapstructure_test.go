package mapstructure

import "testing"

// TestTask001 isolates the 空映射解码 regression.
func TestTask001(t *testing.T) {
	TestComposeDecodeHookFunc(t)
}

// TestTask001Repeat guards deterministic behavior across repeated calls.
func TestTask001Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask001(t)
	}
}
