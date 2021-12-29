package Persistence

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/dlclark/regexp2"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"spider/tool"
	"strings"
	"sync"
	"time"
)

func (w Work) PersistenceUrl(msg, needUrl string) []string {
	if msg != "" {
		doc, err := htmlquery.Parse(strings.NewReader(msg))
		tool.ErrorExit(err)
		find := htmlquery.Find(doc, "//a")
		img := htmlquery.Find(doc, "//img")
		w.FindAllIMG(img)
		m := make([]string, 0)
		for i := range find {
			if val := htmlquery.SelectAttr(find[i], "href"); strings.Index(val, needUrl) != -1 {
				if strings.Index(val, "http") == -1 {
					m = append(m, "https:"+val)
				} else {
					m = append(m, val)
				}
			}
		}
		return m
	}
	return nil
}
func FindAllString(re *regexp2.Regexp, s string) []string {
	var matches []string
	m, _ := re.FindStringMatch(s)
	for m != nil {
		matches = append(matches, m.String())
		m, _ = re.FindNextMatch(m)
	}
	return matches
}

type Work struct {
}

var syt sync.Mutex

func saveIMG(url string) {
	resp, err := http.Get(url)
	defer func() {
		resp.Body.Close()
	}()
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	name := tool.CutStringByName(url)
	if _, ok := imgMap[tool.GetOnlyInt(name)]; !ok {
		if len(imgMap) == 0 {
			imgMap = map[int]string{tool.GetOnlyInt(name): "ok"}
		} else {
			imgMap[tool.GetOnlyInt(name)] = "ok"
		}
		_ = ioutil.WriteFile("downloadRES/"+name, body, 0755)
	}
}

func (w Work) FindAllIMG(urls []*html.Node) {
	for i := range urls {
		syt.Lock()
		if val := htmlquery.SelectAttr(urls[i], "src"); tool.CheckIsImg(val) {
			if strings.Index(val, "http") == -1 {
				saveIMG("https:" + val)
			} else {
				saveIMG(val)
			}
		}
		syt.Unlock()
		time.Sleep(150 * time.Millisecond)
	}
}

var imgMap map[int]string
