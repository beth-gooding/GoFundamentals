package main

import (
	"log"
	"strings"

	"a11.01/pkg/customErrors"
)

func main() {
	validateSSN := []string{"123-45-6789", "012-8-678", "000-12-0962", "999-33-3333", "087-65-4321", "123-45-zzzz"}
	numSSN := len(validateSSN)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Printf("Validating data: %#v\n", validateSSN)

	for index, ssn := range validateSSN {
		processedSSN := strings.ReplaceAll(ssn, "-", "")
		log.Printf("Validate data %s, %d of %d\n", ssn, index+1, numSSN)

		err := customErrors.CheckValidLength(processedSSN)
		if err != nil {
			log.Println(err)
		}

		err = customErrors.CheckValidCharacters(processedSSN)
		if err != nil {
			log.Println(err)
		}

		err = customErrors.CheckValidPrefix(processedSSN)
		if err != nil {
			log.Println(err)
		}

		err = customErrors.CheckValidDigitPlace(processedSSN)
		if err != nil {
			log.Println(err)
		}
	}

}
