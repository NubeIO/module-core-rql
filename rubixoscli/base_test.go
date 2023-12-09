package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/rubix-os/installer"
	"github.com/NubeIO/rubix-os/rubixregistry"
	"github.com/go-resty/resty/v2"
	"testing"
)

// TODO: WARNING: DON'T OVERRIDE FROM COPYING
var installerObj = installer.New(&installer.Installer{}, rubixregistry.New("/data"))
var client = New(&Client{
	Rest:      &resty.Client{},
	Installer: installerObj,
	Ip:        "0.0.0.0",
	Port:      1660,
	HTTPS:     false,
}, installerObj)

func TestClient_Alerts(t *testing.T) {
	alerts, err := client.GetAlerts([]string{"active"})
	fmt.Println("alert", alerts)
	fmt.Println("error", err)
}
