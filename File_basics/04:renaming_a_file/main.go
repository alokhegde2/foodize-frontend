//The standard library provides a convenient function for moving a file.
//Renaming and moving file are synonymous terms;
//If we want to move a file from one directory to another,use os.Rename() function

package main

import (
	"log"
	"os"
)

func main() {
	originalPath := "test.txt"
	newPath := "test2.txt"
	err := os.Rename(originalPath, newPath)
	if err != nil {
		log.Fatal(err)
	}
}
