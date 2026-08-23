package mapstructure

import "testing"

// TestTask008 isolates the 目标容量 regression.
func TestTask008(t *testing.T) {
	TestStringToSliceHookFunc(t)
}
