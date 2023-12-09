package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/rubix-os/services/system"
	"time"

	"github.com/NubeIO/lib-date/datelib"
	"github.com/NubeIO/rubix-os/nresty"
)

func (inst *Client) EdgeSystemTime(hostUUID string) (*datelib.Time, error) {
	url := fmt.Sprintf("/host/ros/api/time")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&datelib.Time{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*datelib.Time), nil
}

func (inst *Client) EdgeGetHardwareTZ(hostUUID string) (string, error) {
	url := fmt.Sprintf("/host/ros/api/timezone")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		Get(url))
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}

func (inst *Client) EdgeGetTimeZoneList(hostUUID string) ([]string, error) {
	url := fmt.Sprintf("/host/ros/api/timezone/list")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&[]string{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return *resp.Result().(*[]string), nil
}

func (inst *Client) EdgeUpdateTimezone(hostUUID string, timeZone string) (*system.Message, error) {
	url := fmt.Sprintf("/host/ros/api/timezone")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetBody(system.DateBody{TimeZone: timeZone}).
		SetResult(&system.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*system.Message), nil
}

func (inst *Client) EdgeUpdateSystemTime(hostUUID, timeString string) (*datelib.Time, error) {
	url := fmt.Sprintf("/host/ros/api/time")
	layout := "2006-01-02 15:04:05"
	// parse time
	_, err := time.Parse(layout, timeString)
	if err != nil {
		return nil, fmt.Errorf("could not parse date try 2006-01-02 15:04:05 %s", err)
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetBody(system.DateBody{DateTime: timeString}).
		SetResult(&datelib.Time{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*datelib.Time), nil
}

func (inst *Client) EdgeNTPEnable(hostUUID string) (*system.Message, error) {
	url := fmt.Sprintf("/host/ros/api/time/ntp/enable")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&system.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*system.Message), nil
}

func (inst *Client) EdgeNTPDisable(hostUUID string) (*system.Message, error) {
	url := fmt.Sprintf("/host/ros/api/time/ntp/disable")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&system.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*system.Message), nil
}
