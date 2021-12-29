package main

import "spider/engine"

func main() {
	r := engine.Req{
		Url:     "https://699pic.com/zhuanti/kejishenghuo.html",
		NeedUrl: "699pic.com",
	}
	r.Start()
}
