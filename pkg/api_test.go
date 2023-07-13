package pkg

import (
	"fmt"
	"testing"
)

func Test_getPath(t *testing.T) {
	url := "points/pnt_d2e2ced50da74b1a"
	//uuid, err := urlGetUUID(url)
	//fmt.Println(err)
	//fmt.Println(uuid)

	s := urlSplit(url)
	fmt.Println(len(s))
	lastItem := s[len(s)-1]
	fmt.Println(lastItem)
	lastItem = s[len(s)-2]
	fmt.Println(lastItem)

}
