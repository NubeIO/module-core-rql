package rubixoscli

import (
	"fmt"

	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-os/nresty"
)

func (inst *Client) EdgeGetPoints(hostIDName string) ([]model.Point, error) {
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&[]model.Point{}).
		Get("/proxy/ros/api/points?with_tags=true&with_meta_tags=true"))
	if err != nil {
		return nil, err
	}
	var out []model.Point
	out = *resp.Result().(*[]model.Point)
	return out, nil
}

type pointPriority struct {
	Priority *model.Priority
}

func (inst *Client) WritePointValue(hostIDName, uuid string, value *model.Priority) (*model.Point, error) {
	body := &pointPriority{
		Priority: value,
	}
	url := fmt.Sprintf("/proxy/ros/api/points/write/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetBody(body).
		SetResult(&model.Point{}).
		Patch(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Point), nil
}

func (inst *Client) AddPoint(hostIDName string, body *model.Point) (*model.Point, error) {
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&model.Point{}).
		SetBody(body).
		Post("/proxy/ros/api/points"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Point), nil
}

func (inst *Client) GetPoints(hostIDName string) ([]model.Point, error) {
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&[]model.Point{}).
		Get("/proxy/ros/api/points?with_tags=true&with_meta_tags=true"))
	if err != nil {
		return nil, err
	}
	var out []model.Point
	out = *resp.Result().(*[]model.Point)
	return out, nil
}

func (inst *Client) GetPointsByModule(hostIDName, moduleName string) ([]*model.Point, error) {
	network, err := inst.GetNetworkByPluginName(hostIDName, moduleName, true)
	if err != nil {
		return nil, err
	}
	var out []*model.Point
	for _, device := range network.Devices {
		out = append(out, device.Points...)
	}
	return out, nil
}

func (inst *Client) GetPointPriority(hostIDName, uuid string) (*model.Point, error) {
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&model.Point{}).
		SetPathParams(map[string]string{"uuid": uuid}).
		Get("/proxy/ros/api/points/{uuid}?with_priority=true"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Point), nil
}

func (inst *Client) GetPoint(hostIDName, uuid string) (*model.Point, error) {
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&model.Point{}).
		SetPathParams(map[string]string{"uuid": uuid}).
		Get("/proxy/ros/api/points/{uuid}?with_tags=true&with_meta_tags=true"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Point), nil
}

func (inst *Client) DeletePoint(hostIDName, uuid string) (bool, error) {
	_, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetPathParams(map[string]string{"uuid": uuid}).
		Delete("/proxy/ros/api/points/{uuid}"))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (inst *Client) EditPoint(hostIDName, uuid string, body *model.Point) (*model.Point, error) {
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetBody(body).
		SetResult(&model.Point{}).
		SetPathParams(map[string]string{"uuid": uuid}).
		Patch("/proxy/ros/api/points/{uuid}"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Point), nil
}

func (inst *Client) FFGetPluginSchemaPoint(hostIDName, pluginName string) ([]byte, error) {
	url := fmt.Sprintf("/proxy/ros/api/plugins/api/%s/schema/json/point", pluginName)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}
