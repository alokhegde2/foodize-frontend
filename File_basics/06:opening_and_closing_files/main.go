package main

import (
	"log"
	"os"
)

func main() {
	//Simple read only open.We will cover actually reading
	//And writing to file in examples further down package
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	//Open file with more options . last params is the permission mode
	//Second param is the is the attributes when opening

	file, err = os.OpenFile("test.txt", os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	//We have to close the file when our work is done
	file.Close()

	//use these attributes individually or combined
	//with an OR for second arg of OpenFile()
	//e.g. os.O_CREATE|os.O_APPEND
	//or os.O_CREATE|os.O_TRUNC|os.O_WRONLY

	//os.O_RDONLY //Read only
	//os.O_WRONLY //Write only
	//os.O_RDWR //Read and Write
	//os.O_APPEND //Append to end of OpenFile
	//os.O_CREATE //Create is none exist
	//os.O_TRUNC  // Truncate file when opening
}
