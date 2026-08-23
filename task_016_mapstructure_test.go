package mapstructure

import "testing"

// TestTask016 isolates the 解码状态恢复 regression.
func TestTask016(t *testing.T) {
	TestDecode_NilValue(t)
}
