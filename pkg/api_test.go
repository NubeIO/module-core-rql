package pkg

import (
	"fmt"
	"testing"
)

func Test_getPathUUID(t *testing.T) {
	urlPath, uuid, combined := getPathUUID("/rules/abc")
	fmt.Println(urlPath, uuid, combined)
}
