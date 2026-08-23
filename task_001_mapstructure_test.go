package mapstructure

import "testing"

// TestTask001 isolates the 空映射解码 regression.
func TestTask001(t *testing.T) {
	TestComposeDecodeHookFunc(t)
}
