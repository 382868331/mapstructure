package mapstructure

import "testing"

func TestTask008DefaultTagName(t *testing.T) {
	var out struct {
		Value int `mapstructure:"custom"`
		Other int `mapstructure:"second"`
	}
	if err := Decode(map[string]interface{}{"custom": 7, "second": 9}, &out); err != nil {
		t.Fatal(err)
	}
	if out.Value != 7 || out.Other != 9 {
		t.Fatalf("out=%+v", out)
	}
}
