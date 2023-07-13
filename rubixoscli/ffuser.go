package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/externaltoken"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/user"
	"github.com/NubeIO/rubix-os/nresty"
)

func (inst *Client) FFLogin(hostIDName string, body *user.User) (*TokenResponse, error) {
	url := fmt.Sprintf("/proxy/ros/api/users/login")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetBody(body).
		SetResult(&TokenResponse{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*TokenResponse), nil
}

func (inst *Client) FFGenerateToken(hostIDName, jtwToken string, body *TokenCreate) (*externaltoken.ExternalToken, error) {
	url := fmt.Sprintf("/proxy/ros/api/tokens/generate")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetHeader("Authorization", jtwToken).
		SetBody(body).
		SetResult(&externaltoken.ExternalToken{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*externaltoken.ExternalToken), nil
}
