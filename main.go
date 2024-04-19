package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func getFile(file *string) (*os.File, error) {
	if *file == "" {
		return nil, fmt.Errorf("file is required")
	}

	fileContents, err := os.Open(*file)
	if err != nil {
		return nil, fmt.Errorf("file not found")
	}

	return fileContents, nil
}

func main() {
	//from := flag.String("from", "", "a string")	
	//to := flag.String("to", "", "a string")	
	file := flag.String("file", "", "a string")	
	//swap := flag.Bool("swap", false, "a bool")	
	
	mapFlag := make(map[string]string)

	flag.Parse()

	fileContents, err := getFile(file)	

	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(fileContents)

	i := 0
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)

		if i == 7 {
			mapFlag["from"] = line
		}

		if i == 3 {
			mapFlag["to"] = line
		}

		i++
	}

	//TODO: manipulate the lines with the flags
	// in the future, use regex to find the line to manipulate
	fmt.Println(lines)
}
