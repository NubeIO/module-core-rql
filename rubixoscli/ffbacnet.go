package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-os/nresty"
)

type WhoIsOpts struct {
	InterfacePort   string `json:"interface_port"`
	LocalDeviceIP   string `json:"local_device_ip"`
	LocalDevicePort int    `json:"local_device_port"`
	LocalDeviceId   int    `json:"local_device_id"`
	Low             int    `json:"low"`
	High            int    `json:"high"`
	GlobalBroadcast bool   `json:"global_broadcast"`
	NetworkNumber   uint16 `json:"network_number"`
	Timeout         int    `json:"timeout,omitempty"`
	ApiTimeout      int    `json:"api_timeout,omitempty"`
}

const bacnetMaster = "bacnetmaster"

type discoveredDevices struct {
	DiscoveredCount int            `json:"discovered_count"`
	AddedCount      int            `json:"added_count"`
	ErrorOnAddCount int            `json:"error_on_add_count"`
	TimeTaken       string         `json:"time_taken"`
	Devices         []model.Device `json:"devices"`
}

// BacnetMasterWhoIs do a whois on an existing network
func (inst *Client) BacnetMasterWhoIs(hostIDName string, body *WhoIsOpts) ([]model.Device, error) {
	url := "/proxy/ros/api/plugins/api/bacnetmaster/master/whois"
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetBody(body).
		SetResult(&discoveredDevices{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*discoveredDevices)
	if out != nil {
		return out.Devices, nil
	}
	return nil, nil
}

type discoveredPoints struct {
	ObjectListLength int           `json:"object_list_length"`
	DiscoveredCount  int           `json:"discovered_count"`
	ErrorOnAddCount  int           `json:"error_on_add_count"`
	TimeTaken        string        `json:"time_taken"`
	Points           []model.Point `json:"points"`
}

type DiscoveredPoints struct {
	DeviceInstance string `json:"deviceInstance"`
	Mac            string `json:"mac"`
	NetworkNumber  string `json:"dnet"`
	MacMSTP        string `json:"dadr"`
	ApiTimeout     int    `json:"api_timeout"`
	Timeout        int    `json:"timeout"`
}

// BacnetDevicePoints get points from an added device
func (inst *Client) BacnetDevicePoints(hostIDName string, opts *DiscoveredPoints) ([]model.Point, error) {
	url := fmt.Sprintf("/proxy/ros/api/plugins/api/%s/read/points", bacnetMaster)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetBody(opts).
		SetResult(&discoveredPoints{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*discoveredPoints)
	if out != nil {
		return out.Points, nil
	}
	return nil, nil
}
