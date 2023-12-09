package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
	"github.com/NubeIO/rubix-os/interfaces"
	"github.com/NubeIO/rubix-os/nresty"
)

func (inst *Client) EdgeSystemCtlAction(hostUUID, serviceName string, action interfaces.Action) (*interfaces.Message, error) {
	url := fmt.Sprintf("/host/bios/api/systemctl/%s?unit=%s", action, serviceName)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&interfaces.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*interfaces.Message), nil
}

func (inst *Client) EdgeSystemCtlState(hostUUID, serviceName string) (*systemctl.SystemState, error) {
	url := fmt.Sprintf("/host/bios/api/systemctl/state?unit=%s", serviceName)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&systemctl.SystemState{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*systemctl.SystemState), nil
}
