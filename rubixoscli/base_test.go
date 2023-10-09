package rubixoscli

import (
	"github.com/NubeIO/rubix-os/installer"
	"github.com/NubeIO/rubix-os/utils/pprint"
	"testing"
)

var client = New(&Client{
	Rest:      nil,
	Installer: nil,
	Ip:        "0.0.0.0",
	Port:      1659,
	HTTPS:     false,
}, &installer.Installer{})

func TestClient_PingRubixOs(t *testing.T) {
	networking, _, err := client.GetNetworking()
	if err != nil {
		return
	}
	pprint.Print(networking)
}
