package main

import (
	"fmt"
	"io/ioutil"
)

func PrintTabDepth(x int) {
	for i := 0; i < x; i++ {
		fmt.Print("  ")
	}
}

func CrawlDir(dir string, dep int) {
	// Get files info in current directory
	fs, err := ioutil.ReadDir(dir)
	if err != nil {
		panic("Could not read directory")
	}

	for _, file := range fs {
		PrintTabDepth(dep)
		fmt.Printf("%s", file.Name())
		if file.IsDir() {
			fmt.Printf(" DIR ")
		}
		fmt.Printf("\n")

		if file.IsDir() {
			CrawlDir(dir+"/"+file.Name(), dep+1)
		}
	}
}

func main() {
	bd := "."

	CrawlDir(bd, 0)

}
