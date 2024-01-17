package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/dto"
	"net/url"

	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	"github.com/NubeIO/rubix-os/nresty"
)

type PointDiscoverableCountType struct {
	DeviceInstance int         `json:"deviceInstance"`
	Mac            string      `json:"mac"`
	ObjectCount    int         `json:"object_count"`
	TimeTaken      string      `json:"time_taken"`
	Errors         interface{} `json:"errors"`
	ObjectsAdded   int         `json:"objects_added"`
	Objects        interface{} `json:"objects"`
}

func (inst *Client) AddDevice(hostUUID string, device *model.Device) (*model.Device, error) {
	url := fmt.Sprintf("/host/ros/api/devices?with_tags=true&with_meta_tags=true")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&model.Device{}).
		SetBody(device).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Device), nil
}

func (inst *Client) GetDevices(hostUUID string, withPoints ...bool) ([]model.Device, error) {
	url := fmt.Sprintf("/host/ros/api/devices?with_tags=true&with_meta_tags=true")
	if len(withPoints) > 0 {
		if withPoints[0] == true {
			url = fmt.Sprintf("/host/ros/api/devices?with_points=true&with_tags=true&with_meta_tags=true")
		}
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&[]model.Device{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	var out []model.Device
	out = *resp.Result().(*[]model.Device)
	return out, nil
}

func (inst *Client) GetPaginatedDevices(hostUUID string, networkUUID string, limit, offset int, search string, withPoints bool) (*dto.PaginationResponse, error) {
	requestURL := fmt.Sprintf("/host/ros/api/devices/paginate?with_tags=true&with_meta_tags=true&network_uuid=%s&limit=%v&offset=%v", networkUUID, limit, offset)
	if withPoints {
		requestURL += "&with_devices=true"
	}
	if search != "" {
		requestURL += "&search_keyword=" + url.QueryEscape(search) // Ensure proper URL encoding for search value
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&dto.PaginationResponse{}).
		Get(requestURL))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*dto.PaginationResponse)
	return out, nil
}

func (inst *Client) GetDevice(hostUUID, uuid string, withPoints ...bool) (*model.Device, error) {
	url := fmt.Sprintf("/host/ros/api/devices/%s?with_tags=true&with_meta_tags=true", uuid)
	if len(withPoints) > 0 {
		if withPoints[0] == true {
			url = fmt.Sprintf("/host/ros/api/devices/%s?with_points=true&with_tags=true&with_meta_tags=true&with_priority=true", uuid)
		}
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&model.Device{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Device), nil
}

func (inst *Client) EditDevice(hostUUID, uuid string, device *model.Device) (*model.Device, error) {
	url := fmt.Sprintf("/host/ros/api/devices/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&model.Device{}).
		SetBody(device).
		Patch(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Device), nil
}

func (inst *Client) DeleteDevice(hostUUID, uuid string) (bool, error) {
	_, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetPathParams(map[string]string{"uuid": uuid}).
		Delete("/host/ros/api/devices/{uuid}"))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (inst *Client) UpdateDeviceTag(hostUUID, deviceUUID string, body []model.Tag) ([]model.Tag, error) {
	url := fmt.Sprintf("/host/ros/api/devices/%s/tags", deviceUUID)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&[]model.Tag{}).
		SetBody(body).
		Put(url))
	if err != nil {
		return nil, err
	}

	var out []model.Tag
	out = *resp.Result().(*[]model.Tag)
	return out, nil
}

func (inst *Client) PingDevice(hostUUID, pluginName string, body interface{}) (bool, error) {
	url := fmt.Sprintf("/host/ros/api/modules/%s/api/devices/ping", pluginName)
	_, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetBody(body).
		Post(url))
	if err != nil {
		return false, err
	}

	return true, nil
}

func (inst *Client) GetDiscoverablePointsCount(hostUUID string, body interface{}) (*PointDiscoverableCountType, error) {
	url := fmt.Sprintf("/host/ros/api/modules/module-core-bacnetmaster/api/read/device/objects/size")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&PointDiscoverableCountType{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	var out PointDiscoverableCountType
	out = *resp.Result().(*PointDiscoverableCountType)
	return &out, nil
}
