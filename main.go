package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type FileMap struct {
	lines []string
}

type Args struct {
	swap bool
	file string
	fromIndex int
	toIndex int
	fromRegex string
	toRegex string
}

func main() {
	allArgs, err := parseArgs()

	if err != nil {
		log.Fatal(err)
		return
	}

	fileMap, err := getFile(allArgs)

	if err != nil {
		log.Fatal(err)
		return
	}

	fileMap.edit(allArgs)
	
	fileContents, err := os.Create(allArgs.file)

	if err != nil {
		log.Fatal(err)
		return
	}

	defer fileContents.Close()

	for _, line := range fileMap.lines {
		fileContents.WriteString(line + "\n")
	}

	fmt.Println("File edited successfully")
}

func parseArgs() (*Args, error) {
	var allArgs Args
	args := os.Args[1:]

	if len(args) < 2 {
		return nil, fmt.Errorf("Usage: go run main.go filename.txt --from=1 --to=2 --swap")
	}

	for _, arg := range args {
		if strings.HasPrefix(arg, "--from") {
			from := strings.Split(arg, "=")[1]
			allArgs.fromIndex = parseInt(from)
			if allArgs.fromIndex == -1 {
				allArgs.fromRegex = from
			}
		}

		if strings.HasPrefix(arg, "--to") {
			to := strings.Split(arg, "=")[1]
			allArgs.toIndex = parseInt(to)
			if allArgs.toIndex == -1 {
				allArgs.toRegex = to 
			}
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

func getFile(allArgs *Args) (*FileMap, error) {
	if allArgs.file == "" {
		return nil, fmt.Errorf("file is required")
	}

	fileContents, err := os.Open(allArgs.file)
	if err != nil {
		return nil, fmt.Errorf("file not found")
	}

	defer fileContents.Close()

	lines := []string{}

	scanner := bufio.NewScanner(fileContents)

	i := 0
	
	var regFrom *regexp.Regexp
	var regTo *regexp.Regexp

	if allArgs.fromRegex != "" {
		regFrom, err = regexp.Compile(allArgs.fromRegex)
		if err != nil {
			return nil, fmt.Errorf("Invalid regex")
		}
	}

	if allArgs.toRegex != "" {
		regTo, err = regexp.Compile(allArgs.toRegex)
		if err != nil {
			return nil, fmt.Errorf("Invalid regex")
		}
	}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		if allArgs.fromRegex != "" && allArgs.fromIndex == -1 {
			if match := regFrom.Find([]byte(line)); match != nil {
				allArgs.fromIndex = i
			}
		}

		if allArgs.toRegex != "" && allArgs.toIndex == -1 {
			if match := regTo.Find([]byte(line)); match != nil {
				allArgs.toIndex = i
			}
		}

		i++
	}

	return &FileMap{
		lines: lines,
	}, nil
}

func (f *FileMap) edit(o *Args) {
	if o.fromIndex == -1 || o.toIndex == -1 {
		return
	}

	var temp string
	if o.swap {
		temp = f.lines[o.toIndex]
	}

	f.lines[o.toIndex] = f.lines[o.fromIndex]

	if o.swap {
		f.lines[o.fromIndex] = temp
	} else {
		f.lines = append(f.lines[:o.fromIndex], f.lines[o.fromIndex+1:]...)
	}
}

func parseInt(input string) int {
	var i int

	if _, err := fmt.Sscanf(input, "%d", &i); err == nil {
		return i
	}

	return -1
}

