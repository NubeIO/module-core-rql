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
	Port:      1660,
	HTTPS:     false,
}, &installer.Installer{})

func TestClient_GetAlerts(t *testing.T) {
	alerts, err := client.GetAlerts()
	if err != nil {
		return
	}
	pprint.Print(alerts)
}
