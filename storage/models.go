package storage

import (
	"time"
)

type Result struct {
	Result    interface{} `json:"result"`
	Timestamp time.Time   `json:"timestamp"`
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

type RQLVariables struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Variable any    `json:"variable"`
	Password string `json:"password"`
}
