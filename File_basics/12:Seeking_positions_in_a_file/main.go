//The seek() is useful for setting the file cursor in a specific location.
//seek() takes two parameters.The first one is distace;
//it can be moved forward with positive integer and backword with negative integer
//for example , if seek(-1,2) was specified,it would set the file cursor one byte back from the end of the file
//seek(2,0) would seek the second byte from the beginning of file

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, _ := os.Open("test.txt")
	defer file.Close()

	//Ofset is how many bytes to move
	//offset can be positive or negative

	var offset int64 = 5

	//Whence is the point of reference for offset
	//o = beginning of file
	//1 = current position
	//2 = End of the file

	var whence int = 0
	newPosition, err := file.Seek(offset, whence)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Just moved to 5:", newPosition)

	//Go back 2 bytes from current position
	newPosition, err = file.Seek(-2, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Just moved back two :", newPosition)

	//Find the current position by getting The
	//return value from seek after moving 0 bytes

	currentPosition, err := file.Seek(0, 1)
	fmt.Println("Current position:", currentPosition)

	//Go to beginning of file
	newPosition, err = file.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Position after seeking 0,0:", newPosition)

}
