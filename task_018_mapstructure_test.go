package mapstructure

import "testing"

// TestTask018 isolates the 字段缓存 regression.
func TestTask018(t *testing.T) {
	TestNestedTypeSliceWithDefaults(t)
}
