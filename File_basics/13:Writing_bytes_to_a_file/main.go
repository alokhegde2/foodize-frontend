//You can write using just the os package , which is needed already to open the file

package main

import (
	"log"
	"os"
)

func main() {
	//open new file for writing only
	file, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//write bytes to file
	byteSlice := []byte("Bytes!\n")
	bytesWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes . \n", bytesWritten)
}
