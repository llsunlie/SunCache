package log

import (
	"log"
	"os"
)

var (
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	// filename := "log_" + time.Now().String() + ".log"
	filename := "log" + ".log"
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	InfoLogger = log.New(file, "[Info] ", log.Ldate|log.Ltime)
	WarningLogger = log.New(file, "[Warning] ", log.Ldate|log.Ltime)
	ErrorLogger = log.New(file, "[Error] ", log.Ldate|log.Ltime)
}

func Info(format string, a ...any) {
	InfoLogger.Printf(format, a...)
}

func Infoln(a ...any) {
	InfoLogger.Println(a...)
}

func Warning(format string, a ...any) {
	WarningLogger.Printf(format, a...)
}

func Warningln(a ...any) {
	WarningLogger.Println(a...)
}

func Error(format string, a ...any) {
	ErrorLogger.Panicf(format, a...)
}

func Errorln(a ...any) {
	ErrorLogger.Panic(a...)
}
