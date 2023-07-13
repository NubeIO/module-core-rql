package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/rubix-os/interfaces"
	"github.com/NubeIO/rubix-os/nresty"
)

func (inst *Client) ConfigureOpenVPN(uuid string) (*interfaces.Message, error) {
	url := fmt.Sprintf("/api/hosts/%s/configure-openvpn", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", uuid).
		SetResult(&interfaces.Message{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*interfaces.Message), nil
}
