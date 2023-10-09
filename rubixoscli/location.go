package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-os/nresty"
)

func (inst *Client) GetLocations() (data []model.Location, response *Response) {
	path := fmt.Sprintf("%s?with_views=true&with_groups=true&with_hosts=true", Paths.Locations.Path)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetResult(&[]model.Location{}).
		Get(path)
	return *resp.Result().(*[]model.Location), response.buildResponse(resp, err)
}

func (inst *Client) GetLocation(uuid string) (data *model.Location, response *Response) {
	path := fmt.Sprintf("%s/%s?with_views=true&with_groups=true&with_hosts=true", Paths.Locations.Path, uuid)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetResult(&model.Location{}).
		Get(path)
	return resp.Result().(*model.Location), response.buildResponse(resp, err)
}

func (inst *Client) AddLocation(body *model.Location) (*model.Location, error) {
	path := fmt.Sprintf(Paths.Locations.Path)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetBody(body).
		SetResult(&model.Location{}).
		Post(path))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Location), nil
}

func (inst *Client) UpdateLocation(uuid string, body *model.Location) (*model.Location, error) {
	path := fmt.Sprintf("%s/%s", Paths.Locations.Path, uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetBody(body).
		SetResult(&model.Location{}).
		Patch(path))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Location), nil
}

func (inst *Client) DeleteLocation(uuid string) (response *Response) {
	path := fmt.Sprintf("%s/%s", Paths.Locations.Path, uuid)
	response = &Response{}
	resp, err := inst.Rest.R().
		Delete(path)
	return response.buildResponse(resp, err)
}

func (inst *Client) GetLocationSchema() string {
	path := fmt.Sprintf("%s/%s", Paths.Locations.Path, "schema")
	resp, err := inst.Rest.R().
		Get(path)
	if err != nil {
		return "{}"
	}
	return string(resp.Body())
}
