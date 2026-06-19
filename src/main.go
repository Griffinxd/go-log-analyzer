package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
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

type Formatter interface {
	Format(entry LogEntry) string
}

type TextFormatter struct {
}

type JSONFormatter struct{}

type CSVFormatter struct{}

type Stats struct {
	TotalEntries    int
	InfoCount       int
	WarningCount    int
	ErrorCount      int
	MostCommonError string
}

func (lf *LogFile) GetStats() Stats {

	var stats Stats
	stats.TotalEntries = len(lf.Entries)
	for i, log := range lf.Entries {
		fmt.Println(log, i)
		switch log.Level {
		case LevelError:
			stats.ErrorCount++
		case LevelWarning:
			stats.WarningCount++
		case LevelInfo:
			stats.InfoCount++
		}
	}
	return stats
}

func (j JSONFormatter) Format(entry LogEntry) string {
	data, err := json.MarshalIndent(entry, "", "\t")
	if err != nil {
		panic("json format error")
	}
	return string(data)
}

func (c CSVFormatter) Format(entry LogEntry) string {
	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	err := w.Write([]string{entry.Timestamp.String(), string(entry.Level), entry.Message})
	if err != nil {
		panic("csv format error")
	}
	w.Flush()
	return buf.String()
}

func (t TextFormatter) Format(entry LogEntry) string {
	return fmt.Sprintf("[%v] %s %s", entry.Timestamp, entry.Level, entry.Message)
}

func ReadFile(FileName string) []string {
	file, err := os.Open(FileName)
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	// for _, line := range fileLines {
	// 	fmt.Println(line)
	// }
	// fmt.Println(string(data))
	return fileLines
}

func CheckLine(Line string) LogEntry {
	var entry LogEntry
	// Line[0] should be '['
	// Line [20] should be ']'
	// Check Timestamp
	timeStampStr := Line[0:21]
	// fmt.Println(timeStampStr)
	if timeStampStr[0] != '[' && timeStampStr[20] != ']' {
		panic("timestamp is not in correct format ")
	}
	timeStampStr = strings.Trim(timeStampStr, "[]")
	t, err := time.Parse("2006-01-02 15:04:05", timeStampStr)
	if err != nil {
		panic(err)
	}
	entry.Timestamp = t

	// After Stimestamp
	remainingLine := Line[22:]
	// fmt.Println(remainingLine)

	strings := strings.SplitN(remainingLine, " ", 2)
	// fmt.Printf("Level: \"%s\"| Message: \"%s\"", strings[0], strings[1])
	entry.Level = LogLevel(strings[0])
	entry.Message = strings[1]
	// for i := 0; i < len(Line); i++ {
	// 	// fmt.Printf("%d, %c\n", i, Line[i])
	// }
	// fmt.Printf("%v \n", entry)
	return entry
}

func ParseLines(FileLines []string) []LogEntry {
	lineCount := len(FileLines)
	entries := make([]LogEntry, 0, lineCount)
	for _, line := range FileLines {
		entries = append(entries, CheckLine(line))
		// fmt.Println(entries[i])
	}
	// fmt.Println("timestamp:", entry.Timestamp)
	// fmt.Println("level:", entry.Level)
	// fmt.Println("message:", entry.Message)
	// Parsing logic
	// for _, line := range FileLines {
	//
	// 	// fmt.Printf("line: %d, log: %s\n", i, line)
	// 	// var entry LogEntry
	// 	CheckLine(line)
	// }
	// fmt.Printf("%v", entries[0])
	return entries
}

func main() {

	args := os.Args
	argc := len(args)
	if argc != 2 {
		log.Fatalf("Invalid number of arguments")
	}

	// fmt.Println("Command Line Arguments")
	// for _, arg := range args {
	// 	fmt.Println(arg)
	// }

	var logfile LogFile
	logfile.Path = args[1]
	// fmt.Println("Filename: " + fileName)
	logfile.Entries = ParseLines(ReadFile(logfile.Path))
	// text := TextFormatter{}
	// jsonForm := JSONFormatter{}
	// csvformat := CSVFormatter{}
	//
	// fmt.Println(text.Format(logfile.Entries[6]))
	// fmt.Println(jsonForm.Format(logfile.Entries[6]))
	// fmt.Println(csvformat.Format(logfile.Entries[6]))
	stats := logfile.GetStats()
	fmt.Println(stats)
}
