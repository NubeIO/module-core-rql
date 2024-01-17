package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/dto"

	"github.com/NubeIO/rubix-os/nresty"
)

func (inst *Client) UpdateRestartJobConfig(hostUUID string, body *dto.RestartJob) (*dto.RestartJob, error) {
	url := "/host/ros/api/restart-jobs"
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetBody(body).
		SetResult(&dto.RestartJob{}).
		Put(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*dto.RestartJob)
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
