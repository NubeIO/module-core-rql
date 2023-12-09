package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/rubix-os/interfaces"
	"github.com/NubeIO/rubix-os/nresty"
	"github.com/NubeIO/rubix-os/src/cli/bioscli/bmodel"
)

func (inst *Client) EdgeBiosRubixOsOnEdgeUpload(hostUUID string, upload interfaces.FileUpload) (*interfaces.Message, error) {
	url := fmt.Sprintf("/api/host/ros/upload")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetBody(upload).
		SetResult(&interfaces.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*interfaces.Message), nil
}

func (inst *Client) EdgeBiosRubixOsOnEdgeInstall(hostUUID string, upload interfaces.FileUpload) (*interfaces.Message, error) {
	url := fmt.Sprintf("/api/host/ros/install")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetBody(upload).
		SetResult(&interfaces.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*interfaces.Message), nil
}

func (inst *Client) EdgeBiosRubixOsOnEdgeVersion(hostUUID string) (*bmodel.Version, error) {
	url := fmt.Sprintf("/api/host/ros/version")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&bmodel.Version{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*bmodel.Version), nil
}
