package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/rubix-os/src/cli/constants"
	"testing"
)

func TestApp_EdgeAppStatus(t *testing.T) {
	status, connectionErr, requestErr := client.EdgeAppStatus("rc", constants.RubixOs)
	fmt.Println("status", status)
	fmt.Println("connectionErr", connectionErr)
	fmt.Println("requestErr", requestErr)
}
