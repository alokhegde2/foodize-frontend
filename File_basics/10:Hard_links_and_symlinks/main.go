//A typical file is just a pointer to a place on the hard disk, called an inode.
//A hard link creates a new pointer to the same place
//A file will only be deleted from the disk after all links to it are removed
//A hard link is what you might consider a "normal" links

//A symbolic link , or a soft link is a little different,it does not point directly to a place on the disk
//Symlink only references other files by name .
//they can point to files on different flile systems
//All sysytem will not supports symlinks

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//create a hard link

	//You will have two file names that point to the same content
	//Changing the contents of one will change the other
	//Deleting/renaming one will not affect the other

	err := os.Link("original.txt", "original_also.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Creating symlink")
	//create a symlinks
	err = os.Symlink("original.txt", "original_sym.txt")
	if err != nil {
		log.Fatal(err)
	}

	//Lstat will return file info , but if it is actually
	//a symlink,it will not follow the link and give information
	//It will not follow the link and give information
	//about the real life
	//Symlinks donot  work in Windows

	fileInfo, err := os.Lstat("Original_sym.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Link info : %+v", fileInfo)

	//Change ownership of symlink only
	//and not the file it points to
	err = os.Lchown("original_sym.txt", os.Getuid(), os.Getgid())

	if err != nil {
		log.Fatal(err)
	}
}
