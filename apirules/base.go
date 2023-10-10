package apirules

import (
	"fmt"
	"github.com/NubeIO/module-core-rql/rubixoscli"
	"github.com/NubeIO/module-core-rql/storage"
	"github.com/NubeIO/rubix-os/installer"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

type RQL struct {
	Return    interface{}      `json:"return"`
	Err       string           `json:"err"`
	TimeTaken string           `json:"time_taken"`
	Storage   storage.IStorage `json:"-"`
}

func getToken() string {
	token, err := ioutil.ReadFile("/data/auth/internal_token.txt")
	if err != nil {
		log.Error(fmt.Sprintf("get ROS token err   #%v ", err))
		return ""
	}
	if len(string(token)) < 40 {
		log.Error(fmt.Sprintf("ROS token lenght is to short len: %d", len(string(token))))
		return ""
	}
	return string(token)

}

var cli = rubixoscli.New(&rubixoscli.Client{
	Rest:          nil,
	Installer:     nil,
	Ip:            "0.0.0.0",
	Port:          1659,
	HTTPS:         false,
	ExternalToken: getToken(),
}, &installer.Installer{})

func errorString(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
