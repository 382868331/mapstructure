package mapstructure

import "testing"

// TestTask017 isolates the 标签匹配 regression.
func TestTask017(t *testing.T) {
	TestNestedTypePointerWithDefaults(t)
}
