package mapstructure

import "testing"

func TestTask009DefaultCaseInsensitiveMatch(t *testing.T) {
	for _, key := range []string{"name", "NAME"} {
		var out struct{ Name string }
		if err := Decode(map[string]interface{}{key: "Ada"}, &out); err != nil {
			t.Fatal(err)
		}
		if out.Name != "Ada" {
			t.Fatalf("key=%q name=%q", key, out.Name)
		}
	}
}
