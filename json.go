package cmn

import "github.com/json-iterator/go"

var Json = myJson{}

type myJson struct{}

func (myJson) Get(data []byte, path ...interface{}) jsoniter.Any {
	return json.Get(data, path...)
}

func (myJson) MarshalToString(v interface{}) (string, error) {
	return json.MarshalToString(v)
}

func (myJson) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (myJson) MarshalStructWithSorted(s interface{}) ([]byte, error) {
	s1, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	b := map[string]interface{}{}
	if err := json.Unmarshal(s1, &b); err != nil {
		return nil, err
	}

	b1, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	return b1, nil
}

func (myJson) UnmarshalFromString(str string, v interface{}) error {
	return json.UnmarshalFromString(str, v)
}

func (myJson) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (myJson) Valid(data []byte) bool {
	return json.Valid(data)
}
