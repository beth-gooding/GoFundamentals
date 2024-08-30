package customErrors

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrInvalidSSNPrefix = errors.New("ssn has three zeros as prefix")
)

func CheckValidPrefix(ssn string) error {
	if strings.HasPrefix(ssn, "000") {
		return fmt.Errorf("the value %s has caused an error: %v", ssn, ErrInvalidSSNPrefix)
	}
	return nil
}
