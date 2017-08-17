package main

import (
	"flag"
	"./du"
	"fmt"
	"time"
)

var verbose = flag.Bool("v",false,"show verbose progress messages")

func main() {
	flag.Parse()
	roots := flag.Args()
	if(len(roots)==0){
		roots = []string{"C:\\Program Files (x86)","C:\\Users\\"}
	}

	fileSizes := make(chan int64)

	go func() {
		for _,root := range roots{
			du.WalkDir(root,fileSizes)
		}
		close(fileSizes)
	}()

	var tick <-chan time.Time
	if *verbose{
		tick = time.Tick(500*time.Millisecond)
	}
	var nfiles,nbytes int64

	loop:
	for{
		select {
		case size,ok := <-fileSizes :
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles,nbytes)
		}
	}

	printDiskUsage(nfiles,nbytes)
}

func printDiskUsage(nfiles ,nbytes int64)  {
	fmt.Printf("%d files %d B\n",nfiles,nbytes)
}
