package rubixoscli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestClient_ProxyPOST(t *testing.T) {
	plan, _ := ioutil.ReadFile("proxy-example-io16.json")
	var data interface{}
	err := json.Unmarshal(plan, &data)
	fmt.Println(err)

	post, err := client.ProxyHostRosPOST("rc", "/api/system/scanner", data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(post)
}
