//The ioutil package provides two functions:
//TempDir() and TempFile()
//It is the callers duty to delete the temporary items when done
//The only benifit of these functions provide is the that you can pass it an empty string for the directory
//And it will automically create the item in the system

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//create a temp file dir in the system default temp folder

	tempDirPath, err := ioutil.TempDir("", "myTempDir")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Temp dir created :", tempDirPath)

	//Create a file in new temp directory
	tempFile, err := ioutil.TempFile(tempDirPath, "myTempFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Temp file created:", tempFile.Name())

	//... do something with temp file/dir ...

	//close file
	err = tempFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	//Delete the resources we created
	err = os.Remove(tempFile.Name())
	if err != nil {
		log.Fatal(err)
	}
	err = os.Remove(tempDirPath)
	if err != nil {
		log.Fatal(err)
	}
}
