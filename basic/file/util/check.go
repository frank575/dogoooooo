package util

import "log"

func CheckOpen(err error) {
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
}

func CheckGetwd(err error) {
	if err != nil {
		log.Fatalf("Error when getting wd: %s", err)
	}
}

func CheckRead(err error) {
	if err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}
}

func CheckReadDir(err error) {
	if err != nil {
		log.Fatalf("Error while reading dir: %s", err)
	}
}
