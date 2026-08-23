package mapstructure

import "testing"

func TestTask011ErrorUnusedReportsExtraKeys(t *testing.T) {
	var out struct{ Name string }
	cfg := &DecoderConfig{Result: &out, ErrorUnused: true}
	d, err := NewDecoder(cfg)
	if err != nil { t.Fatal(err) }
	if err := d.Decode(map[string]interface{}{"Name": "Ada", "Age": 37}); err == nil {
		t.Fatal("expected an error for the unused Age key")
	}
}
