package main

import (
	"flag"
	"./du"
	"fmt"
)
func main() {
	flag.Parse()
	roots := flag.Args()
	if(len(roots)==0){
		roots = []string{"C:\\Go\\","C:\\Program Files (x86)"}
	}

	fileSizes := make(chan int64)

	go func() {
		for _,root := range roots{
			du.WalkDir(root,fileSizes)
		}
		close(fileSizes)
	}()

	var nfiles,nbytes int64
	for size := range fileSizes{
		nfiles++
		nbytes += size
	}
	printDiskUsage(nfiles,nbytes)
}

func printDiskUsage(nfiles ,nbytes int64)  {
	fmt.Printf("%d files %d B\n",nfiles,nbytes)
}
