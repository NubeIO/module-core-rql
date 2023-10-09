package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/rubix-os/interfaces"
	"github.com/NubeIO/rubix-os/nresty"
	"github.com/NubeIO/rubix-os/src/cli/bioscli/bmodel"
)

func (inst *Client) EdgeBiosRubixOsOnEdgeUpload(hostIDName string, upload interfaces.FileUpload) (*interfaces.Message, error) {
	url := fmt.Sprintf("/api/eb/ros/upload")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetBody(upload).
		SetResult(&interfaces.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*interfaces.Message), nil
}

func (inst *Client) EdgeBiosRubixOsOnEdgeInstall(hostIDName string, upload interfaces.FileUpload) (*interfaces.Message, error) {
	url := fmt.Sprintf("/api/eb/ros/install")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetBody(upload).
		SetResult(&interfaces.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*interfaces.Message), nil
}

func (inst *Client) EdgeBiosRubixOsOnEdgeVersion(hostIDName string) (*bmodel.Version, error) {
	url := fmt.Sprintf("/api/eb/ros/version")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&bmodel.Version{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*bmodel.Version), nil
}
