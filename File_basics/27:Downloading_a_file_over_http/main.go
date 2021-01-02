//A comman tsk in modern computing is downloading a file over http protocol.
//The following example shows how to quickly download a specific url to a file

package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	//Create a output file

	newFile, err := os.Create("devdungeon.html")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	//HTTP GET request devdungeon.com

	url := "http://www.devdungeon.com/archives"
	// url := "https://meet.google.com"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	//Write bytes from HTTP response to file.
	//response.Body satisfies the reader interface
	//newFile satisfies the writer interface
	//That allows us to use io.Copy which accepts
	//any type that implements reader and writer interface

	numBytesWritten, err := io.Copy(newFile, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Downloaded %d bytes file.\n", numBytesWritten)
}
