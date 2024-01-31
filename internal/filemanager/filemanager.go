package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not open file: [%v]", fm.InputFilePath))
	}

	// make sure file is closed when done
	defer CloseIt(file)

	scanner := bufio.NewScanner(file)

	// for each line in the file (until EOF returns false)...
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// check for any error emitted by the scanner
	err = scanner.Err()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error reading file: [%v]", fm.InputFilePath))
	}

	return lines, nil
}

func (fm FileManager) WriteResult(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to create file: [%v]", fm.OutputFilePath))
	}

	// artificially add a sleep to help illustrate usage of goroutines
	time.Sleep(3 * time.Second)

	// make sure file is closed when done
	defer CloseIt(file)

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // pretty prints the JSON
	err = encoder.Encode(data)
	if err != nil {
		return errors.New("failed to convert data to JSON")
	}

	return nil
}

func CloseIt(file *os.File) {
	err := file.Close()
	if err != nil {
		fmt.Printf("could not CloseIt file")
		fmt.Println(err)
	}
}
