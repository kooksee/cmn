package cmn

import (
	"errors"
	"fmt"
)

func MustNotErr(errs ... error) {
	for _, err := range errs {
		if err != nil {
			panic(err.Error())
		}
	}
}

func Err(data string, params ... interface{}) error {
	return errors.New(fmt.Sprintf(data, params...))
}

func ErrWithMsg(msg string, errs ... error) error {
	for _, err := range errs {
		if err != nil {
			return errors.New(fmt.Sprintf("%s --> %s", msg, err.Error()))
		}
	}
	return nil
}
