package rubixoscli

import (
	"errors"
	"fmt"
	"github.com/NubeIO/rubix-os/nresty"
	"strings"
)

type FilesExists struct {
	File   string `json:"file"`
	Exists bool   `json:"exists"`
}

func (inst *Client) EdgeFileExists(hostUUID, path string) (*FilesExists, error) {
	url := fmt.Sprintf("/host/ros/api/files/exists/?file=%s", path)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&FilesExists{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*FilesExists), nil
}

func (inst *Client) EdgeDeleteDataFile(hostUUID, path string) (*Message, error) {
	if path == "/" {
		return nil, errors.New("the root dir can not be deleted")
	}
	if path == "/data" {
		return nil, errors.New("the /data dir can not be deleted")
	}
	if !strings.Contains(path, "/data") {
		return nil, errors.New(fmt.Sprintf("path %s must be in the /data dir", path))
	}
	url := fmt.Sprintf("/host/ros/api/files/delete/?file=%s", path)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&Message{}).
		Delete(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*Message), nil
}
