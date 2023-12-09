package rubixoscli

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"path"
)

func (inst *Client) ProxyHostRosGET(hostUUID, url string) (*resty.Response, error) {
	url = path.Join("/host/ros", url)
	resp, err := inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		Get(url)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (inst *Client) ProxyHostRosPOST(hostUUID, url string, body interface{}) (*resty.Response, error) {
	url = path.Join("/host/ros", url)
	resp, err := inst.Rest.R().
		SetBody(body).
		SetHeader("X-Host", hostUUID).
		Post(url)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (inst *Client) ProxyHostRosPATCH(hostUUID, url string, body interface{}) (*resty.Response, error) {
	url = path.Join("/host/ros", url)
	resp, err := inst.Rest.R().
		SetBody(body).
		SetHeader("X-Host", hostUUID).
		Patch(url)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (inst *Client) ProxyHostRosPUT(hostUUID, url string, body interface{}) (*resty.Response, error) {
	url = path.Join("/host/ros", url)
	resp, err := inst.Rest.R().
		SetBody(body).
		SetHeader("X-Host", hostUUID).
		Put(url)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (inst *Client) ProxyHostRosDELETE(hostUUID, url string) (*resty.Response, error) {
	url = fmt.Sprintf("/host/ros/%s", url)
	resp, err := inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		Delete(url)
	if err != nil {
		return nil, err
	}
	return resp, err
}
