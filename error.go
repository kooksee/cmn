package cmn

import (
	"errors"
	"fmt"
	"reflect"
)

var Err = myErr{}

type myErr struct{}

// MustNotErr ,支持func()error,func()([]reflect.Value, error),error
func (m myErr) MustNotErr(errs ... interface{}) {
	if err := m.ErrWithMsg("myErr.MustNotErr", errs...); err != nil {
		panic(err.Error())
	}
}

func (m myErr) GetResultWithoutErr(data interface{}, err error) interface{} {
	m.MustNotErr(err)
	return data
}

// Err 封装errors
func (myErr) Err(data string, params ... interface{}) error {
	return errors.New(fmt.Sprintf(data, params...))
}

func (m myErr) ErrWithMsg(msg string, errs ... interface{}) error {
	for _, err := range errs {
		var ee error
		if err != nil {
			switch err.(type) {
			default:
				return m.Err("myErr.ErrWithMsg error params type: %s", reflect.TypeOf(err).String())
			case error:
				ee = err.(error)
			case func() error:
				ee = (err.(func() error))()
			case func() ([]reflect.Value, error):
				ee = m.FilterErr((err.(func() ([]reflect.Value, error)))())
			}

			if ee != nil {
				return errors.New(fmt.Sprintf("%s -> %s", msg, ee.Error()))
			}
		}
	}
	return nil
}

func (m myErr) Wrap(f interface{}, params ... interface{}) func() error {
	return func() error {
		return m.FilterErr(m.WrapM(f, params...)())
	}
}

func (myErr) FilterErr(params ... interface{}) error {
	if len(params) < 1 {
		return errors.New("err -> Wrap: the params must be more than one value")
	}

	p := params[len(params)-1]
	if p == nil {
		return nil
	}

	value := reflect.ValueOf(p)
	if value.IsNil() {
		return nil
	}

	if e, ok := value.Interface().(error); ok {
		return e
	} else {
		return errors.New("err -> Wrap: the last param must be error type")
	}
}

func (myErr) WrapM(f interface{}, params ... interface{}) func() (ds []reflect.Value, err error) {
	return func() (ds []reflect.Value, err error) {
		t := reflect.TypeOf(f)
		if t.Kind() != reflect.Func {
			return nil, errors.New("err -> Wrap: please input func")
		}

		var vs []reflect.Value
		for _, p := range params {
			vs = append(vs, reflect.ValueOf(p))
		}

		v := reflect.ValueOf(f)
		out := v.Call(vs)
		if len(out) < 1 {
			return nil, errors.New("err -> Wrap: the func output must be more than one value")
		}

		value := out[len(out)-1]
		if value.IsNil() {
			return out[:len(out)-1], nil
		}

		if e, ok := value.Interface().(error); ok {
			return nil, e
		} else {
			return nil, errors.New("err -> Wrap: the func last output must be error type")
		}
	}
}
