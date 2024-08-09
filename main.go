package main

import (
	"bufio"
	"fmt"
	"os"
)

func countRows(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// Cria um scanner para ler o arquivo linha por linha
	scanner := bufio.NewScanner(file)

	//Contador de linhas
	count := 0
	for scanner.Scan() {
		count++
	}

	return count, nil
}

func main() {
	files := []string{"file1.txt", "file2.txt", "file3.txt"}
	rowChannel := make(chan int, len(files))

	for _, file := range files {
		fileToRead := file

		go func() {
			fmt.Println("Reading go routine")
			numRows, err := countRows(fileToRead)
			if err != nil {
				panic("ops")
			}

			rowChannel <- numRows
		}()
	}

	var totalRows int

	for range files {
		rows := <-rowChannel
		totalRows += rows
	}

	fmt.Println(totalRows)
}
