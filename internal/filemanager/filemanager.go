package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func ReadLines(fileToRead string) ([]string, error) {
	file, err := os.Open(fileToRead)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not open file: [%v]", fileToRead))
	}

	// make sure file is closed when done
	defer close(file)

	scanner := bufio.NewScanner(file)

	// for each line in the file (until EOF returns false)...
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// check for any error emitted by the scanner
	err = scanner.Err()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error reading file: [%v]", fileToRead))
	}

	return lines, nil
}

func WriteJson(fileToMake string, data interface{}) error {
	file, err := os.Create(fileToMake)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to create file: [%v]", fileToMake))
	}

	// make sure file is closed when done
	defer close(file)

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // pretty prints the JSON
	err = encoder.Encode(data)
	if err != nil {
		return errors.New("failed to convert data to JSON")
	}

	return nil
}

func close(file *os.File) {
	err := file.Close()
	if err != nil {
		fmt.Printf("could not close file")
		fmt.Println(err)
	}
}
