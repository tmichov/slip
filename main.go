package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

type FileMap struct {
	fromIndex int
	toIndex int
	lines []string
	swap bool
}

func getFile(file *string) (*FileMap, error) {
	if *file == "" {
		return nil, fmt.Errorf("file is required")
	}

	fileContents, err := os.Open(*file)
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

func main() {
	//from := flag.String("from", "", "a string")	
	//to := flag.String("to", "", "a string")	
	file := flag.String("file", "", "a string")	
	//swap := flag.Bool("swap", false, "a bool")	
	
	flag.Parse()

	fileMap, err := getFile(file)	

	if err != nil {
		log.Fatal(err)
		return
	}

	fileMap.edit()

	fileContents, err := os.Create(*file)
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
