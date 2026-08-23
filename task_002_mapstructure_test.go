package mapstructure

import "testing"

// TestTask002 isolates the 目标切片 regression.
func TestTask002(t *testing.T) {
	TestComposeDecodeHookFunc_err(t)
}
