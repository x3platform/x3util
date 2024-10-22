package errutil

import (
	"log"
)

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
