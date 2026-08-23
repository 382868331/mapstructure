package mapstructure
import "testing"
func TestTask008DefaultTagName(t *testing.T){var out struct{Value int `mapstructure:"custom"`};if err:=Decode(map[string]interface{}{"custom":7},&out);err!=nil{t.Fatal(err)};if out.Value!=7{t.Fatalf("value=%d",out.Value)}}
