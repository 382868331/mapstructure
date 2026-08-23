package mapstructure

import (
	"strings"
	"testing"
)

func TestTask020DecoderExplainsNonPointerResult(t *testing.T) {
	_, err := NewDecoder(&DecoderConfig{Result: struct{}{}})
	if err == nil || !strings.Contains(err.Error(), "pointer") {
		t.Fatalf("error=%v", err)
	}
	_, err = NewDecoder(&DecoderConfig{Result: map[string]int{}})
	if err == nil || !strings.Contains(err.Error(), "pointer") {
		t.Fatalf("map error=%v", err)
	}
	_, err = NewDecoder(&DecoderConfig{Result: "value"})
	if err == nil || !strings.Contains(err.Error(), "pointer") {
		t.Fatalf("string error=%v", err)
	}
}
