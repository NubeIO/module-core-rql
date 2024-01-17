package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/dto"
	"github.com/NubeIO/rubix-os/nresty"
	"github.com/NubeIO/rubix-os/src/cli/bioscli/bmodel"
)

func (inst *Client) EdgeBiosRubixOsOnEdgeUpload(hostUUID string, upload dto.FileUpload) (*dto.Message, error) {
	url := fmt.Sprintf("/api/host/ros/upload")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetBody(upload).
		SetResult(&dto.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*dto.Message), nil
}

func (inst *Client) EdgeBiosRubixOsOnEdgeInstall(hostUUID string, upload dto.FileUpload) (*dto.Message, error) {
	url := fmt.Sprintf("/api/host/ros/install")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetBody(upload).
		SetResult(&dto.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*dto.Message), nil
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
