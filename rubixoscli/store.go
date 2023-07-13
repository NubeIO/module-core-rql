package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/rubix-os/interfaces"
	"github.com/NubeIO/rubix-os/nresty"
	"io"
)

func (inst *Client) UploadAppOnAppStore(appName, version, arch, fileName string, reader io.Reader) (
	*interfaces.UploadResponse, error) {
	url := fmt.Sprintf("/api/store/apps?name=%s&version=%s&arch=%s", appName, version, arch)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&interfaces.UploadResponse{}).
		SetFileReader("file", fileName, reader).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*interfaces.UploadResponse), nil
}

func (inst *Client) CheckAppExistence(appName, version, arch string) error {
	url := fmt.Sprintf("/api/store/apps/exists?name=%s&arch=%s&version=%s", appName, arch, version)
	_, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&interfaces.FoundMessage{}).
		Get(url))
	return err
}
