package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/rubix-os/services/system"

	systats "github.com/NubeIO/lib-system"
	"github.com/NubeIO/rubix-os/nresty"
)

func (inst *Client) EdgeHostReboot(hostIDName string) (*system.Message, error) {
	url := fmt.Sprintf("/proxy/ros/api/system/reboot/")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&system.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*system.Message), nil
}

func (inst *Client) EdgeGetSystem(hostIDName string) (*systats.System, error) {
	url := fmt.Sprintf("/proxy/ros/api/system/info/")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&systats.System{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*systats.System), nil
}

func (inst *Client) EdgeGetMemoryUsage(hostIDName string) (*system.MemoryUsage, error) {
	url := fmt.Sprintf("/proxy/ros/api/system/usage/")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&system.MemoryUsage{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*system.MemoryUsage), nil
}

func (inst *Client) EdgeGetMemory(hostIDName string) (*systats.Memory, error) {
	url := fmt.Sprintf("/proxy/ros/api/system/memory/")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&systats.Memory{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*systats.Memory), nil
}

func (inst *Client) GetTopProcesses(hostIDName, sort string, count int) (*[]systats.Process, error) {
	url := fmt.Sprintf("/proxy/ros/api/system/processes?sort=%s&count=%d", sort, count)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&[]systats.Process{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*[]systats.Process), nil
}

func (inst *Client) GetSwap(hostIDName string) (*systats.Swap, error) {
	url := fmt.Sprintf("/proxy/ros/api/system/swap/")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&systats.Swap{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*systats.Swap), nil
}

func (inst *Client) EdgeDiscUsage(hostIDName string) (*[]systats.Disk, error, error) {
	url := fmt.Sprintf("/proxy/ros/api/system/disc/")
	resp, connectionErr, requestErr := nresty.FormatRestyV2Response(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&[]systats.Disk{}).
		Get(url))
	if connectionErr != nil || requestErr != nil {
		return nil, connectionErr, requestErr
	}
	data := resp.Result().(*[]systats.Disk)
	return data, nil, nil
}

func (inst *Client) DiscUsagePretty(hostIDName string) (*[]system.Disk, error) {
	url := fmt.Sprintf("/proxy/ros/api/system/disc/pretty/")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&[]system.Disk{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*[]system.Disk), nil
}
