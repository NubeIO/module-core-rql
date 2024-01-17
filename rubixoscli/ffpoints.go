package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/dto"
	"net/url"

	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	"github.com/NubeIO/rubix-os/nresty"
)

func (inst *Client) EdgeGetPoints(hostUUID string) ([]model.Point, error) {
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&[]model.Point{}).
		Get("/host/ros/api/points?with_tags=true&with_meta_tags=true"))
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

func (inst *Client) WritePointValue(hostUUID, uuid string, value *model.Priority) (*model.Point, error) {
	body := &pointPriority{
		Priority: value,
	}
	url := fmt.Sprintf("/host/ros/api/points/%s/write", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetBody(body).
		SetResult(&model.Point{}).
		Patch(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Point), nil
}

func (inst *Client) AddPoint(hostUUID string, body *model.Point) (*model.Point, error) {
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&model.Point{}).
		SetBody(body).
		Post("/host/ros/api/points"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Point), nil
}

func (inst *Client) GetPoints(hostUUID string) ([]model.Point, error) {
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&[]model.Point{}).
		Get("/host/ros/api/points?with_tags=true&with_meta_tags=true"))
	if err != nil {
		return nil, err
	}
	var out []model.Point
	out = *resp.Result().(*[]model.Point)
	return out, nil
}

func (inst *Client) GetPointPriority(hostUUID, uuid string) (*model.Point, error) {
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&model.Point{}).
		SetPathParams(map[string]string{"uuid": uuid}).
		Get("/host/ros/api/points/{uuid}?with_priority=true"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Point), nil
}

func (inst *Client) GetPoint(hostUUID, uuid string) (*model.Point, error) {
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&model.Point{}).
		SetPathParams(map[string]string{"uuid": uuid}).
		Get("/host/ros/api/points/{uuid}?with_tags=true&with_meta_tags=true"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Point), nil
}

func (inst *Client) DeletePoint(hostUUID, uuid string) (bool, error) {
	_, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetPathParams(map[string]string{"uuid": uuid}).
		Delete("/host/ros/api/points/{uuid}"))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (inst *Client) EditPoint(hostUUID, uuid string, body *model.Point) (*model.Point, error) {
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetBody(body).
		SetResult(&model.Point{}).
		SetPathParams(map[string]string{"uuid": uuid}).
		Patch("/host/ros/api/points/{uuid}"))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Point), nil
}

func (inst *Client) GetPaginatedPointsByDeviceUUID(hostUUID, deviceUUID string, limit, offset int, search string) (*dto.PaginationResponse, error) {
	requestURL := fmt.Sprintf("/host/ros/api/points/paginate?with_tags=true&with_meta_tags=true&with_priority=true&device_uuid=%v&limit=%v&offset=%v", deviceUUID, limit, offset)
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

func (inst *Client) GetSearchedPointsByDevice(hostUUID, deviceUUID, search string, limit, offset int) ([]model.Point, error) {
	url := fmt.Sprintf("/host/ros/api/points?with_tags=true&with_meta_tags=true&with_priority=true&device_uuid=%v&search_keyword=%v&limit=%v&offset=%v", deviceUUID, search, limit, offset)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&[]model.Point{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	var out []model.Point
	out = *resp.Result().(*[]model.Point)
	return out, nil
}

func (inst *Client) CountSearchedPointsByDevice(hostUUID, deviceUUID, search string) (*dto.Count, error) {
	url := fmt.Sprintf("/host/ros/api/points/count?device_uuid=%v&search_keyword=%v", deviceUUID, search)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&dto.Count{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	var out dto.Count
	out = *resp.Result().(*dto.Count)
	return &out, nil
}

func (inst *Client) UpdatePointTags(hostUUID, pointUUID string, body []model.Tag) ([]model.Tag, error) {
	url := fmt.Sprintf("/host/ros/api/points/%s/tags", pointUUID)
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
