package du

import (
	"os"
	"io/ioutil"
	"fmt"
	"path/filepath"
)

func WalkDir(dir string,fileSizes chan <- int64 ){
	for _,entry := range  dirents(dir){
		if entry.IsDir(){
			subdir := filepath.Join(dir,entry.Name())
			WalkDir(subdir,fileSizes)
		}else {
			fileSizes <- entry.Size()
		}
	}
}
func dirents(dir string) []os.FileInfo {
	entries,err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr,"du1: %v\n",err)
		return nil
	}
	return entries
}
