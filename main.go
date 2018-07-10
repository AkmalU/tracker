package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/AkmalUr/test/transactions"
)

func main() {
	args := os.Args[1:]
	filePath := args[0]
	resultFilePath := args[1]

	inLines, err := readLines(filePath)
	if err != nil {
		log.Fatal(err)
	}

	service := transactions.NewService(transactions.NewInMemoryRepo())
	outLines := [][]byte{}
	for _, inLineStr := range inLines {
		inLine := []byte(inLineStr)

		transaction := &transactions.Transaction{}
		err := json.Unmarshal(inLine, transaction)
		if err != nil {
			fmt.Printf("boo\n%s\n\n------\n%+v", string(inLine), err)
			os.Exit(-1)
		}

		result := service.ProcessTransaction(transaction)
		if result == nil {
			continue
		}
		outLine, err := json.Marshal(result)
		if err != nil {
			log.Fatal(err)
		}
		outLines = append(outLines, outLine)
	}

	err = writeLines(outLines, resultFilePath)
	if err != nil {
		log.Fatal(err)
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func writeLines(lines [][]byte, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, string(line))
	}
	return w.Flush()
}
