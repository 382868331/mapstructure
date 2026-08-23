package mapstructure

import "testing"

// TestTask019 isolates the 空结构 regression.
func TestTask019(t *testing.T) {
	TestNestedTypeWithDefaults(t)
}
