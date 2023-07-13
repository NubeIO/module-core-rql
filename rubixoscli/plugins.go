package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/rubix-os/installer"
	"github.com/NubeIO/rubix-os/interfaces"
	"github.com/NubeIO/rubix-os/nresty"
	"io"
)

func (inst *Client) StoreListPlugins() ([]installer.BuildDetails, error, error) {
	url := fmt.Sprintf("/api/store/plugins")
	resp, connectionErr, requestErr := nresty.FormatRestyV2Response(inst.Rest.R().
		SetResult(&[]installer.BuildDetails{}).
		Get(url))
	if connectionErr != nil || requestErr != nil {
		return nil, connectionErr, requestErr
	}
	return *resp.Result().(*[]installer.BuildDetails), nil, nil
}

func (inst *Client) StoreUploadPlugin(fileName string, reader io.Reader) (*interfaces.UploadResponse, error) {
	url := fmt.Sprintf("/api/store/plugins")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&interfaces.UploadResponse{}).
		SetFileReader("file", fileName, reader).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*interfaces.UploadResponse), nil
}
