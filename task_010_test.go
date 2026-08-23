package mapstructure
import "testing"
func TestTask010SquashTag(t *testing.T){type inner struct{Value int};var out struct{inner `mapstructure:",squash"`};if err:=Decode(map[string]interface{}{"Value":8},&out);err!=nil{t.Fatal(err)};if out.Value!=8{t.Fatalf("value=%d",out.Value)}}
