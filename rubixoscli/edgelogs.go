package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	"github.com/NubeIO/rubix-os/nresty"
)

// EdgeCreateLog will make, get and delete a log all in the one api
func (inst *Client) EdgeCreateLog(hostIDName string, body *model.StreamLog) (*model.StreamLog, error) {
	url := fmt.Sprintf("/proxy/ros/api/logs/create/")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&model.StreamLog{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.StreamLog), nil
}

func (inst *Client) EdgeGetLogs(hostIDName string) ([]model.StreamLog, error) {
	url := fmt.Sprintf("/proxy/ros/api/logs/")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&[]model.StreamLog{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return *resp.Result().(*[]model.StreamLog), nil
}
