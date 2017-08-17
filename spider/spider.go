package main

import (
	"net/http"
	"log"
	"fmt"
	"io/ioutil"
	"regexp"
	//"strings"
	"os"
	"io"
	"bytes"
	"sync"
)



func main() {

	dir := "E:\\P\\test\\"

	url := "https://exhentai.org/s/8713d5c4d5/482563-182"

	var wg  = sync.WaitGroup{}
	for i:=0;;i++ {
		html,err := connect(url)
		if err != nil {
			fmt.Println(err)
		}

		reg := regexp.MustCompile(`id="i3"><a onclick="return load_image(.+?)" href="(.+?)"><img id="img" src="(.+?)"`)
		str := reg.FindAllSubmatch(html,-1)

		//fmt.Printf("%q\n",str)

		//fmt.Printf("%s\n",str[0][3])

		name := fmt.Sprintf("%04d.jpg",i)

		wg.Add(1)
		go getImg(string(str[0][3]),dir,name,&wg)

		fmt.Printf("next url is %s \n",string(str[0][2]))
		if string(str[0][2])== url{
			break
		}else {
			url = string(str[0][2])
		}
	}

	wg.Wait()


}

func connect(url string) ([]byte,error){

	client := &http.Client{}

	req ,err := http.NewRequest("GET",url,nil)

	//client.CheckRedirect = CheckRedirect;
	//cookie := http.Cookie{}
	//req.AddCookie(cookie)
	req.Header.Set("Cookie","igneous=791d8b38d; ipb_member_id=897616; ipb_pass_hash=a3bfabedfa06e3c2898340495e3cab26; lv=1479975677-1479975745; s=2fa97f5c7; yay=louder")


	if(err != nil){
		return make([]byte,0),fmt.Errorf("create client error")
	}

	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return make([]byte,0),fmt.Errorf("connect %s error: %s",url,err)
	}
	if resp.StatusCode != 200 {
		return make([]byte,0),fmt.Errorf("getting %s StatusCode: %d",url,resp.StatusCode)
	}
	body ,err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	return body,nil
}

func getImg(url, dir, name string,wg *sync.WaitGroup) (n int64, err error) {
	defer wg.Done()

	name = dir +name

	out, err := os.Create(name)
	defer out.Close()
	pix, err := connect(url)

	fmt.Printf("%s 已经下载\n",name)
	n, err = io.Copy(out, bytes.NewReader(pix))
	return
}

//func CheckRedirect (req *http.Request, via []*http.Request) error{
//	if len(via) >= 0 {
//		str := fmt.Sprintf("stopped after %d redirects",len(via))
//		return errors.New(str)
//	}
//	return nil
//}