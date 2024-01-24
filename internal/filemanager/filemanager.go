package filemanager

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func ReadLines(fileToRead string) ([]string, error) {
	file, err := os.Open(fileToRead)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Could not open file: [%v]", fileToRead))
	}

	// make sure file is closed when done
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Could not close file: [%v]\n", fileToRead)
			fmt.Println(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)

	// for each line in the file (until EOF returns false)...
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// check for any error emitted by the scanner
	err = scanner.Err()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error reading file: [%v]", fileToRead))
	}

	return lines, nil
}
