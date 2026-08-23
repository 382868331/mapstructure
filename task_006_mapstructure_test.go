package mapstructure

import "testing"

// TestTask006 isolates the 字段标签去重 regression.
func TestTask006(t *testing.T) {
	TestOrComposeDecodeHookFunc_err(t)
}
