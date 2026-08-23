package mapstructure
import "testing"
func TestTask010SquashTag(t *testing.T){
	type inner struct{Value int}
	for _, value := range []int{8, 13} {
		var out struct{inner `mapstructure:",squash"`}
		if err:=Decode(map[string]interface{}{"Value":value},&out);err!=nil{t.Fatal(err)}
		if out.Value!=value{t.Fatalf("value=%d, want %d",out.Value,value)}
	}
}
