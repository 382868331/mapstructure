package mapstructure

import "testing"

// TestTask003 isolates the 弱类型数字 regression.
func TestTask003(t *testing.T) {
	TestComposeDecodeHookFunc_kinds(t)
}
