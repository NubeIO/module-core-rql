package rubixoscli

import (
	"github.com/NubeIO/rubix-os/interfaces"
	"github.com/NubeIO/rubix-os/nresty"
)

func (inst *Client) SyncEdgeDevices(hostIDName string) (*interfaces.Message, error) {
	url := "/api/clone_edges"
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
