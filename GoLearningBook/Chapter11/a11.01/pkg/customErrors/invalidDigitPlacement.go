package customErrors

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrInvalidSSNDigitPlace = errors.New("ssn starting with a 9 requires 7 or 9 in fourth place")
)

func CheckValidDigitPlace(ssn string) error {
	if strings.HasPrefix(ssn, "9") && string(ssn[3]) != "7" && string(ssn[3]) != "9" {
		return fmt.Errorf("the value of %s caused an error: %v", ssn, ErrInvalidSSNDigitPlace)
	}
	return nil
}
