package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/dto"
	"github.com/NubeIO/rubix-os/nresty"
)

func (inst *Client) ConfigureOpenVPN(hostUUID string) (*dto.Message, error) {
	url := fmt.Sprintf("/api/host/configure-openvpn")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&dto.Message{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*dto.Message), nil
}
