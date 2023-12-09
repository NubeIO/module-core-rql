package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/lib-utils-go/pprint"
	"github.com/NubeIO/rubix-os/services/system"
	"testing"
)

func TestClient_EdgeGetNetworks(t *testing.T) {
	data, err := client.EdgeGetNetworks("rc")
	fmt.Println(err)
	pprint.PrintJSON(data)
}

func TestClient_EdgeDHCPPortExists(t *testing.T) {
	data, err := client.EdgeDHCPPortExists("rc", &system.NetworkingBody{PortName: "eth0"})
	fmt.Println(err)
	fmt.Println(data)
}
