package logger

import "log"

func Fatal(err error) {
	log.Fatal(err)
}