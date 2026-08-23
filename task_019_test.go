package mapstructure
import("reflect";"strconv";"testing")
func TestTask019ComposedHooksRunInOrder(t *testing.T){
	h1:=DecodeHookFuncType(func(f,t reflect.Type,data interface{})(interface{},error){if f.Kind()!=reflect.String{return data,nil};return data.(string)+"0",nil})
	h2:=DecodeHookFuncType(func(f,t reflect.Type,data interface{})(interface{},error){if f.Kind()!=reflect.String{return data,nil};return strconv.Atoi(data.(string))})
	var out struct{Count int};d,err:=NewDecoder(&DecoderConfig{Result:&out,DecodeHook:ComposeDecodeHookFunc(h1,h2)});if err!=nil{t.Fatal(err)}
	if err=d.Decode(map[string]interface{}{"Count":"7"});err!=nil{t.Fatal(err)}
	if out.Count!=70{t.Fatalf("count=%d",out.Count)}
	var second struct{Count int};d2,err:=NewDecoder(&DecoderConfig{Result:&second,DecodeHook:ComposeDecodeHookFunc(h1,h2)});if err!=nil{t.Fatal(err)}
	if err=d2.Decode(map[string]interface{}{"Count":"12"});err!=nil{t.Fatal(err)}
	if second.Count!=120{t.Fatalf("second count=%d",second.Count)}
}
