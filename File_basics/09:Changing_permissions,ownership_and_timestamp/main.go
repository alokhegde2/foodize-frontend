//If we have privileges to change the ownership,
//A set of function are provided by the standard librery , they are
//os.Chmod()
//os.Chown()
//os.Chtime()

package main

import (
	"log"
	"os"
	"time"
)

func main() {
	//Change the permissions using Linuc style
	err := os.Chmod("test.txt", 0777)
	if err != nil {
		log.Println(err)
	}

	//change ownership
	err1 := os.Chown("test.txt", os.Getuid(), os.Getgid())

	if err1 != nil {
		log.Println(err1)
	}

	//Change timestamps

	twoDaysFromNow := time.Now().Add(48 * time.Hour)
	lastModifyTime := twoDaysFromNow
	lastAccessTime := twoDaysFromNow

	err2 := os.Chtimes("test.txt", lastAccessTime, lastModifyTime)
	if err2 != nil {
		log.Println(err2)
	}
}
