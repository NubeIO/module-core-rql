package rubixoscli

import (
	"fmt"

	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	"github.com/NubeIO/rubix-os/nresty"
)

func (inst *Client) GetPointHistories(hostIDName string, pointUUIDs []string) ([]model.PointHistory, error) {
	url := fmt.Sprintf("/proxy/ros/api/histories/points/point_uuids")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&[]model.PointHistory{}).
		SetBody(pointUUIDs).
		Post(url))
	if err != nil {
		return nil, err
	}
	var out []model.PointHistory
	out = *resp.Result().(*[]model.PointHistory)
	return out, nil
}

func (inst *Client) GetPointHistoriesWithInterval(hostIDName, lowerBound, upperBound string, pointUUIDs []string) ([]model.PointHistory, error) {
	url := fmt.Sprintf("/proxy/ros/api/histories/points/point_uuids?timestamp_gt=%s&&timestamp_lt=%s", lowerBound, upperBound)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&[]model.PointHistory{}).
		SetBody(pointUUIDs).
		Post(url))
	if err != nil {
		return nil, err
	}
	var out []model.PointHistory
	out = *resp.Result().(*[]model.PointHistory)
	return out, nil
}
