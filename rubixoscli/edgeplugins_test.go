package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/rubix-os/interfaces"
	"github.com/NubeIO/rubix-os/utils/pprint"
	"testing"
)

func TestClient_EdgeUploadPlugin(t *testing.T) {
	data, err := client.EdgeUploadPlugin("rc", &interfaces.Plugin{
		Name:    "bacnetserver",
		Arch:    "amd64",
		Version: "v0.6.6",
	})
	fmt.Println(err)
	pprint.PrintJSON(data)
}

func TestClient_DeleteDownloadPlugins(t *testing.T) {
	data, connectionErr, requestErr := client.EdgeDeleteDownloadPlugins("rc")
	fmt.Println(data)
	fmt.Println(connectionErr)
	fmt.Println(requestErr)
}
