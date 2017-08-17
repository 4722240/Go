package main

import (
	//"bytes"
	"fmt"
	//"io"
	//"io/ioutil"
	//"net/http"
	//"os"
	//"strings"
	"math/cmplx"
	"time"
)

//func getImg(url string) (n int64, err error) {


	//path := strings.Split(url, "/")
	//var name string
	//if len(path) > 1 {
	//	name = path[len(path)-1]
	//}
	//fmt.Println(name)
	//out, err := os.Create(name)
	//defer out.Close()
	//resp, err := http.Get(url)
	//defer resp.Body.Close()
	//pix, err := ioutil.ReadAll(resp.Body)
	//n, err = io.Copy(out, bytes.NewReader(pix))
	//return

//}

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)
func main() {
	//fmt.Println("1111"=="1111")
	//getImg("http://img2.bdstatic.com/img/image/166314e251f95cad1c8f496ad547d3e6709c93d5197.jpg")
	//const f = "%T(%v)\n"
	//fmt.Printf(f, ToBe, ToBe)
	//fmt.Printf(f, MaxInt, MaxInt)
	//fmt.Printf(f, z, z)
	fmt.Println(time.Now().UTC())
}