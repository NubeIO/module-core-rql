package rubixoscli

import (
	"fmt"

	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-os/nresty"
)

func (inst *Client) AttachViews(teamUUID string, viewUUIDs []string) ([]model.View, error) {
	url := fmt.Sprintf("/api/teams/%s/views", teamUUID)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetBody(viewUUIDs).
		SetResult(&[]model.View{}).
		Put(url))
	if err != nil {
		return nil, err
	}
	var out []model.View
	out = *resp.Result().(*[]model.View)
	return out, nil
}

func (inst *Client) GetTeam(teamUUID string) (*model.Team, error) {
	url := fmt.Sprintf("/api/teams/%s?with_members=true&with_views=true", teamUUID)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.Team{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*model.Team)
	return out, nil
}

func (inst *Client) GetTeams() ([]model.Team, error) {
	url := "/api/teams?with_members=true&with_views=true"
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&[]model.Team{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	var out []model.Team
	out = *resp.Result().(*[]model.Team)
	return out, nil
}

func (inst *Client) AddTeam(body *model.Team) (*model.Team, error) {
	url := "/api/teams"
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetBody(body).
		SetResult(&model.Team{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*model.Team)
	return out, nil
}

func (inst *Client) DeleteTeam(teamUUID string) (bool, error) {
	url := fmt.Sprintf("/api/teams/%s", teamUUID)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		Delete(url))
	if err != nil {
		return false, err
	}
	return resp.String() == "true", nil
}

func (inst *Client) BulkDeleteTeam(teamUUIDs []string) (bool, error) {
	for _, aUUID := range teamUUIDs {
		url := fmt.Sprintf("/api/teams/%s", aUUID)
		resp, err := nresty.FormatRestyResponse(inst.Rest.R().
			Delete(url))
		if err != nil || resp.String() != "true" {
			return false, err
		}
	}
	return true, nil
}

func (inst *Client) UpdateTeam(teamUUID string, body *model.Team) (*model.Team, error) {
	url := fmt.Sprintf("/api/teams/%s", teamUUID)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetBody(body).
		SetResult(&model.Team{}).
		Patch(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*model.Team)
	return out, nil
}

func (inst *Client) UpdateTeamMember(teamUUID string, memberUUIDs []string) ([]model.Team, error) {
	url := fmt.Sprintf("/api/teams/%s/members", teamUUID)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetBody(memberUUIDs).
		SetResult(&[]model.Team{}).
		Put(url))
	if err != nil {
		return nil, err
	}
	var out []model.Team
	out = *resp.Result().(*[]model.Team)
	return out, nil
}
