//The following program will print out all the meatadata available about a file
//It includes file name,size,permissions,lastmodified time and weather it is a directory

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//Stat returns file info .
	//it will returns an error if there is no file.

	fileInfo, err := os.Stat("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File Name : ", fileInfo.Name())
	fmt.Println("File Size : ", fileInfo.Size())
	fmt.Println("Permissions : ", fileInfo.Mode())
	fmt.Println("Last Modified : ", fileInfo.ModTime())
	fmt.Println("Is Directory : ", fileInfo.IsDir())
	fmt.Printf("File Interface type : %T\n", fileInfo.Sys())
	fmt.Printf("File Info : %+v\n\n", fileInfo.Sys())
}
