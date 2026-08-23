package mapstructure

import "testing"

// TestTask011 isolates the 并发元数据 regression.
func TestTask011(t *testing.T) {
	TestStringToIPHookFunc(t)
}
