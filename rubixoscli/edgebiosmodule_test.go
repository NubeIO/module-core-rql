package rubixoscli

import (
	"fmt"
	pprint "github.com/NubeIO/rubix-ui/backend/helpers/print"
	"testing"
)

func TestClient_EdgeListModules(t *testing.T) {
	modules, connectionErr, requestErr := client.EdgeListModules("rc")
	pprint.PrintJOSN(modules)
	fmt.Println(connectionErr)
	fmt.Println(requestErr)
}
