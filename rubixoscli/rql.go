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

func (inst *Client) TestRunRule(hostIDName string, body *RQLRule) (*RQLRuleResponse, error) {
	url := fmt.Sprintf("/api/modules/module-core-rql/rules/dry")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&RQLRuleResponse{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*RQLRuleResponse), nil
}

func (inst *Client) SelectAllRules(hostIDName string) ([]RQLRule, error) {
	url := fmt.Sprintf("/api/modules/module-core-rql/rules")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
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

func (inst *Client) SelectRule(hostIDName, uuid string) (*RQLRule, error) {
	url := fmt.Sprintf("/api/modules/module-core-rql/rules/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&RQLRule{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*RQLRule), nil
}

func (inst *Client) RunExistingRule(hostIDName, uuid string) (*RQLRuleResponse, error) {
	url := fmt.Sprintf("/api/modules/module-core-rql/rules/run/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&RQLRuleResponse{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*RQLRuleResponse), nil
}

func (inst *Client) AddRule(hostIDName string, body *RQLRule) (*RQLRule, error) {
	url := fmt.Sprintf("/api/modules/module-core-rql/rules")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
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

func (inst *Client) UpdateRule(hostIDName, uuid string, body *RQLRule) (*RQLRule, error) {
	url := fmt.Sprintf("/api/modules/module-core-rql/rules/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&RQLRule{}).
		SetBody(body).
		Patch(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*RQLRule), nil
}

func (inst *Client) DeleteRule(hostIDName, uuid string) (*Message, error) {
	url := fmt.Sprintf("/api/modules/module-core-rql/rules/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&Message{}).
		Delete(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*Message), nil
}

type RQLVariables struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Variable any    `json:"variable"`
	Password string `json:"password"`
}

func (inst *Client) SelectAllVars(hostIDName string) ([]RQLVariables, error) {
	url := fmt.Sprintf("/api/modules/module-core-rql/vars")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
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

func (inst *Client) SelectOneVar(hostIDName, uuid string) (*RQLVariables, error) {
	url := fmt.Sprintf("/api/modules/module-core-rql/vars/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&RQLVariables{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*RQLVariables), nil
}

func (inst *Client) AddVar(hostIDName string, body *RQLVariables) (*RQLVariables, error) {
	url := fmt.Sprintf("/api/modules/module-core-rql/vars")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&RQLVariables{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*RQLVariables), nil
}

func (inst *Client) UpdateVar(hostIDName, uuid string, body *RQLVariables) (*RQLVariables, error) {
	url := fmt.Sprintf("/api/modules/module-core-rql/vars/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&RQLVariables{}).
		SetBody(body).
		Patch(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*RQLVariables), nil
}

func (inst *Client) DeleteVar(hostIDName, uuid string) (*Message, error) {
	url := fmt.Sprintf("/api/modules/module-core-rql/vars/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&Message{}).
		Delete(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*Message), nil
}
