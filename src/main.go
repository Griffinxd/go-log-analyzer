package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type LogLevel string

const (
	LevelInfo    LogLevel = "INFO"
	LevelWarning LogLevel = "WARNING"
	LevelError   LogLevel = "ERROR"
)

type LogEntry struct {
	Timestamp time.Time
	Level     LogLevel
	Message   string
}

type LogFile struct {
	Entries []LogEntry
	Path    string
}

func ReadFile(FileName string) {

	data, err := os.ReadFile(FileName)
	if err != nil {
		log.Fatalf("Error reading file: %s", FileName)
	}

	fmt.Println(string(data))

}

func main() {

	args := os.Args
	argc := len(args)
	if argc != 2 {
		log.Fatalf("Invalid number of arguments")
	}

	fmt.Println("Command Line Arguments")
	for _, arg := range args {
		fmt.Println(arg)
	}

	fmt.Println("#################################################\n\n")
	fileName := args[1]
	fmt.Println("Filename: " + fileName)

	ReadFile(fileName)
}
