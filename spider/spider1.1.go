package main

import (
	"net/http"
	"log"
	"fmt"
	"io/ioutil"
	"regexp"
	"os"
	"io"
	"bytes"
	"sync"
)

type downloadUrl struct{
	url string
	name string
	dir string
	end int//0为结束
}

func main() {

	dir := "E:\\P\\test\\"
	url := "https://exhentai.org/s/8713d5c4d5/482563-182"
	thread := 3

	var wg  = sync.WaitGroup{}

	urlch := make(chan downloadUrl,10)

	for i:=0;i<thread; i++{
		wg.Add(1)
		go getImg(urlch,&wg,i)
	}


	for i:=1;;i++ {
		html,err := connect(url)
		fmt.Println(len(html))
		if err != nil {
			fmt.Println(err)
			i--
			continue
		}

		reg := regexp.MustCompile(`id="i3"><a onclick="return load_image(.+?)" href="(.+?)"><img id="img" src="(.+?)"`)
		str := reg.FindAllSubmatch(html,-1)

		//fmt.Printf("%q\n",str)

		name := fmt.Sprintf("%04d.jpg",i)

		urlch <- downloadUrl{string(str[0][3]),name,dir,1}


		fmt.Printf("next url is %s \n",string(str[0][2]))

		if string(str[0][2])== url{
			fmt.Println("url collected end")
			break
		}else {
			url = string(str[0][2])
		}
	}
	close(urlch)

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

	if err != nil {
		return make([]byte,0),fmt.Errorf("connect %s error: %s",url,err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return make([]byte,0),fmt.Errorf("getting %s : %s",url,resp.StatusCode)
	}
	body ,err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	return body,nil
}

func getImg(downUrl <-chan downloadUrl,wg *sync.WaitGroup,No int) ( err error) {
	defer wg.Done()

	for  {
		urlStruct := <- downUrl
		fmt.Println(urlStruct)

		if urlStruct.end == 0{
			return nil
		}

		name := urlStruct.dir + urlStruct.name

		out, err := os.Create(name)
		defer out.Close()
		pix, err := connect(urlStruct.url)

		fmt.Printf("%02d 线程已经下载 %s \n",No,urlStruct.name)
		_,err = io.Copy(out, bytes.NewReader(pix))

		if err != nil {
			return err
		}
	}
}
