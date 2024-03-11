package util

import (
	"fmt"
	"strings"
)

func ArrayToString(a []int, delim string) (res string) {
	res = strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
	return
	//return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
	//return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}
