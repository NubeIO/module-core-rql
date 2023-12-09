package rubixoscli

import (
	"fmt"

	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	"github.com/NubeIO/rubix-os/nresty"
)

func (inst *Client) GetFCMServerKey() (data *model.FcmServer, err error) {
	path := fmt.Sprintf("/api/fcm-server")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.FcmServer{}).
		Get(path))
	return resp.Result().(*model.FcmServer), err
}

func (inst *Client) UpsertFCMServerKey(body *model.FcmServer) (*model.FcmServer, error) {
	path := fmt.Sprintf("/api/fcm-server")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.FcmServer{}).
		SetBody(body).
		Put(path))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.FcmServer), nil
}
