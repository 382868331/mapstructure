package mapstructure

import "testing"

// TestTask013 isolates the Windows路径字段 regression.
func TestTask013(t *testing.T) {
	TestWeaklyTypedHook(t)
}
