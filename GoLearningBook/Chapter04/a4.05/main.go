package main

import (
	"fmt"
	"os"
)

type locale struct {
	language string
	country  string
}

func createValidLocales() []locale {
	enUS := locale{
		"en",
		"US",
	}

	enUK := locale{
		"en",
		"GB",
	}

	frCN := locale{
		"fr",
		"CN",
	}

	frFR := locale{
		"fr",
		"FR",
	}

	ruRu := locale{
		"ru",
		"RU",
	}

	var validLocales = []locale{enUK, enUS, frCN, frFR, ruRu}
	return validLocales
}

func checkValidLocale(input string) bool {
	validLocales := createValidLocales()

	potentialLocale := locale{
		input[0:2],
		input[3:],
	}

	for i := 0; i < len(validLocales); i++ {
		if potentialLocale == validLocales[i] {
			return true
		}
	}

	return false

}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No locale passed to command line")
		os.Exit(1)
	}
	isValid := checkValidLocale(os.Args[1])

	if isValid {
		fmt.Println("Locale passed is supported")
	} else {
		fmt.Printf("Locale not support %v\n", os.Args[1])
	}
}
