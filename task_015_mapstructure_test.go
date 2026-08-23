package mapstructure

import "testing"

// TestTask015 isolates the 目标句柄 regression.
func TestTask015(t *testing.T) {
	TestTextUnmarshallerHookFunc(t)
}
