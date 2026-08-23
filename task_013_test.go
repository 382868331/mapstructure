package mapstructure
import "testing"
func TestTask013RemainCollectsUnusedKeys(t *testing.T){
	var out struct{Name string; Extra map[string]interface{} `mapstructure:",remain"`}
	if err:=Decode(map[string]interface{}{"Name":"Ada","Age":37},&out);err!=nil{t.Fatal(err)}
	if out.Extra["Age"]!=37{t.Fatalf("extra=%#v",out.Extra)}
}
