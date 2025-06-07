package logger

import (
	"log"
)

var debugMode bool

func Init(debug bool) {
	debugMode = debug
}

func Info(msg string) {
	log.Printf("[INFO] %s", msg)
}

func Debug(msg string) {
	if debugMode {
		log.Printf("[DEBUG] %s", msg)
	}
}

func Error(err error) {
	log.Printf("[ERROR] %v", err)
}
