package main 

import (
	"flag"
	"fmt"
)

//TODO: Implement grep, find all files that match the pattern
func grep(file string) []string {
	return []string{}
}

func main() {
	from := flag.String("from", "", "a string")	
	to := flag.String("to", "", "a string")	
	file := flag.String("file", "", "a string")	
	
	flag.Parse()

	fmt.Println("from:", *from)
	fmt.Println("to:", *to)
	fmt.Println("file:", *file)
}
