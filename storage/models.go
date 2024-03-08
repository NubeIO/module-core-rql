package storage

import (
	"time"
)

type Result struct {
	Result    interface{} `json:"result"`
	Timestamp string      `json:"timestamp"`
	Time      time.Time   `json:"time"`
}

type RQLRule struct {
	UUID              string   `json:"uuid"`
	Name              string   `json:"name"`
	LatestRunDate     string   `json:"latest_run_date"`
	Script            string   `json:"script"`
	Schedule          string   `json:"schedule"`
	Enable            bool     `json:"enable"`
	ResultStorageSize int      `json:"result_storage_size"`
	Result            []Result `json:"result"`
}

func (c RQLRule) ID() (jsonField string, value interface{}) {
	value = c.UUID
	jsonField = "uuid"
	return
}

type RQLVariables struct {
	UUID       string `json:"uuid"`
	ScriptUUID string `json:"script_uuid"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Value      any    `json:"value"`
	Required   bool   `json:"required"`
}

func (c RQLVariables) ID() (jsonField string, value interface{}) {
	value = c.UUID
	jsonField = "uuid"
	return
}
