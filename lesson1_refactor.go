package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type ErrorWithTimestamp struct {
	text      string
	timestamp string
}

//исправил нейминг

func NewErrorWithTimestamp(text string) error {
	return &ErrorWithTimestamp{
		text:      text,
		timestamp: time.Now().String(),
	}
}

func (e *ErrorWithTimestamp) Error() string {
	return fmt.Sprintf("error: %s\noccur at: %s \n", e.text, e.timestamp)
}

func handlePanic() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Printf("Handle panic with details\n%v\n", New("Error from panic"))
		}
	}()

	fmt.Println("### A panic example ###")
	var arr []int
	fmt.Println(5 / len(arr))
}

//добавил проверку ошибки

func createAndCloseFile(index int) {
	f, err := os.Create(fmt.Sprintf("files/File_%d.txt", index))
	err := f.Close()
	if err != nil {
		return err
	}
	return nil

}

//добавил проверку ошибки

func createEmptyFiles() {
	fmt.Println("### Creating 1M of empty files ###")

	n := 1_000_000
	dir := "files"

	err := os.Mkdir(dir, 0700)
	if err != nil {
		return err
	}
	for i := 0; i < n; i++ {
		createAndCloseFile(i + 1)
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	fmt.Printf("Created %v files.\n", len(files))

	fmt.Println("Cleaning up...\n")
	err := os.RemoveAll(dir)
	if err != nil {
		return err
	}
}

func panicInGoroutine() {
	fmt.Println("### Handling panic in goroutine ###")

	go func() {
		defer func() {
			if v := recover(); v != nil {
				fmt.Println("Recovered from panic in goroutine:", v)
			}
		}()

		panic("A-A-A!!!")
	}()
	time.Sleep(time.Second)

}

func main() {
	handlePanic()
	createEmptyFiles()
	panicInGoroutine()
}
