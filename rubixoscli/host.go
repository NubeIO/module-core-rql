package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	"github.com/NubeIO/rubix-os/nresty"
)

func (inst *Client) GetHosts() (data []model.Host, err error) {
	path := fmt.Sprintf("%s?with_tags=true&with_comments=true&with_views=true", Paths.Hosts.Path)
	resp, err := inst.Rest.R().
		SetResult(&[]model.Host{}).
		Get(path)
	if err != nil {
		return nil, err
	}
	return *resp.Result().(*[]model.Host), nil
}

func (inst *Client) GetHost(uuid string) (data *model.Host, err error) {
	path := fmt.Sprintf("%s/%s?with_tags=true&with_comments=true&with_views=true", Paths.Hosts.Path, uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.Host{}).
		Get(path))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Host), nil
}

func (inst *Client) AddHost(body *model.Host) (data *model.Host, err error) {
	path := fmt.Sprintf(Paths.Hosts.Path)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetBody(body).
		SetResult(&model.Host{}).
		Post(path))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Host), nil
}

func (inst *Client) UpdateHost(uuid string, body *model.Host) (data *model.Host, err error) {
	path := fmt.Sprintf("%s/%s", Paths.Hosts.Path, uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetBody(body).
		SetResult(&model.Host{}).
		Patch(path))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Host), nil
}

func (inst *Client) DeleteHost(uuid string) (response *Response) {
	path := fmt.Sprintf("%s/%s", Paths.Hosts.Path, uuid)
	response = &Response{}
	resp, err := inst.Rest.R().
		Delete(path)
	return response.buildResponse(resp, err)
}

func (inst *Client) GetHostSchema() string {
	path := fmt.Sprintf("%s/%s", Paths.Hosts.Path, "schema")
	resp, err := inst.Rest.R().
		Get(path)
	if err != nil {
		return "{}"
	}
	return string(resp.Body())
}

func (inst *Client) AddHostComments(body *model.HostComment) (*model.HostComment, error) {
	path := fmt.Sprintf("%s/comments", Paths.Hosts.Path)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetBody(body).
		SetResult(&model.HostComment{}).
		Post(path))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.HostComment), nil
}

func (inst *Client) UpdateHostComments(uuid string, body *model.HostComment) (*model.HostComment, error) {
	path := fmt.Sprintf("%s/comments/%s", Paths.Hosts.Path, uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetBody(body).
		SetResult(&model.HostComment{}).
		Patch(path))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.HostComment), nil
}

func (inst *Client) DeleteHostComments(uuid string) (*Message, error) {
	path := fmt.Sprintf("%s/comments/%s", Paths.Hosts.Path, uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&Message{}).
		Delete(path))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*Message), nil
}

func (inst *Client) UpdateHostTags(hostUUID string, body []*model.HostTag) ([]model.HostTag, error) {
	path := fmt.Sprintf("%s/tags/host_uuid/%s", Paths.Hosts.Path, hostUUID)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetBody(body).
		SetResult(&[]model.HostTag{}).
		Put(path))
	if err != nil {
		return nil, err
	}
	var out []model.HostTag
	out = *resp.Result().(*[]model.HostTag)
	return out, nil
}
