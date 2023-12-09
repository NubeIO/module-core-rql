package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/externaltoken"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/user"
	"github.com/NubeIO/rubix-os/nresty"
)

func (inst *Client) FFLogin(hostUUID string, body *user.User) (*TokenResponse, error) {
	url := fmt.Sprintf("/host/ros/api/users/login")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetBody(body).
		SetResult(&TokenResponse{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*TokenResponse), nil
}

func (inst *Client) FFGenerateToken(hostUUID, jtwToken string, body *TokenCreate) (*externaltoken.ExternalToken, error) {
	url := fmt.Sprintf("/host/ros/api/tokens/generate")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetHeader("Authorization", jtwToken).
		SetBody(body).
		SetResult(&externaltoken.ExternalToken{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*externaltoken.ExternalToken), nil
}
