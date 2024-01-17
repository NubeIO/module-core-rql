package rubixoscli

import (
	"github.com/NubeIO/nubeio-rubix-lib-models-go/dto"
	"github.com/NubeIO/rubix-os/nresty"
)

func (inst *Client) SyncEdgeDevices(hostUUID string) (*dto.Message, error) {
	url := "/api/host/clone-things-to-cloud"
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&dto.Message{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*dto.Message), nil
}
