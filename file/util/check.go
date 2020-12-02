package util

import "log"

func CheckOpen(err error) {
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
}

func CheckRead(err error) {
	if err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}
}
