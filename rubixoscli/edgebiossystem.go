package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/rubix-os/interfaces"
	"github.com/NubeIO/rubix-os/nresty"
	"github.com/NubeIO/rubix-ui/backend/rumodel"
)

func (inst *Client) EdgeBiosPing(hostIDName string) (*interfaces.Message, error) {
	url := fmt.Sprintf("/proxy/eb/api/system/ping")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&interfaces.Message{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*interfaces.Message), nil
}

func (inst *Client) EdgeBiosArch(hostIDName string) (*rumodel.Arch, error) {
	url := fmt.Sprintf("/proxy/eb/api/system/arch")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&rumodel.Arch{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*rumodel.Arch), nil
}
