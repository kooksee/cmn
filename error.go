package cmn

import (
	"errors"
	"fmt"
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
