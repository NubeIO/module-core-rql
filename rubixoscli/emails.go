package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/dto"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	"github.com/NubeIO/rubix-os/nresty"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

func (inst *Client) GetEmail(emailUUID string) (*model.Email, error) {
	url := fmt.Sprintf("/api/emails/%s", emailUUID)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.Ticket{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*model.Email)
	return out, nil
}

func (inst *Client) GetEmails() ([]model.Email, error) {
	url := "/api/emails"
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&[]model.Email{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	var out []model.Email
	out = *resp.Result().(*[]model.Email)
	return out, nil
}

func (inst *Client) SendEmail(body *model.Email) (*model.Email, error) {
	url := "/api/emails/send"
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetBody(body).
		SetResult(&model.Email{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*model.Email)
	return out, nil
}

func (inst *Client) DeleteEmail(emailUUID string) (bool, error) {
	url := fmt.Sprintf("/api/emails/%s", emailUUID)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		Delete(url))
	if err != nil {
		return false, err
	}
	return resp.String() == "true", nil
}

func (inst *Client) ResendEmail(emailUUID string) (*model.Email, error) {
	url := fmt.Sprintf("/api/emails/%s/resend", emailUUID)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.Email{}).
		Put(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*model.Email)
	return out, nil
}

func (inst *Client) GetEmailConfiguration() (*model.EmailConfig, error) {
	url := "/api/email-config"
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.EmailConfig{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*model.EmailConfig)
	return out, nil
}

func (inst *Client) UpdateEmailConfiguration(body *model.EmailConfig) (bool, error) {
	url := "/api/email-config"
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetBody(body).
		SetResult(&model.EmailConfig{}).
		Put(url))
	if err != nil {
		return false, err
	}

	return resp.String() == "true", nil
}

func (inst *Client) AttachmentDownload(fileName, destination string) (*Message, error) {
	url := fmt.Sprintf("/attachments/%s", fileName)
	fmt.Println(url)

	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		Get(url))
	if err != nil {
		return nil, err
	}
	SaveFile(resp.Body(), destination)
	return &Message{
		Message: fmt.Sprintf("File downloaded successfully to %s", destination),
	}, nil
}

func (inst *Client) UploadAttachments(attachments []string) (*[]dto.EmailAttachment, error) {
	url := fmt.Sprintf("/api/emails/attachments/upload")

	attachmentsData := make([]*resty.MultipartField, 0)
	for _, attachment := range attachments {
		file, err := os.Open(attachment)
		if err != nil {
			log.Error(err)
			return nil, err
		}

		attachmentsData = append(attachmentsData, &resty.MultipartField{
			Param:    "attachments",
			FileName: filepath.Base(attachment),
			Reader:   file,
		})
	}

	// Send a multipart form data request with all attachments
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("Content-Type", "multipart/form-data").
		SetMultipartFields(attachmentsData...).
		SetResult(&[]dto.EmailAttachment{}).
		Post(url))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	log.Debugf("Response body: %s", resp.Body())
	respAttachments := resp.Result().(*[]dto.EmailAttachment)
	return respAttachments, nil
}
