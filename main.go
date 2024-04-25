package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type FileMap struct {
	fromIndex int
	toIndex int
	lines []string
	swap bool
}

func getFile(allArgs *Args) (*FileMap, error) {
	if allArgs.file == "" {
		return nil, fmt.Errorf("file is required")
	}

	fileContents, err := os.Open(allArgs.file)
	if err != nil {
		return nil, fmt.Errorf("file not found")
	}

	fileMap := FileMap{
		fromIndex: -1,
		toIndex: -1,
		lines: []string{},
		swap: true,
	}

	scanner := bufio.NewScanner(fileContents)

	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		fileMap.lines = append(fileMap.lines, line)

		if i == 6 {
			fileMap.fromIndex = i
		}

		if i == 3 {
			fileMap.toIndex = i
		}

		i++
	}

	return &fileMap, nil
}

func (f *FileMap) edit() {
	if f.fromIndex == -1 || f.toIndex == -1 {
		return
	}

	var temp string
	if f.swap {
		temp = f.lines[f.toIndex]
	}

	f.lines[f.toIndex] = f.lines[f.fromIndex]

	if f.swap {
		f.lines[f.fromIndex] = temp
	} else {
		f.lines = append(f.lines[:f.fromIndex], f.lines[f.fromIndex+1:]...)
	}
}

type Options struct {
	fromIndex int
	toIndex int
	swap bool
	fromRegex string
	toRegex string
}

type Args struct {
	from string
	to string
	swap bool
	file string
}

func parseArgs() (*Args, error) {
	var allArgs Args
	args := os.Args[1:]

	if len(args) < 2 {
		return nil, fmt.Errorf("Usage: go run main.go filename.txt --from=1 --to=2 --swap")
	}

	for _, arg := range args {
		if strings.HasPrefix(arg, "--from") {
			allArgs.from = strings.Split(arg, "=")[1]
		}

		if strings.HasPrefix(arg, "--to") {
			allArgs.to = strings.Split(arg, "=")[1]
		}

		if strings.HasPrefix(arg, "--swap") {
			allArgs.swap = true
		}

		if !strings.HasPrefix(arg, "--") {
			allArgs.file = arg
		}

		if strings.HasPrefix(arg, "--file") {
			allArgs.file = strings.Split(arg, "=")[1]
		}
	}

	return &allArgs, nil
}

func main() {

	allArgs, err := parseArgs()

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(allArgs.file)

	fromIndex := parseInt(allArgs.from)
	toIndex := parseInt(allArgs.to)

	fmt.Println(fromIndex, toIndex)

	fileMap, err := getFile(allArgs)

	if err != nil {
		log.Fatal(err)
		return
	}

	fileMap.edit()

	fileContents, err := os.Create(allArgs.file)

	if err != nil {
		log.Fatal(err)
		return
	}

	for _, line := range fileMap.lines {
		fileContents.WriteString(line + "\n")
	}

	fileContents.Close()

	fmt.Println("File edited successfully")
}

func parseInt(input string) int {
	var i int

	fmt.Println(input)
	if _, err := fmt.Sscanf(input, "%d", &i); err == nil {
		return i
	}

	return -1
}

