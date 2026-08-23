package mapstructure

import "testing"

// TestTask005 isolates the hook收尾 regression.
func TestTask005(t *testing.T) {
	TestOrComposeDecodeHookFunc_correctValueIsLast(t)
}
