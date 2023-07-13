package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/rubix-os/utils/pprint"
	"testing"
)

func TestClient_EdgeGetPoints(t *testing.T) {
	resp, err := client.EdgeGetPoints("rc")
	fmt.Println(err)
	if err != nil {
		return
	}
	pprint.Print(resp)
}
