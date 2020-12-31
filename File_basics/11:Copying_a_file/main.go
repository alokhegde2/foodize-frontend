//we are using io.Copy() function to copy the content

package main

import (
	"io"
	"log"
	"os"
)

func main() {
	//Open original file
	originalFile, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	//A defer statement defers the execution of a function until the surrounding function returns.
	//The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding
	//function returns.
	defer originalFile.Close()

	//create a new file
	newFile, err := os.Create("test_copy.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	//Copy the bytes to destination from source
	bytesWritten, err := io.Copy(newFile, originalFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes.", bytesWritten)

	//Commit the file contents
	//Flushes memory to disk
	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
	}
}
