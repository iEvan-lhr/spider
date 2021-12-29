package engine

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"runtime"
	"spider/Persistence"
	"spider/tool"
	"time"
)

type AcceptMODE int

const (
	admin   AcceptMODE = 0
	public  AcceptMODE = 1
	private AcceptMODE = 2
)

func Accept(mode AcceptMODE) {
	switch mode {
	case admin:
	case public:
	case private:

	}
}

func init() {
	HeaderPublic()
}

var Header http.Header

func HeaderPublic() {
	header := http.Header{}
	header.Set("Accept", "*/*")
	//header.Set("Accept-Encoding","gzip, deflate, br")
	header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	header.Set("Connection", "keep-alive")
	header.Set("sec-ch-ua", `"Google Chrome";v="95", "Chromium";v="95", ";Not A Brand";v="99"`)
	header.Set("sec-ch-ua-mobile", "?0")
	header.Set("sec-ch-ua-platform", "Windows")
	header.Set("Sec-Fetch-Dest", "script")
	header.Set("Sec-Fetch-Mode", "no-cors")
	header.Set("Sec-Fetch-Site", "same-site")
	header.Set("User-Agent", " Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36")
	Header = header
}
func HeaderPrivate(cookie string) http.Header {
	if cookie == "" {
		tool.ErrorExit(errors.New("Please User Public Header! "))
	}
	header := http.Header{}
	header.Set("Accept", "*/*")
	//header.Set("Accept-Encoding","gzip, deflate, br")
	header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	header.Set("Connection", "keep-alive")
	header.Set("sec-ch-ua", `"Google Chrome";v="95", "Chromium";v="95", ";Not A Brand";v="99"`)
	header.Set("sec-ch-ua-mobile", "?0")
	header.Set("sec-ch-ua-platform", "Windows")
	header.Set("Sec-Fetch-Dest", "script")
	header.Set("Sec-Fetch-Mode", "no-cors")
	header.Set("Sec-Fetch-Site", "same-site")
	header.Set("User-Agent", " Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36")
	header.Set("Cookie", cookie)
	return header
}
func GetHttpRequest(url string, t int) *http.Request {
	if t == 1 {
		url = "https" + url[4:]
	}
	request, err := http.NewRequest("GET", url, nil)
	tool.ErrorExit(err)
	request.Header = Header
	return request
}

type Req struct {
	Url     string
	NeedUrl string
	Val     map[string]interface{}
}

func (r Req) Start() {
	request := GetHttpRequest(r.Url, 0)
	client := http.Client{}
	resp, err := client.Do(request)
	tool.ErrorExit(err)
	defer resp.Body.Close()
	all, err := ioutil.ReadAll(resp.Body)
	tool.ErrorExit(err)
	w := Persistence.Work{}
	i := 0
	url := w.PersistenceUrl(string(all), r.NeedUrl)
	i += len(url)
	channels := make([]chan int, runtime.NumCPU())
	for k := range channels {
		channels[k] = make(chan int, runtime.NumCPU()*2)
		go func(i int, c chan int) {
			for {
				fmt.Println(<-c)
			}
		}(i, channels[k])
	}
	for j := 0; j < len(url); j++ {
		time.Sleep(200 * time.Millisecond)
		go func(h chan int, str string) {
			cli := http.Client{}
			do, err2 := cli.Do(GetHttpRequest(str, 0))
			tool.ErrorExit(err2)
			defer do.Body.Close()
			allString, err := ioutil.ReadAll(do.Body)
			tool.ErrorExit(err)
			urls := w.PersistenceUrl(string(allString), r.NeedUrl)
			i += len(urls)
			h <- i
		}(channels[rand.Intn(runtime.NumCPU())], url[j])
	}
	time.Sleep(5 * time.Second)
	fmt.Println(i)
}
