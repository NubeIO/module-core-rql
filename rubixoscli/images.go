package rubixoscli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/NubeIO/rubix-os/nresty"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

type UploadImageResponse struct {
	Image string `json:"image"`
}

func (inst *Client) UploadImage(path string) (*UploadImageResponse, error) {
	url := fmt.Sprintf("/api/images")

	file, err := os.Open(path)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// send multipart form data request
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("Content-Type", "multipart/form-data").
		SetMultipartFields(&resty.MultipartField{
			Param:    "file",
			FileName: filepath.Base(path),
			Reader:   file,
			// ContentType: "application/file",
		}).
		SetResult(&UploadImageResponse{}).
		Post(url))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	log.Debugf("Response body: %s", resp.Body())
	return resp.Result().(*UploadImageResponse), nil
}
