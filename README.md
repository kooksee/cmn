# cmn
golang公共函数库

```
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
```