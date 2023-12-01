package util

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func ReadFile(fileName string) []string {
	fileHandle, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer fileHandle.Close()
	return scan(fileHandle)
}

func ReadString(text string) []string {
	return scan(strings.NewReader(text))
}

func scan(r io.Reader) []string {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
