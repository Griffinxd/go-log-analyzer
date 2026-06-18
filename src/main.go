package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

type LogLevel string

const StringSliceCapacity = 100

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

	file, err := os.Open(FileName)
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string = make([]string, StringSliceCapacity)

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	for _, line := range fileLines {
		fmt.Println(line)
	}
	// fmt.Println(string(data))

}

func main() {

	args := os.Args
	argc := len(args)
	if argc != 2 {
		log.Fatalf("Invalid number of arguments")
	}

	fmt.Println("Command Line Arguments")
	// for _, arg := range args {
	// 	fmt.Println(arg)
	// }

	fileName := args[1]
	fmt.Println("Filename: " + fileName)

	ReadFile(fileName)
}
