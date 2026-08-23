package mapstructure

import "testing"

// TestTask010 isolates the 字段遍历 regression.
func TestTask010(t *testing.T) {
	TestStringToTimeHookFunc(t)
}
