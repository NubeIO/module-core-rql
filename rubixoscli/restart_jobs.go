package rubixoscli

import (
	"fmt"

	"github.com/NubeIO/rubix-os/interfaces"
	"github.com/NubeIO/rubix-os/nresty"
)

func (inst *Client) UpdateRestartJobConfig(hostUUID string, body *interfaces.RestartJob) (*interfaces.RestartJob, error) {
	url := "/host/ros/api/restart-jobs"
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetBody(body).
		SetResult(&interfaces.RestartJob{}).
		Put(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*interfaces.RestartJob)
	return out, nil
}

func (inst *Client) DeleteRestartJobConfig(hostUUID string, unit string) (bool, error) {
	url := fmt.Sprintf("/host/ros/api/restart-jobs/unit/%s", unit)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		Delete(url))
	if err != nil {
		return false, err
	}
	return resp.String() == "true", nil
}
