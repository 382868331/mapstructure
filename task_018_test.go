package mapstructure

import "testing"

func TestTask018WeakMetadataConvertsValues(t *testing.T) {
	var out struct{ Count int }
	md := new(Metadata)
	if err := WeakDecodeMetadata(map[string]interface{}{"Count": "7", "Extra": 1}, &out, md); err != nil {
		t.Fatal(err)
	}
	if out.Count != 7 || len(md.Unused) != 1 {
		t.Fatalf("out=%#v metadata=%#v", out, md)
	}
	var second struct{ Enabled bool }
	md2 := new(Metadata)
	if err := WeakDecodeMetadata(map[string]interface{}{"Enabled": "1"}, &second, md2); err != nil {
		t.Fatal(err)
	}
	if !second.Enabled {
		t.Fatal("weak bool conversion was not applied")
	}
}
