package mapstructure

import "testing"

// TestTask009 isolates the 键值结果 regression.
func TestTask009(t *testing.T) {
	TestStringToTimeDurationHookFunc(t)
}
