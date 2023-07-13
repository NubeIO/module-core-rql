package rubixoscli

import (
	"fmt"

	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-os/interfaces"
	"github.com/NubeIO/rubix-os/nresty"
)

func (inst *Client) GetMembers() ([]model.Member, error) {
	url := "/api/members"
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&[]model.Member{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	var out []model.Member
	out = *resp.Result().(*[]model.Member)
	return out, nil
}

func (inst *Client) AddMember(body *model.Member) (*model.Member, error) {
	url := "/api/apps/members"
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetBody(body).
		SetResult(&model.Member{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*model.Member)
	return out, nil
}

func (inst *Client) UpdateMember(memberUUID string, body *model.Member) (*model.Member, error) {
	url := fmt.Sprintf("/api/members/%s", memberUUID)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetBody(body).
		SetResult(&model.Member{}).
		Patch(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*model.Member)
	return out, nil
}

func (inst *Client) DeleteMember(memberUUID string) (bool, error) {
	url := fmt.Sprintf("/api/members/%s", memberUUID)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		Delete(url))
	if err != nil {
		return false, err
	}
	return resp.String() == "true", nil
}

func (inst *Client) BulkDeleteMember(memberUUID []string) (bool, error) {
	for _, aUUID := range memberUUID {
		url := fmt.Sprintf("/api/members/%s", aUUID)
		resp, err := nresty.FormatRestyResponse(inst.Rest.R().
			Delete(url))
		if err != nil || resp.String() != "true" {
			return false, err
		}
	}
	return true, nil
}

func (inst *Client) VerifyMember(memberName string) (*interfaces.Message, error) {
	url := fmt.Sprintf("/api/members/verify/%s", memberName)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&interfaces.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*interfaces.Message), nil
}
