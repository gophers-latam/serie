package app

import (
	"log"
	"os"
)

type Application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}

func NewInfoLogger() *log.Logger {
	return log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
}

func NewErrorLogger() *log.Logger {
	return log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
}
