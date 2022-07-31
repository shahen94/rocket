package logging

import (
	"log"
	"os"
)

func New() (*log.Logger, *log.Logger) {
	infoLog := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)

	return infoLog, errorLog
}
