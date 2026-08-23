package mapstructure
import "testing"
func TestTask001WeakBoolToString(t *testing.T){var out string;if err:=WeakDecode(true,&out);err!=nil{t.Fatal(err)};if out!="1"{t.Fatalf("out=%q want=1",out)}}
