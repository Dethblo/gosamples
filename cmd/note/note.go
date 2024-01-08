package main

import (
	"bufio"
	"fmt"
	"github.com/Dethblo/godethblo/internal/note"
	"os"
	"strings"
)

func main() {

	userNote, err := getNoteData()
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		fmt.Println(err)
		return
	}

	fmt.Println("Saving the note succeeded!!.")

}

func getNoteData() (note.Note, error) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")
	theNote, err := note.New(title, content)
	return theNote, err
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	// remove the line feed
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
