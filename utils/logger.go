package utils

import (
	"log"
	"os"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	DebugLogger *log.Logger
)

func init() {
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	DebugLogger = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(message string) {
	InfoLogger.Println(message)
}

func Error(message string, err error) {
	if err != nil {
		ErrorLogger.Printf("%s: %v\n", message, err)
	} else {
		ErrorLogger.Println(message)
	}
}

func Debug(message string) {
	DebugLogger.Println(message)
}
