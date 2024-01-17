package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/dto"
	"github.com/NubeIO/rubix-os/nresty"
	"io"
)

func (inst *Client) UploadAppOnAppStore(appName, version, arch, fileName string, reader io.Reader) (
	*dto.UploadResponse, error) {
	url := fmt.Sprintf("/api/store/apps?name=%s&version=%s&arch=%s", appName, version, arch)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&dto.UploadResponse{}).
		SetFileReader("file", fileName, reader).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*dto.UploadResponse), nil
}

func (inst *Client) CheckAppExistence(appName, version, arch string) error {
	url := fmt.Sprintf("/api/store/apps/exists?name=%s&arch=%s&version=%s", appName, arch, version)
	_, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&dto.FoundMessage{}).
		Get(url))
	return err
}
