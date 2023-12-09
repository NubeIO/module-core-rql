package rubixoscli

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/NubeIO/rubix-os/nresty"
	"github.com/NubeIO/rubix-ui/backend/helpers/kb"
	"github.com/NubeIO/rubix-ui/backend/ttime"
	"github.com/go-resty/resty/v2"
)

type Snapshots struct {
	Name              string    `json:"name"`
	Size              int64     `json:"size"`
	SizeReadable      string    `json:"size_readable"`
	Description       string    `json:"description" get:"true" post:"true" patch:"true"`
	CreatedAt         time.Time `json:"created_at"`
	CreatedAtReadable string    `json:"created_at_readable"`
}

type CreateStatus string

const (
	Creating     CreateStatus = "Creating"
	Created      CreateStatus = "Created"
	CreateFailed CreateStatus = "Failed"
)

type SnapshotCreateLog struct {
	UUID        string       `json:"uuid" gorm:"primary_key" get:"true" delete:"true"`
	HostUUID    string       `json:"host_uuid" get:"true" post:"true"`
	Msg         string       `json:"msg" get:"true" post:"true" patch:"true"`
	Description string       `json:"description" get:"true" post:"true" patch:"true"`
	Status      CreateStatus `json:"status" get:"true" post:"true" patch:"true"`
	CreatedAt   time.Time    `json:"created_at" get:"true"`
}

type SnapshotCreateRequest struct {
	Description string `json:"description" get:"true" post:"true" patch:"true"`
}

type SnapshotRestoreRequest struct {
	FileName    string `json:"file" get:"true" post:"true" patch:"true"`
	Description string `json:"description" get:"true" post:"true" patch:"true"`
}

type RestoreStatus string

const (
	Restoring     RestoreStatus = "Restoring"
	Restored      RestoreStatus = "Restored"
	RestoreFailed RestoreStatus = "Failed"
)

type SnapshotRestoreLog struct {
	UUID        string        `json:"uuid" gorm:"primary_key" get:"true" delete:"true"`
	HostUUID    string        `json:"host_uuid" get:"true" post:"true" patch:"true"`
	Msg         string        `json:"msg" get:"true" post:"true" patch:"true"`
	Description string        `json:"description" get:"true" post:"true" patch:"true"`
	Status      RestoreStatus `json:"status" get:"true" post:"true" patch:"true"`
	CreatedAt   time.Time     `json:"created_at" get:"true"`
}

func (inst *Client) EdgeGetSnapshots(hostUUID string) ([]Snapshots, error) {
	url := fmt.Sprintf("/api/host/snapshots")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&[]Snapshots{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	var res []Snapshots
	var out []Snapshots
	res = *resp.Result().(*[]Snapshots)
	for _, snapshots := range res {
		snapshots.SizeReadable = kb.PrettyByteSize(int(snapshots.Size))
		snapshots.CreatedAtReadable = ttime.TimeSince(snapshots.CreatedAt)
		out = append(out, snapshots)
	}
	return out, nil
}

func (inst *Client) EdgeGetSnapshotsCreateLogs(hostUUID string) ([]SnapshotCreateLog, error) {
	url := fmt.Sprintf("/api/host/snapshots/create-logs")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&[]SnapshotCreateLog{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	var out []SnapshotCreateLog
	out = *resp.Result().(*[]SnapshotCreateLog)
	return out, nil
}

func (inst *Client) EdgeEditSnapshotLog(hostUUID string, uuid string, body *SnapshotCreateRequest) (*Message, error) {
	url := fmt.Sprintf("/api/host/snapshots/create-logs/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetBody(body).
		SetResult(&Message{}).
		Patch(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*Message), nil
}

func (inst *Client) EdgeDeleteSnapshotLog(hostUUID string, uuid string) (*Message, error) {
	url := fmt.Sprintf("/api/host/snapshots/create-logs/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&Message{}).
		Delete(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*Message), nil
}

func (inst *Client) EdgeGetSnapshotsRestoreLogs(hostUUID string) ([]SnapshotRestoreLog, error) {
	url := fmt.Sprintf("/api/host/snapshots/restore-logs")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&[]SnapshotRestoreLog{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	var out []SnapshotRestoreLog
	out = *resp.Result().(*[]SnapshotRestoreLog)
	return out, nil
}

func (inst *Client) EdgeEditSnapshotRestoreLog(hostUUID string, uuid string, body *SnapshotCreateRequest) (*Message, error) {
	url := fmt.Sprintf("/api/host/snapshots/restore-logs/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetBody(body).
		SetResult(&Message{}).
		Patch(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*Message), nil
}

func (inst *Client) EdgeDeleteSnapshotRestoreLog(hostUUID string, uuid string) (*Message, error) {
	url := fmt.Sprintf("/api/host/snapshots/restore-logs/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&Message{}).
		Delete(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*Message), nil
}

func (inst *Client) EdgeCreateSnapshot(hostUUID string, body *SnapshotCreateRequest) (*Message, error) {
	url := fmt.Sprintf("/api/host/snapshots/create")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetBody(body).
		SetResult(&Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*Message), nil
}

func (inst *Client) EdgeEditSnapshot(hostUUID, fileName string, body *SnapshotCreateRequest) (*Message, error) {
	url := fmt.Sprintf("/api/host/snapshots/file/%s", fileName)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetBody(body).
		SetResult(&Message{}).
		Patch(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*Message), nil
}

func (inst *Client) EdgeDeleteSnapshot(hostUUID, fileName string) (*Message, error) {
	url := fmt.Sprintf("/api/host/snapshots?file=%s", fileName)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&Message{}).
		Delete(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*Message), nil
}

func (inst *Client) EdgeUploadSnapshot(hostUUID, description, path string) (*Message, error) {
	url := fmt.Sprintf("/api/host/snapshots/upload")

	file, err := os.Open(path)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// set query parameters
	queryParams := map[string]string{
		"description": description,
	}

	// send multipart form data request
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetQueryParams(queryParams).
		SetHeader("Content-Type", "multipart/form-data").
		SetMultipartFields(&resty.MultipartField{
			Param:       "file",
			FileName:    filepath.Base(path),
			Reader:      file,
			ContentType: "application/zip",
		}).
		SetResult(&Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	log.Debugf("Response body: %s", resp.Body())
	return resp.Result().(*Message), nil
}

func (inst *Client) EdgeDownloadSnapshot(hostUUID, fileName, destination string) (*Message, error) {
	url := fmt.Sprintf("/api/host/snapshots/download?file=%s", fileName)
	fmt.Println(url)

	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		Post(url))
	if err != nil {
		return nil, err
	}
	SaveFile(resp.Body(), destination)
	return &Message{
		Message: fmt.Sprintf("File downloaded successfully to %s", destination),
	}, nil
}

func (inst *Client) EdgeRestoreSnapshot(hostUUID string, body *SnapshotRestoreRequest) (*Message, error) {
	url := fmt.Sprintf("/api/host/snapshots/restore")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetBody(body).
		SetResult(&Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*Message), nil
}

func SaveFile(data []byte, path string) error {
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
