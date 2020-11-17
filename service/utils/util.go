package utils

import (
	"log"
)

func PrintError(err error, ignore bool) {
	if err != nil {
		if ignore {
			log.Println("Add failed:", err)
		} else {
			log.Fatal("Add failed:", err)
		}
	}
}
