package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/dto"
	"github.com/NubeIO/rubix-os/nresty"
	"github.com/NubeIO/rubix-ui/backend/rumodel"
)

func (inst *Client) EdgeBiosPing(hostUUID string) (*dto.Message, error) {
	url := fmt.Sprintf("/host/bios/api/system/ping")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&dto.Message{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*dto.Message), nil
}

func (inst *Client) EdgeBiosArch(hostUUID string) (*rumodel.Arch, error) {
	url := fmt.Sprintf("/host/bios/api/system/arch")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&rumodel.Arch{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*rumodel.Arch), nil
}
