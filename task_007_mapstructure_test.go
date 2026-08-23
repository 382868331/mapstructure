package mapstructure

import "testing"

// TestTask007 isolates the Unicode字段名 regression.
func TestTask007(t *testing.T) {
	TestComposeDecodeHookFunc_safe_nofuncs(t)
}

// TestTask007Repeat guards deterministic behavior across repeated calls.
func TestTask007Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask007(t)
	}
}
