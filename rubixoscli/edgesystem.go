package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/rubix-os/services/system"

	systats "github.com/NubeIO/lib-system"
	"github.com/NubeIO/rubix-os/nresty"
)

func (inst *Client) EdgeHostReboot(hostUUID string) (*system.Message, error) {
	url := fmt.Sprintf("/host/ros/api/system/reboot")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&system.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*system.Message), nil
}

func (inst *Client) EdgeGetSystem(hostUUID string) (*systats.System, error) {
	url := fmt.Sprintf("/host/ros/api/system/info")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&systats.System{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*systats.System), nil
}

func (inst *Client) EdgeGetMemoryUsage(hostUUID string) (*system.MemoryUsage, error) {
	url := fmt.Sprintf("/host/ros/api/system/usage")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&system.MemoryUsage{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*system.MemoryUsage), nil
}

func (inst *Client) EdgeGetMemory(hostUUID string) (*systats.Memory, error) {
	url := fmt.Sprintf("/host/ros/api/system/memory")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&systats.Memory{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*systats.Memory), nil
}

func (inst *Client) GetTopProcesses(hostUUID, sort string, count int) (*[]systats.Process, error) {
	url := fmt.Sprintf("/host/ros/api/system/processes?sort=%s&count=%d", sort, count)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&[]systats.Process{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*[]systats.Process), nil
}

func (inst *Client) GetSwap(hostUUID string) (*systats.Swap, error) {
	url := fmt.Sprintf("/host/ros/api/system/swap")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&systats.Swap{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*systats.Swap), nil
}

func (inst *Client) EdgeDiscUsage(hostUUID string) (*[]systats.Disk, error, error) {
	url := fmt.Sprintf("/host/ros/api/system/disc")
	resp, connectionErr, requestErr := nresty.FormatRestyV2Response(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&[]systats.Disk{}).
		Get(url))
	if connectionErr != nil || requestErr != nil {
		return nil, connectionErr, requestErr
	}
	data := resp.Result().(*[]systats.Disk)
	return data, nil, nil
}

func (inst *Client) DiscUsagePretty(hostUUID string) (*[]system.Disk, error) {
	url := fmt.Sprintf("/host/ros/api/system/disc/pretty")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&[]system.Disk{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*[]system.Disk), nil
}
