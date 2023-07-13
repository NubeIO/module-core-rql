package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/rubix-os/interfaces"
	"github.com/NubeIO/rubix-os/nresty"
)

func (inst *Client) FFSystemPing(hostIDName string) (*interfaces.Message, error) {
	url := fmt.Sprintf("/proxy/ros/api/system/ping")
	_, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		Get(url))
	if err != nil {
		return nil, err
	}
	return &interfaces.Message{Message: "ping success"}, nil
}
