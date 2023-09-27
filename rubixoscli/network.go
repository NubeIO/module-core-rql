package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-os/nresty"
)

func (inst *Client) GetHostNetworks() ([]model.Group, error) {
	path := fmt.Sprintf("%s?with_hosts=true", Paths.Groups.Path)
	resp, err := inst.Rest.R().
		SetResult(&[]model.Group{}).
		Get(path)
	if err != nil {
		return nil, err
	}
	return *resp.Result().(*[]model.Group), nil
}

func (inst *Client) GetHostNetwork(uuid string) (*model.Group, error) {
	path := fmt.Sprintf("%s/%s?with_hosts=true", Paths.Groups.Path, uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.Group{}).
		Get(path))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Group), nil
}

func (inst *Client) AddHostNetwork(body *model.Group) (*model.Group, error) {
	path := fmt.Sprintf(Paths.Groups.Path)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetBody(body).
		SetResult(&model.Group{}).
		Post(path))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Group), nil
}

func (inst *Client) UpdateHostNetwork(uuid string, body *model.Group) (*model.Group, error) {
	path := fmt.Sprintf("%s/%s", Paths.Groups.Path, uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetBody(body).
		SetResult(&model.Group{}).
		Patch(path))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Group), nil
}

func (inst *Client) UpdateHostsStatus(uuid string) (*model.Group, error) {
	path := fmt.Sprintf("%s/%s/%s", Paths.Groups.Path, uuid, "update-hosts-status")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.Group{}).
		Get(path))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Group), nil
}

func (inst *Client) DeleteHostNetwork(uuid string) error {
	path := fmt.Sprintf("%s/%s", Paths.Groups.Path, uuid)
	_, err := nresty.FormatRestyResponse(inst.Rest.R().
		Delete(path))
	if err != nil {
		return err
	}
	return nil
}

func (inst *Client) GetNetworkSchema() string {
	path := fmt.Sprintf("%s/%s", Paths.Groups.Path, "schema")
	resp, err := inst.Rest.R().
		Get(path)
	if err != nil {
		return "{}"
	}
	return string(resp.Body())
}
