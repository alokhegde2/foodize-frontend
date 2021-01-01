//This example uses zip but standard library
//also supports tar archives

package main

import (
	"archive/zip"
	"log"
	"os"
)

func main() {
	//Create a file to write the archive buffer todo
	//Could also use an in memory
	outFile, err := os.Create("test.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	//Create a zip writer on top of the file writter
	zipWriter := zip.NewWriter(outFile)

	//Add files to archives
	//We use some hard coded through all the files
	//in a directory add pass the name and contents
	//of each file,or you can take data from your
	//progeam and write it write in to the archive without

	var fileToArchive = []struct {
		Name, Body string
	}{
		{"test.txt", "String content of the file"},
		{"test2.txt", "\x61\x62\x63\n"},
	}

	//Create and write files to the archive which in turn
	//are getting written to the underlying writer to fileToArchive
	//.zip file we created at the beginning

	for _, file := range fileToArchive {
		fileWritter, err := zipWriter.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = fileWritter.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	//clean up
	err = zipWriter.Close()
	if err != nil {
		log.Fatal(err)
	}
}
