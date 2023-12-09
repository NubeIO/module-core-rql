package rubixoscli

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	"github.com/NubeIO/rubix-os/nresty"
)

func (inst *Client) AddAlert(hostUUID string, body *model.Alert) (*model.Alert, error) {
	if body == nil {
		return nil, errors.New("alerts body can not be empty")
	}
	if body.HostUUID == "" {
		body.HostUUID = hostUUID
	}
	if body.Severity == "" {
		body.Severity = "info"
	}
	if body.Status == "" {
		body.Status = "active"
	}
	path := fmt.Sprintf("%s", Paths.Alerts.Path)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&model.Alert{}).
		SetBody(body).
		Post(path))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Alert), nil
}

func (inst *Client) GetAlerts(status []string) ([]model.Alert, error) {
	statusParams := url.Values{}
	for _, s := range status {
		statusParams.Add("status", s)
	}
	path := fmt.Sprintf("%s?%s&with_teams=true&with_tags=true&with_meta_tags=true&with_transactions=true&with_tickets=true", Paths.Alerts.Path, statusParams.Encode())
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&[]model.Alert{}).
		Get(path))
	if err != nil {
		return nil, err
	}
	var out []model.Alert
	out = *resp.Result().(*[]model.Alert)
	return out, nil
}

func (inst *Client) GetAlert(alertUUID string) (*model.Alert, error) {
	path := fmt.Sprintf("%s/%s?with_transactions=true&with_teams=true&with_tags=true&with_meta_tags=true", Paths.Alerts.Path, alertUUID)

	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.Alert{}).
		Get(path))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*model.Alert), nil
}

func (inst *Client) GetAlertsByHost(hostUUID string, status []string) ([]model.Alert, error) {
	statusParams := url.Values{}
	for _, s := range status {
		statusParams.Add("status", s)
	}
	path := fmt.Sprintf("%s/host-uuid/%s?with_teams=true&with_tags=true&with_meta_tags=true&with_transactions=true&with_tickets=true&%s", Paths.Alerts.Path, hostUUID, statusParams.Encode())
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&[]model.Alert{}).
		Get(path))
	if err != nil {
		return nil, err
	}
	var out []model.Alert
	out = *resp.Result().(*[]model.Alert)
	return out, nil
}

func (inst *Client) DeleteAlert(hostUUID, uuid string) (*Message, error) {
	path := fmt.Sprintf("%s/%s", Paths.Alerts.Path, uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&Message{}).
		Delete(path))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*Message), nil
}

type AlertUpdateStatusBody struct {
	Status string `json:"status"`
}

func (inst *Client) UpdateAlertStatus(hostUUID, uuid, status string) (*model.Alert, error) {
	path := fmt.Sprintf("%s/%s/status", Paths.Alerts.Path, uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetBody(AlertUpdateStatusBody{
			Status: status,
		}).
		SetResult(&model.Alert{}).
		Patch(path))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Alert), nil
}

func (inst *Client) AssignAlertTeams(hostUUID, uuid string, teams []string) ([]*model.AlertTeam, error) {
	url := fmt.Sprintf("%s/%s/teams", Paths.Alerts.Path, uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetBody(teams).
		SetResult(&[]*model.AlertTeam{}).
		Put(url))
	if err != nil {
		return nil, err
	}

	out := resp.Result().(*[]*model.AlertTeam)
	return *out, nil
}

func (inst *Client) UpdateAlertTag(hostUUID, alertUUID string, alertBody model.Alert) (*model.Alert, error) {
	path := fmt.Sprintf("%s/%s/tags", Paths.Alerts.Path, alertUUID)

	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetBody(alertBody).
		SetResult(&model.Alert{}).
		Put(path))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Alert), nil
}

type AlertMetaTags struct {
	AlertUuid string `json:"alert_uuid"`
	Key       string `json:"key"`
	Value     string `json:"value"`
}

func (inst *Client) UpdateAlertMetaTag(hostUUID, alertUUID string, metaTags []*MetaTags) ([]AlertMetaTags, error) {
	path := fmt.Sprintf("%s/%s/meta-tags", Paths.Alerts.Path, alertUUID)

	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetBody(metaTags).
		SetResult(&[]AlertMetaTags{}).
		Put(path))

	if err != nil {
		if err.Error() == "<nil>" {
			return nil, errors.New(resp.Status())
		}
		return nil, err
	}
	if resp.StatusCode() > 300 {
		return nil, errors.New(resp.Status())
	}
	var out []AlertMetaTags
	out = *resp.Result().(*[]AlertMetaTags)
	return out, nil
}
