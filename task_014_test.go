package mapstructure
import "testing"
func TestTask014IgnoreUntaggedFields(t *testing.T){
	in:=struct{Visible string `mapstructure:"visible"`; Hidden string}{Visible:"yes",Hidden:"no"}
	out:=map[string]interface{}{}
	d,err:=NewDecoder(&DecoderConfig{Result:&out,IgnoreUntaggedFields:true});if err!=nil{t.Fatal(err)}
	if err=d.Decode(in);err!=nil{t.Fatal(err)}
	if out["visible"]!="yes"{t.Fatalf("out=%#v",out)}
	if _,ok:=out["Hidden"];ok{t.Fatalf("untagged field leaked: %#v",out)}
}
