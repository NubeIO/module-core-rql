package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-os/nresty"
)

type WhoIsOpts struct {
	Low                   int                `json:"low,omitempty"`
	High                  int                `json:"high,omitempty"`
	NetworkNumber         uint16             `json:"network_number,omitempty"`
	DevicePerPublish      bool               `json:"device_per_publish"`
	DiscoveredDevicesList []discoveredDevice `json:"discovered_devices_list,omitempty"`
	Timeout               int                `json:"timeout,omitempty"`
	ApiTimeout            int                `json:"api_timeout,omitempty"`
	TxnSource             string             `json:"txn_source,omitempty"`
	TxnNumber             string             `json:"txn_number,omitempty"`
}

const bacnetMaster = "bacnetmaster"

type DiscoveredBACnetDevices struct {
	DiscoveredCount  int                `json:"discovered_count"`
	TimeTaken        string             `json:"time_taken"`
	Errors           []string           `json:"errors"`
	DiscoveredDevice []discoveredDevice `json:"discovered_device"`
}

type discoveredDevice struct {
	DeviceId      int    `json:"device_id"`
	Ip            string `json:"ip"`
	Port          int    `json:"port"`
	NetworkNumber int    `json:"network_number"`
	MacMSTP       int    `json:"mac_mstp"`
	Apdu          int    `json:"apdu"`
}

// BacnetWhoIs do a whois on an existing network
func (inst *Client) BacnetWhoIs(hostIDName string, body *WhoIsOpts) (*DiscoveredBACnetDevices, error) {
	url := "/proxy/ros/api/plugins/api/bacnetmaster/master/whois"
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetBody(body).
		SetResult(&DiscoveredBACnetDevices{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*DiscoveredBACnetDevices)
	if out != nil {
		return out, nil
	}
	return nil, nil
}

type DiscoveredDevices struct {
	DiscoveredCount int            `json:"discovered_count"`
	AddedCount      int            `json:"added_count"`
	ErrorOnAddCount int            `json:"error_on_add_count"`
	Errors          []string       `json:"errors"`
	TimeTaken       string         `json:"time_taken"`
	Devices         []model.Device `json:"devices"`
}

// BacnetDiscoveredDevices discover devices will do a whois
func (inst *Client) BacnetDiscoveredDevices(hostIDName string, body *WhoIsOpts) (*DiscoveredDevices, error) {
	url := "/proxy/ros/api/plugins/api/bacnetmaster/read/devices"
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetBody(body).
		SetResult(&DiscoveredDevices{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*DiscoveredDevices)
	if out != nil {
		return out, nil
	}
	return nil, nil
}

type DiscoveredPoints struct {
	ObjectListLength int           `json:"object_list_length"`
	DiscoveredCount  int           `json:"discovered_count"`
	ErrorOnAddCount  int           `json:"error_on_add_count"`
	TimeTaken        string        `json:"time_taken"`
	Errors           []string      `json:"errors"`
	Points           []model.Point `json:"points"`
}

type BACnetOpts struct {
	DeviceInstance   string   `json:"deviceInstance"`
	Mac              string   `json:"mac"`
	NetworkNumber    string   `json:"dnet"`
	MacMSTP          string   `json:"dadr"`
	ApiTimeout       int      `json:"api_timeout"`
	Timeout          int      `json:"timeout"`
	KnownObjectsList []object `json:"known_object_list"`
}

// BacnetDevicePoints get points from an added device
func (inst *Client) BacnetDevicePoints(hostIDName string, opts *BACnetOpts) (*DiscoveredPoints, error) {
	url := fmt.Sprintf("/proxy/ros/api/plugins/api/%s/read/points", bacnetMaster)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetBody(opts).
		SetResult(&DiscoveredPoints{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*DiscoveredPoints)
	if out != nil {
		return out, nil
	}
	return nil, nil
}

type object struct {
	ObjectType     string `json:"objectType"`
	ObjectInstance int    `json:"objectInstance"`
}

type ObjectsList struct {
	DeviceInstance int      `json:"deviceInstance"`
	Mac            string   `json:"mac"`
	ObjectCount    int      `json:"object_count"`
	TimeTaken      string   `json:"time_taken"`
	Errors         []string `json:"errors"`
	ObjectsAdded   int      `json:"objects_added"`
	Objects        []object `json:"objects"`
}

// BacnetDeviceObjects get device objects
func (inst *Client) BacnetDeviceObjects(hostIDName string, opts *BACnetOpts) (*ObjectsList, error) {
	url := fmt.Sprintf("/proxy/ros/api/plugins/api/%s/read/objects/discover", bacnetMaster)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetBody(opts).
		SetResult(&ObjectsList{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*ObjectsList)
	if out != nil {
		return out, nil
	}
	return nil, nil
}

// BacnetDeviceObjectsSize get device objects size/count
func (inst *Client) BacnetDeviceObjectsSize(hostIDName string, opts *BACnetOpts) (*ObjectsList, error) {
	url := fmt.Sprintf("/proxy/ros/api/plugins/api/%s/read/objects/discover", bacnetMaster)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetBody(opts).
		SetResult(&ObjectsList{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*ObjectsList)
	if out != nil {
		return out, nil
	}
	return nil, nil
}
