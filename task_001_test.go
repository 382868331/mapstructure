package mapstructure

import "testing"

func TestTask001WeakBoolToString(t *testing.T) {
	for _, tc := range []struct {
		in   bool
		want string
	}{{true, "1"}, {false, "0"}} {
		var out string
		if err := WeakDecode(tc.in, &out); err != nil {
			t.Fatal(err)
		}
		if out != tc.want {
			t.Fatalf("out=%q want=%q", out, tc.want)
		}
	}
}
