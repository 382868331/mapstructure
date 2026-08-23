package mapstructure

import "testing"

// TestTask004 isolates the 解码取消 regression.
func TestTask004(t *testing.T) {
	TestOrComposeDecodeHookFunc(t)
}
