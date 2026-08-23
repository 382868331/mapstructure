package mapstructure

import "testing"

// TestTask003 isolates the 弱类型数字 regression.
func TestTask003(t *testing.T) {
	TestComposeDecodeHookFunc_kinds(t)
}

// TestTask003Repeat guards deterministic behavior across repeated calls.
func TestTask003Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask003(t)
	}
}
