package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/rubix-os/nresty"
	"time"
)

type Result struct {
	Result    interface{} `json:"result"`
	Timestamp string      `json:"timestamp"`
	Time      time.Time   `json:"time"`
}

type RQLRule struct {
	UUID              string   `json:"uuid"`
	Name              string   `json:"name"`
	LatestRunDate     string   `json:"latest_run_date"`
	Script            string   `json:"script"`
	Schedule          string   `json:"schedule"`
	Enable            bool     `json:"enable"`
	ResultStorageSize int      `json:"result_storage_size"`
	Result            []Result `json:"result"`
}

type RQLRuleResponse struct {
	Return    interface{} `json:"return"`
	Err       interface{} `json:"err"`
	TimeTaken string      `json:"time_taken"`
}

func (inst *Client) TestRunRule(hostUUID string, body *RQLRule) (*RQLRuleResponse, error) {
	url := fmt.Sprintf("/api/modules/module-core-rql/api/rules/dry")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&RQLRuleResponse{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*RQLRuleResponse), nil
}

func (inst *Client) SelectAllRules(hostUUID string) ([]RQLRule, error) {
	url := fmt.Sprintf("/api/modules/module-core-rql/api/rules")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&[]RQLRule{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	var out []RQLRule
	out = *resp.Result().(*[]RQLRule)
	return out, nil
}

func (inst *Client) SelectRule(hostUUID, uuid string) (*RQLRule, error) {
	url := fmt.Sprintf("/api/modules/module-core-rql/api/rules/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&RQLRule{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*RQLRule), nil
}

func (inst *Client) RunExistingRule(hostUUID, uuid string) (*RQLRuleResponse, error) {
	url := fmt.Sprintf("/api/modules/module-core-rql/api/rules/%s/run", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&RQLRuleResponse{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*RQLRuleResponse), nil
}

func (inst *Client) AddRule(hostUUID string, body *RQLRule) (*RQLRule, error) {
	url := fmt.Sprintf("/api/modules/module-core-rql/api/rules")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&RQLRule{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return resp.Result().(*RQLRule), nil
}

func (inst *Client) UpdateRule(hostUUID, uuid string, body *RQLRule) (*RQLRule, error) {
	url := fmt.Sprintf("/api/modules/module-core-rql/api/rules/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&RQLRule{}).
		SetBody(body).
		Patch(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*RQLRule), nil
}

func (inst *Client) DeleteRule(hostUUID, uuid string) (*Message, error) {
	url := fmt.Sprintf("/api/modules/module-core-rql/api/rules/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&Message{}).
		Delete(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*Message), nil
}

type RQLVariables struct {
	UUID       string `json:"uuid"`
	ScriptUUID string `json:"script_uuid"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Value      any    `json:"value"`
	Required   bool   `json:"required"`
}

func (inst *Client) SelectAllVars(hostUUID string) ([]RQLVariables, error) {
	url := fmt.Sprintf("/api/modules/module-core-rql/api/vars")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&[]RQLVariables{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	var out []RQLVariables
	out = *resp.Result().(*[]RQLVariables)
	return out, nil
}

func (inst *Client) SelectOneVar(hostUUID, uuid string) (*RQLVariables, error) {
	url := fmt.Sprintf("/api/modules/module-core-rql/api/vars/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&RQLVariables{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*RQLVariables), nil
}

func (inst *Client) AddVar(hostUUID string, body *RQLVariables) (*RQLVariables, error) {
	url := fmt.Sprintf("/api/modules/module-core-rql/api/vars")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&RQLVariables{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*RQLVariables), nil
}

func (inst *Client) UpdateVar(hostUUID, uuid string, body *RQLVariables) (*RQLVariables, error) {
	url := fmt.Sprintf("/api/modules/module-core-rql/api/vars/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&RQLVariables{}).
		SetBody(body).
		Patch(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*RQLVariables), nil
}

func (inst *Client) DeleteVar(hostUUID, uuid string) (*Message, error) {
	url := fmt.Sprintf("/api/modules/module-core-rql/api/vars/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&Message{}).
		Delete(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*Message), nil
}
