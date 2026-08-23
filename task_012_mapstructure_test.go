package mapstructure

import "testing"

// TestTask012 isolates the ANSI字段名 regression.
func TestTask012(t *testing.T) {
	TestStringToIPNetHookFunc(t)
}
