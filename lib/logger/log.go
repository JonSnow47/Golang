// package logger provide 5 simple method for log
package logger

import (
	"log"
)

func Debug(v ...interface{}) {
	log.Println("DEBUG:", v)
}

func Info(v ...interface{}) {
	log.Println("INFO:", v)
}

func Warning(v ...interface{}) {
	log.Println("WARN:", v)
}

func Error(v ...interface{}) {
	log.Println("ERROR:", v)
}

func Critical(v ...interface{}) {
	log.Println("CRITICAL:", v)
}
