package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/rubix-os/interfaces"
	"github.com/NubeIO/rubix-os/nresty"
	"io"
)

func (inst *Client) StoreListModules() ([]interfaces.Module, error, error) {
	url := fmt.Sprintf("/api/store/modules")
	resp, connectionErr, requestErr := nresty.FormatRestyV2Response(inst.Rest.R().
		SetResult(&[]interfaces.Module{}).
		Get(url))
	if connectionErr != nil || requestErr != nil {
		return nil, connectionErr, requestErr
	}
	return *resp.Result().(*[]interfaces.Module), nil, nil
}

func (inst *Client) StoreUploadModule(fileName string, reader io.Reader) (*interfaces.UploadResponse, error) {
	url := fmt.Sprintf("/api/store/modules")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&interfaces.UploadResponse{}).
		SetFileReader("file", fileName, reader).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*interfaces.UploadResponse), nil
}
