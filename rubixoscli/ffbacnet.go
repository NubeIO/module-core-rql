package rubixoscli

import (
	"errors"
	"fmt"
	"strings"

	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
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
func (inst *Client) BacnetWhoIs(hostUUID, pluginName string, body *WhoIsOpts) (*DiscoveredBACnetDevices, error) {
	var url string
	if strings.HasPrefix(pluginName, ModulePrefix) {
		url = "/host/ros/api/modules/module-core-bacnetmaster/api/master/whois"
	} else {
		url = "/host/ros/api/plugins/api/bacnetmaster/master/whois"
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
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
func (inst *Client) BacnetDiscoveredDevices(hostUUID, pluginName string, body *WhoIsOpts) (*DiscoveredDevices, error) {
	var url string
	if strings.HasPrefix(pluginName, ModulePrefix) {
		url = "/host/ros/api/modules/module-core-bacnetmaster/api/read/devices"
	} else {
		url = "/host/ros/api/plugins/api/bacnetmaster/read/devices"
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
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
	DeviceInstance   string         `json:"deviceInstance"`
	Mac              string         `json:"mac"`
	NetworkNumber    string         `json:"dnet"`
	MacMSTP          string         `json:"dadr"`
	ApiTimeout       int            `json:"api_timeout"`
	Timeout          int            `json:"timeout"`
	KnownObjectsList []bacnetObject `json:"known_object_list"`
}

// BacnetDevicePoints get points from an added device
func (inst *Client) BacnetDevicePoints(hostUUID, pluginName string, opts *BACnetOpts) (*DiscoveredPoints, error) {
	var url string
	if strings.HasPrefix(pluginName, ModulePrefix) {
		url = "/host/ros/api/modules/module-core-bacnetmaster/api/read/points"
	} else {
		url = fmt.Sprintf("/host/ros/api/plugins/api/%s/read/points", bacnetMaster)
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
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

type bacnetObject struct {
	ObjectType     string `json:"objectType"`
	ObjectInstance int    `json:"objectInstance"`
}

type ObjectsList struct {
	DeviceInstance int            `json:"deviceInstance"`
	Mac            string         `json:"mac"`
	ObjectCount    int            `json:"object_count"`
	TimeTaken      string         `json:"time_taken"`
	Errors         []string       `json:"errors"`
	ObjectsAdded   int            `json:"objects_added"`
	Objects        []bacnetObject `json:"objects"`
}

// BacnetDeviceObjects get device objects
func (inst *Client) BacnetDeviceObjects(hostUUID, pluginName string, opts *BACnetOpts) (*ObjectsList, error) {
	var url string
	if strings.HasPrefix(pluginName, ModulePrefix) {
		url = "/host/ros/api/modules/module-core-bacnetmaster/api/read/objects/discover"
	} else {
		url = fmt.Sprintf("/host/ros/api/plugins/api/%s/read/objects/discover", bacnetMaster)
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
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

type PriArray struct {
	P1  *float64 `json:"_1"`
	P2  *float64 `json:"_2"`
	P3  *float64 `json:"_3"`
	P4  *float64 `json:"_4"`
	P5  *float64 `json:"_5"`
	P6  *float64 `json:"_6"`
	P7  *float64 `json:"_7"`
	P8  *float64 `json:"_8"`
	P9  *float64 `json:"_9"`
	P10 *float64 `json:"_10"`
	P11 *float64 `json:"_11"`
	P12 *float64 `json:"_12"`
	P13 *float64 `json:"_13"`
	P14 *float64 `json:"_14"`
	P15 *float64 `json:"_15"`
	P16 *float64 `json:"_16"`
}

type WriteArrayBody struct {
	TimeOut  int      `json:"timeout"`
	Priority PriArray `json:"priority"`
}

type PayloadReadPV struct {
	ObjectType     string  `json:"objectType"`
	ObjectInstance string  `json:"objectInstance"`
	DeviceInstance int     `json:"deviceInstance"`
	Mac            string  `json:"mac"`
	Value          float64 `json:"value"`
	TxnSource      string  `json:"txn_source"`
	TxnNumber      string  `json:"txn_number"`
	Error          string  `json:"error"`
}

func (inst *Client) ReadBacnetPointValue(hostUUID, pointUUID, pluginName string) (string, error) {
	var url string
	if strings.HasPrefix(pluginName, ModulePrefix) {
		url = fmt.Sprintf("/host/ros/api/modules/module-core-bacnetmaster/api/%s/read/pv", pointUUID)
	} else {
		url = fmt.Sprintf("/host/ros/api/plugins/api/bacnetmaster/%s/read/pv", pointUUID)
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		Get(url))
	if err != nil {
		return "", err
	}
	out := resp.String()
	return out, nil
}

func (inst *Client) WriteBacnetPointValue(hostUUID, pointUUID, pluginName string, body interface{}) (*PayloadReadPV, error) {
	var url string
	if strings.HasPrefix(pluginName, ModulePrefix) {
		url = fmt.Sprintf("/host/ros/api/modules/module-core-bacnetmaster/api/%s/write/pv", pointUUID)
	} else {
		url = fmt.Sprintf("/host/ros/api/plugins/api/bacnetmaster/%s/write/pv", pointUUID)
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetBody(body).
		SetResult(&PayloadReadPV{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*PayloadReadPV)
	if out != nil {
		return out, nil
	}
	return nil, nil
}

type PayloadPri struct {
	ObjectType     string    `json:"objectType"`
	ObjectInstance string    `json:"objectInstance"`
	Value          *PriArray `json:"value"`
	DeviceInstance int       `json:"deviceInstance"`
	Mac            string    `json:"mac"`
	TxnSource      string    `json:"txn_source"`
	TxnNumber      string    `json:"txn_number"`
	Error          string    `json:"error"`
}

func (inst *Client) ReadBacnetPriorityArray(hostUUID, pointUUID, pluginName string) (*PayloadPri, error) {
	var url string
	if strings.HasPrefix(pluginName, ModulePrefix) {
		url = fmt.Sprintf("/host/ros/api/modules/module-core-bacnetmaster/api/%s/read/pri", pointUUID)
	} else {
		url = fmt.Sprintf("/host/ros/api/plugins/api/bacnetmaster/%s/read/pri", pointUUID)
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&PayloadPri{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*PayloadPri)
	if out != nil {
		return out, nil
	}
	return nil, nil
}

func (inst *Client) WriteBacnetPriorityArray(hostUUID, pointUUID, pluginName string, body interface{}) (*PriArray, error) {
	var url string
	if strings.HasPrefix(pluginName, ModulePrefix) {
		url = fmt.Sprintf("/host/ros/api/modules/module-core-bacnetmaster/api/%s/write/pri", pointUUID)
	} else {
		url = fmt.Sprintf("/host/ros/api/plugins/api/bacnetmaster/%s/write/pri", pointUUID)
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetBody(body).
		SetResult(&PriArray{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*PriArray)
	if out != nil {
		return out, nil
	}
	return nil, nil
}

type WriteBacnetPriorityArrayNull struct {
	Timeout        int `json:"timeout"`
	PriorityNumber int `json:"priority_number"`
}

func (inst *Client) WriteBacnetPriorityArrayNull(hostUUID, pointUUID, pluginName string, body *WriteBacnetPriorityArrayNull) (*PriArray, error) {
	var url string
	if strings.HasPrefix(pluginName, ModulePrefix) {
		url = fmt.Sprintf("/host/ros/api/modules/module-core-bacnetmaster/api/%s/write/pri/null", pointUUID)
	} else {
		url = fmt.Sprintf("/host/ros/api/plugins/api/bacnetmaster/%s/write/pri/null", pointUUID)
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetBody(body).
		SetResult(&PriArray{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*PriArray)
	if out != nil {
		return out, nil
	}
	return nil, nil
}

// ------------ BACNET Objects Discovery

type DeviceObjectsBody struct {
	DeviceUUID       string         `json:"deviceUUID"`
	ApiTimeout       int            `json:"api_timeout"`
	Timeout          int            `json:"timeout"`
	KnownObjectsList []bacnetObject `json:"known_object_list"`
}

// BacnetDeviceObjectsSize get device objects size/count
func (inst *Client) BacnetDeviceObjectsSize(hostUUID, pluginName string, opts *DeviceObjectsBody) (*ObjectsList, error) {
	var url string
	if strings.HasPrefix(pluginName, ModulePrefix) {
		url = fmt.Sprintf("/host/ros/api/modules/module-core-bacnetmaster/api/read/device/objects/size")
	} else {
		url = fmt.Sprintf("/host/ros/api/plugins/api/bacnetmaster/read/device/objects/size")
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
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

type ErrorsWithTime struct {
	Errors    []string `json:"errors"`
	TimeTaken string   `json:"time_taken"`
}

type DeviceObjectsList struct {
	DeviceInstance int            `json:"deviceInstance"`
	Mac            string         `json:"mac"`
	ObjectCount    int            `json:"object_count"`
	ObjectsAdded   int            `json:"objects_added"`
	Objects        []bacnetObject `json:"objects"`
	ErrorsWithTime
}

// BacnetDeviceObjectsList get device objects list
func (inst *Client) BacnetDeviceObjectsList(hostUUID, pluginName string, opts *DeviceObjectsBody) (*DeviceObjectsList, error) {
	var url string
	if strings.HasPrefix(pluginName, ModulePrefix) {
		url = fmt.Sprintf("/host/ros/api/modules/module-core-bacnetmaster/api/read/objects/discover/device")
	} else {
		url = fmt.Sprintf("/host/ros/api/plugins/api/bacnetmaster/read/objects/discover/device")
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetBody(opts).
		SetResult(&DeviceObjectsList{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*DeviceObjectsList)
	if out != nil {
		return out, nil
	}
	return nil, nil
}

// BacnetDeviceObjectsListCount get device objects list count
func (inst *Client) BacnetDeviceObjectsListCount(hostUUID, deviceUUID, pluginName string) (any, error) {
	var url string
	if strings.HasPrefix(pluginName, ModulePrefix) {
		url = fmt.Sprintf("/host/ros/api/modules/module-core-bacnetmaster/api/read/objects/discovered/count/object-count-%s", deviceUUID)
	} else {
		url = fmt.Sprintf("/host/ros/api/plugins/api/bacnetmaster/read/objects/discovered/count/object-count-%s", deviceUUID)
	}
	var temp *int
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&temp).
		Get(url))
	if err != nil {
		return 0, err
	}
	if resp != nil {
		return resp.Result(), nil
	}
	return 0, nil
}

// ------------ BACNET Points Discovery

type DeviceDiscoveredPoints struct {
	ObjectListLength int           `json:"object_list_length"`
	DiscoveredCount  int           `json:"discovered_count"`
	ErrorOnAddCount  int           `json:"error_on_add_count"`
	Points           []model.Point `json:"points"`
	ErrorsWithTime
}

// BacnetDevicePointsList get device points by its objects list
func (inst *Client) BacnetDevicePointsList(hostUUID, pluginName string, opts *DeviceObjectsBody) (*DeviceDiscoveredPoints, error) {
	if opts == nil {
		return nil, errors.New("options body can not be empty")
	}
	if opts.KnownObjectsList == nil {
		return nil, errors.New("please pass in the objects to discover")
	}
	var url string
	if strings.HasPrefix(pluginName, ModulePrefix) {
		url = fmt.Sprintf("/host/ros/api/modules/module-core-bacnetmaster/api/read/nonpics/points/device")
	} else {
		url = fmt.Sprintf("/host/ros/api/plugins/api/bacnetmaster/read/nonpics/points/device")
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetBody(opts).
		SetResult(&DeviceDiscoveredPoints{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*DeviceDiscoveredPoints)
	if out != nil {
		return out, nil
	}
	return nil, nil
}

// BacnetDevicePointsListCount get device points list count
func (inst *Client) BacnetDevicePointsListCount(hostUUID, deviceUUID, pluginName string) (any, error) {
	var url string
	if strings.HasPrefix(pluginName, ModulePrefix) {
		url = fmt.Sprintf("/host/ros/api/modules/module-core-bacnetmaster/api/read/objects/discovered/count/points-count-%s", deviceUUID)
	} else {
		url = fmt.Sprintf("/host/ros/api/plugins/api/bacnetmaster/read/objects/discovered/count/points-count-%s", deviceUUID)
	}
	var temp *int
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&temp).
		Get(url))
	if err != nil {
		return 0, err
	}
	if resp != nil {
		return resp.Result(), nil
	}
	return 0, nil
}
