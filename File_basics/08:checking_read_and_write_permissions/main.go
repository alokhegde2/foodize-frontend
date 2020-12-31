package main

import (
	"log"
	"os"
)

func main() {
	//Test write permission . it is possible thr file
	//Does not exist and that will return a different
	//Error that can be checked with os.IsPermission(err)

	file, err := os.OpenFile("test.txt", os.O_WRONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("Error : Write permission denied.")
		}
	}
	file.Close()

	//Test read permissions

	file1, err1 := os.OpenFile("test.txt", os.O_RDONLY, 0666)
	if err1 != nil {
		if os.IsPermission(err1) {
			log.Println("Error: Read permission denied.")
		}
	}
	file1.Close()
}
