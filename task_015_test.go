package mapstructure
import "testing"
func TestTask015ZeroFieldsReplacesMap(t *testing.T){
	out:=map[string]int{"stale":9}
	d,err:=NewDecoder(&DecoderConfig{Result:&out,ZeroFields:true});if err!=nil{t.Fatal(err)}
	if err=d.Decode(map[string]int{"fresh":2});err!=nil{t.Fatal(err)}
	if _,ok:=out["stale"];ok{t.Fatalf("stale key retained: %#v",out)}
	if out["fresh"]!=2{t.Fatalf("out=%#v",out)}
}
