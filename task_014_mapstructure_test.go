package mapstructure

import "testing"

// TestTask014 isolates the 维度字段 regression.
func TestTask014(t *testing.T) {
	TestStructToMapHookFuncTabled(t)
}

// TestTask014Repeat guards deterministic behavior across repeated calls.
func TestTask014Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask014(t)
	}
}
