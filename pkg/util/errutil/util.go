package errutil

import (
	"errors"
	"fmt"
	"log"
)

func Error(code string, text string) error {
	return errors.New(fmt.Sprintf("%s %s", code, text))
}

func TryPaincError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func TryFatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
