package mapstructure

import "testing"

// TestTask020 isolates the 解码器版本 regression.
func TestTask020(t *testing.T) {
	TestDecodeSliceToEmptySliceWOZeroing(t)
}
