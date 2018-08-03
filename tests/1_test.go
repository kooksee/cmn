package tests

import (
	"testing"
	"github.com/kooksee/cmn"
	"fmt"
	"reflect"
)

type Js1 struct {
	A string `json:"a"`
	C string `json:"c"`
	D string `json:"d"`
	M string `json:"m"`
	I string `json:"i"`
}

type Js struct {
	A string `json:"a"`
	C string `json:"c"`
	D string `json:"d"`
	Q Js1    `json:"q"`
	M string `json:"m"`
	I string `json:"i"`
}

func TestName(t *testing.T) {
	cmn.P(cmn.Color.Blue("ffff"))
	cmn.P(cmn.Rand.RandStr(4))

	cmn.Assert.StrEqual(cmn.Str(cmn.Err.GetResultWithoutErr(cmn.Json.Marshal(&Js{
		A: "ddd", C: "dd", D: "dd", M: "dd", I: "fff", Q: Js1{
			A: "ddd", C: "dd", D: "dd", M: "dd", I: "fff",
		},
	}))), `{"a":"ddd","c":"dd","d":"dd","q":{"a":"ddd","c":"dd","d":"dd","m":"dd","i":"fff"},"m":"dd","i":"fff"}`)

	cmn.Assert.StrEqual(cmn.Str(cmn.Err.GetResultWithoutErr(cmn.Json.MarshalStructWithSorted(&Js{
		A: "ddd", C: "dd", D: "dd", M: "dd", I: "fff", Q: Js1{
			A: "ddd", C: "dd", D: "dd", M: "dd", I: "fff",
		},
	}))), `{"a":"ddd","c":"dd","d":"dd","i":"fff","m":"dd","q":{"a":"ddd","c":"dd","d":"dd","i":"fff","m":"dd"}}`)

	f := cmn.Err.Wrap(cmn.Json.UnmarshalFromString, `{"a":"ddd","c":"dd","d":"dd","i":"fff","m":"dd","q":{"a":"ddd","c":"dd","d":"dd","i":"fff","m":"dd"}}`,&map[string]interface{}{})
	fmt.Println(f())

	reflect.TypeOf(nil)
}
