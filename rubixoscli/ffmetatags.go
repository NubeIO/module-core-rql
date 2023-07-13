package rubixoscli

import (
	"errors"
	"fmt"
	"github.com/NubeIO/rubix-os/nresty"
)

type MetaTags struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type NetworkTags struct {
	NetworkUuid string `json:"network_uuid"`
	Key         string `json:"key"`
	Value       string `json:"value"`
}

func (inst *Client) FFAddNetworkTags(hostIDName, uuid string, body []*MetaTags) ([]NetworkTags, error) {
	tagsType := "networks"
	url := fmt.Sprintf("/proxy/ros/api/%s/meta_tags/uuid/%s", tagsType, uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&[]NetworkTags{}).
		SetBody(body).
		Put(url))
	if err != nil {
		if err.Error() == "<nil>" {
			return nil, errors.New(resp.Status())
		}
		return nil, err
	}
	if resp.StatusCode() > 300 {
		return nil, errors.New(resp.Status())
	}
	var out []NetworkTags
	out = *resp.Result().(*[]NetworkTags)
	return out, nil
}

type DeviceTags struct {
	DeviceUuid string `json:"device_uuid"`
	Key        string `json:"key"`
	Value      string `json:"value"`
}

func (inst *Client) FFAddDeviceTags(hostIDName, uuid string, body []*MetaTags) ([]DeviceTags, error) {
	tagsType := "devices"
	url := fmt.Sprintf("/proxy/ros/api/%s/meta_tags/uuid/%s", tagsType, uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&[]DeviceTags{}).
		SetBody(body).
		Put(url))
	if err != nil {
		if err.Error() == "<nil>" {
			return nil, errors.New(resp.Status())
		}
		return nil, err
	}
	if resp.StatusCode() > 300 {
		return nil, errors.New(resp.Status())
	}
	var out []DeviceTags
	out = *resp.Result().(*[]DeviceTags)
	return out, nil
}

type PointTags struct {
	DeviceUuid string `json:"device_uuid"`
	Key        string `json:"key"`
	Value      string `json:"value"`
}

func (inst *Client) FFAddPointTags(hostIDName, uuid string, body []*MetaTags) ([]PointTags, error) {
	tagsType := "points"
	url := fmt.Sprintf("/proxy/ros/api/%s/meta_tags/uuid/%s", tagsType, uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&[]PointTags{}).
		SetBody(body).
		Put(url))
	if err != nil {
		if err.Error() == "<nil>" {
			return nil, errors.New(resp.Status())
		}
		return nil, err
	}
	if resp.StatusCode() > 300 {
		return nil, errors.New(resp.Status())
	}
	var out []PointTags
	out = *resp.Result().(*[]PointTags)
	return out, nil
}
