package main

import (
	"log"
	"os"
)

func main() {
	//Stat returns file info. It will returns
	//an error if there is  no file.
	fileInfo, err := os.Stat("test.txt")
	if err != nil {
		if os.IsExist(err) {
			log.Fatal("File does not exist.")
		}
	}
	log.Println("File does Exist . File information :")
	log.Println(fileInfo)
}
