package customErrors

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidSSNLength = errors.New("ssn is not 9 characters long")
)

func CheckValidLength(ssn string) error {
	if len(ssn) != 9 {
		return fmt.Errorf("the value of %s caused an error: %v", ssn, ErrInvalidSSNLength)
	}
	return nil
}
