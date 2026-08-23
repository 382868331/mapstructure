package mapstructure
import "testing"
func TestTask007MetadataTracksUnused(t *testing.T){var out struct{Known int};var md Metadata;if err:=DecodeMetadata(map[string]interface{}{"Known":1,"Extra":2},&out,&md);err!=nil{t.Fatal(err)};if len(md.Unused)!=1||md.Unused[0]!="Extra"{t.Fatalf("unused=%v",md.Unused)}}
