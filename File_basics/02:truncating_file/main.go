// when log file become too big , they can be truncated to save the disk space
//if we are attacking we may wnt to truncate .bash_history and other log files to cover your tracks
//Genuinly , malicious actors may truncate files just for the sake of destroying data

package main

import (
	"log"
	"os"
)

func main() {
	//Truncate a file to 100 bytes . if file is less than 100 bytes the original content will remains same
	//At the beginning ,  an d rest of the space is fill null bytes . If it is over 100 bytes
	//Everything past 100 bytes will be lost .
	//Either way
	//we will end up with exactly 100 bytes .
	//Pass in 0 in to truncate to completely empty files

	err := os.Truncate("test.txt", 100)
	if err != nil {
		log.Fatal(err)
	}
}
