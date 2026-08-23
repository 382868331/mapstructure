package mapstructure
import "testing"
func TestTask009DefaultCaseInsensitiveMatch(t *testing.T){var out struct{Name string};if err:=Decode(map[string]interface{}{"name":"Ada"},&out);err!=nil{t.Fatal(err)};if out.Name!="Ada"{t.Fatalf("name=%q",out.Name)}}
