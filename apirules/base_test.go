package apirules

import (
	"fmt"
	"testing"
)

func Test_getToken(t *testing.T) {
	got := getToken()
	fmt.Println(got)
}
