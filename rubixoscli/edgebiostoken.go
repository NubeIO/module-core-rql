package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/externaltoken"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/user"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/dto"
	"github.com/NubeIO/rubix-os/nresty"
)

func (inst *Client) EdgeBiosLogin(hostUUID string, body *user.User) (*dto.TokenResponse, error) {
	url := "/host/bios/api/users/login"
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetBody(body).
		SetResult(&dto.TokenResponse{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*dto.TokenResponse), nil
}

func (inst *Client) EdgeBiosTokens(hostUUID, jwtToken string) (*[]externaltoken.ExternalToken, error) {
	url := "/host/bios/api/tokens"
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetHeader("jwt-token", jwtToken).
		SetResult(&[]externaltoken.ExternalToken{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*[]externaltoken.ExternalToken)
	return data, nil
}

func (inst *Client) EdgeBiosToken(hostUUID, jwtToken, uuid string) (*externaltoken.ExternalToken, error) {
	url := fmt.Sprintf("/host/bios/api/tokens/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetHeader("jwt-token", jwtToken).
		SetResult(&externaltoken.ExternalToken{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*externaltoken.ExternalToken)
	return data, nil
}

func (inst *Client) EdgeBiosTokenGenerate(hostUUID, jwtToken, name string) (*externaltoken.ExternalToken, error) {
	url := "/host/bios/api/tokens/generate"
	body := externaltoken.ExternalToken{Name: name, Blocked: false}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetHeader("jwt-token", jwtToken).
		SetBody(body).
		SetResult(&externaltoken.ExternalToken{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*externaltoken.ExternalToken)
	return data, nil
}

func (inst *Client) EdgeBiosTokenBlock(hostUUID, jwtToken, uuid string, block bool) (*externaltoken.ExternalToken, error) {
	url := fmt.Sprintf("/host/bios/api/tokens/%s/block", uuid)
	body := map[string]bool{"blocked": block}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetHeader("jwt-token", jwtToken).
		SetBody(body).
		SetResult(&externaltoken.ExternalToken{}).
		Put(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*externaltoken.ExternalToken)
	return data, nil
}

func (inst *Client) EdgeBiosTokenRegenerate(hostUUID, jwtToken, uuid string) (*externaltoken.ExternalToken, error) {
	url := fmt.Sprintf("/host/bios/api/tokens/%s/regenerate", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetHeader("jwt-token", jwtToken).
		SetResult(&externaltoken.ExternalToken{}).
		Put(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*externaltoken.ExternalToken)
	return data, nil
}

func (inst *Client) EdgeBiosTokenDelete(hostUUID, jwtToken, uuid string) (bool, error) {
	url := fmt.Sprintf("/host/bios/api/tokens/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetHeader("jwt-token", jwtToken).
		Delete(url))
	if err != nil {
		return false, err
	}
	return resp.String() == "true", nil
}

func (inst *Client) EdgeBiosUpdateUser(hostUUID, jwtToken, username, password string) (bool, error) {
	url := "/host/bios/api/users"
	_, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetHeader("jwt-token", jwtToken).
		SetBody(map[string]string{"username": username, "password": password}).
		Put(url))
	if err != nil {
		return false, err
	}
	return true, nil
}
