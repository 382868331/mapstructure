package mapstructure

import "testing"

// TestTask007 isolates the Unicode字段名 regression.
func TestTask007(t *testing.T) {
	TestComposeDecodeHookFunc_safe_nofuncs(t)
}
