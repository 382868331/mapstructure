package mapstructure

import "testing"

// TestTask009 isolates the 键值结果 regression.
func TestTask009(t *testing.T) {
	TestStringToTimeDurationHookFunc(t)
}

// TestTask009Repeat guards deterministic behavior across repeated calls.
func TestTask009Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask009(t)
	}
}
