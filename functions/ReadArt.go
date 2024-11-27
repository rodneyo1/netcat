package functions

import (
	"bufio"
	"fmt"
	"os"
)

func ReadArt(filename string) string {
	art := ""
	// fmt.Println(filename)
	filecontent, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file", err)
	}
	defer filecontent.Close()

	scanner := bufio.NewScanner(filecontent)

	for scanner.Scan() {
		art += scanner.Text() + "\n"
	}

	errs := scanner.Err()
	if errs != nil {
		fmt.Println("Error reading file", errs)
	}
	return art
}
