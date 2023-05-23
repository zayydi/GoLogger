package watcher

import (
	"bufio"
	"encoding/json"
	"fmt"
	"loggerStruct"
	"os"
	"path/filepath"
	"strings"

	"github.com/gologger/eslogger"
)

func analyzeLogFile(filePath string) error {
	// Check if the file extension is .txt or .log
	ext := strings.ToLower(filepath.Ext(filePath))
	if ext != ".txt" && ext != ".log" {
		return fmt.Errorf("unsupported file extension: %s", ext)
	}

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read and analyze each line
	for scanner.Scan() {
		line := scanner.Text()
		analyzeLine(line)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func analyzeLine(line string) {
	var log loggerStruct.Logger
	err := json.Unmarshal([]byte(line), &log)
	if err != nil {
		log.Level = "unknown"
		log.Description = line
	}

	eslogger.LogError(log)
}

func ReadFile(filePath string) {
	err := analyzeLogFile(filePath)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
