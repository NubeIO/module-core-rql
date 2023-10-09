package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
	"github.com/NubeIO/rubix-os/interfaces"
	"github.com/NubeIO/rubix-os/nresty"
)

func (inst *Client) EdgeSystemCtlAction(hostIDName, serviceName string, action interfaces.Action) (*interfaces.Message, error) {
	url := fmt.Sprintf("/proxy/eb/api/systemctl/%s?unit=%s", action, serviceName)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&interfaces.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*interfaces.Message), nil
}

func (inst *Client) EdgeSystemCtlState(hostIDName, serviceName string) (*systemctl.SystemState, error) {
	url := fmt.Sprintf("/proxy/eb/api/systemctl/state?unit=%s", serviceName)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&systemctl.SystemState{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*systemctl.SystemState), nil
}
