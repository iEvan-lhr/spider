package tool

import (
	"log"
	"strings"
)

func CutStringByName(s string) string {
	is := strings.Split(s, "/")
	if len(is) > 0 {
		return is[len(is)-1]
	} else {
		return s
	}
}

func ErrorExit(err error) {
	if err != nil {
		panic(err)
	}
}

func ErrorDontExit(err error) {
	if err != nil {
		log.Panicln(err)
	}
}

func GetOnlyInt(s string) int {
	i := 0
	for k := range s {
		i += int(s[k])
	}
	return i
}

func CheckIsImg(str string) bool {
	if strings.Index(str, "jpg") != -1 {
		return true
	}
	if strings.Index(str, "png") != -1 {
		return true
	}
	if strings.Index(str, "jpeg") != -1 {
		return true
	}
	if strings.Index(str, "gif") != -1 {
		return true
	}

	return false
}
