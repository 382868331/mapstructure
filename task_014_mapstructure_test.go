package mapstructure

import "testing"

// TestTask014 isolates the 维度字段 regression.
func TestTask014(t *testing.T) {
	TestStructToMapHookFuncTabled(t)
}
