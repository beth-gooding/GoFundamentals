package customErrors

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	ErrInvalidSSNNumbers = errors.New("ssn has non-numeric digits")
)

func CheckValidCharacters(ssn string) error {
	_, err := strconv.Atoi(ssn)
	if err != nil {
		return fmt.Errorf("the value of %s caused an error: %v", ssn, ErrInvalidSSNNumbers)
	}
	return nil
}
