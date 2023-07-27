package apirules

import (
	"github.com/NubeIO/module-core-rql/rubixoscli"
	"github.com/NubeIO/module-core-rql/storage"
	"github.com/NubeIO/rubix-os/installer"
)

type RQL struct {
	Return    interface{}      `json:"return"`
	Err       string           `json:"err"`
	TimeTaken string           `json:"time_taken"`
	Storage   storage.IStorage `json:"-"`
}

var cli = rubixoscli.New(&rubixoscli.Client{
	Rest:          nil,
	Installer:     nil,
	Ip:            "0.0.0.0",
	Port:          1659,
	HTTPS:         false,
	ExternalToken: "",
}, &installer.Installer{})

func errorString(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
