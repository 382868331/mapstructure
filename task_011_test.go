package mapstructure

import "testing"

func TestTask011ErrorUnusedReportsExtraKeys(t *testing.T) {
	for _, key := range []string{"Age", "Country"} {
		var out struct{ Name string }
		cfg := &DecoderConfig{Result: &out, ErrorUnused: true}
		d, err := NewDecoder(cfg)
		if err != nil { t.Fatal(err) }
		if err := d.Decode(map[string]interface{}{"Name": "Ada", key: 37}); err == nil {
			t.Fatalf("expected an error for unused key %s", key)
		}
	}
}
