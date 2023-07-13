package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/rubix-os/interfaces"
	"github.com/NubeIO/rubix-os/nresty"
)

// EdgeCreateLog will make, get and delete a log all in the one api
func (inst *Client) EdgeCreateLog(hostIDName string, body *interfaces.StreamLog) (*interfaces.StreamLog, error) {
	url := fmt.Sprintf("/proxy/ros/api/logs/create/")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&interfaces.StreamLog{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*interfaces.StreamLog), nil
}

func (inst *Client) EdgeGetLogs(hostIDName string) ([]interfaces.StreamLog, error) {
	url := fmt.Sprintf("/proxy/ros/api/logs/")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&[]interfaces.StreamLog{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return *resp.Result().(*[]interfaces.StreamLog), nil
}
