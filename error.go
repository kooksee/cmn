package cmn

import (
	"errors"
	"fmt"
	"reflect"
)

var Err = myErr{}

type myErr struct{}

func (myErr) MustNotErr(errs ... error) {
	for _, err := range errs {
		if err != nil {
			panic(err.Error())
		}
	}
}

func (m myErr) GetResultWithoutErr(data interface{}, err error) interface{} {
	m.MustNotErr(err)
	return data
}

func (myErr) Err(data string, params ... interface{}) error {
	return errors.New(fmt.Sprintf(data, params...))
}

func (myErr) ErrWithMsg(msg string, errs ... error) error {
	for _, err := range errs {
		if err != nil {
			return errors.New(fmt.Sprintf("%s -> %s", msg, err.Error()))
		}
	}
	return nil
}

func (myErr) Wrap(f interface{}, params ... interface{}) func() error {
	return func() error {
		t := reflect.TypeOf(f)
		if t.Kind() != reflect.Func {
			return errors.New("err -> Wrap: please input func")
		}

		var vs []reflect.Value
		for _, p := range params {
			vs = append(vs, reflect.ValueOf(p))
		}

		v := reflect.ValueOf(f)
		out := v.Call(vs)
		if len(out) != 1 {
			return errors.New("err -> Wrap: the func output must one value")
		}
		return out[0].Interface().(error)
	}
}
