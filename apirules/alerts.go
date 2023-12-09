package apirules

import (
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	"github.com/mitchellh/mapstructure"
)

/*
let body = {
  hostUUID: "hos_c8ab2fe07e7a413c",
  entityType: "device",
  type: "ping",
  status: "active",
  severity: "crucial",
};

let alert = RQL.AddAlert("hos_c8ab2fe07e7a413c", body);

RQL.Result = alert.CreatedAt;
*/

func alertBody(body any) (*model.Alert, error) {
	result := &model.Alert{}
	err := mapstructure.Decode(body, &result)
	return result, err
}

func (inst *RQL) GetAlerts() any {
	resp, err := cli.GetAlerts([]string{"active", "acknowledged", "closed"})
	if err != nil {
		return err
	}
	return resp
}

type alertsByHost struct {
	HostUUID               string
	HostName               string
	AlertsActive           []model.Alert
	AlertsActiveCount      int
	AlertsAcknowledge      []model.Alert
	AlertsAcknowledgeCount int
	AlertsClosed           []model.Alert
	AlertsClosedCount      int
	TotalAlertsCount       int
	Error                  error
}

// 	GetAlertsByHost filter active | acknowledge | closed
func (inst *RQL) GetAlertsByHost() any {
	resp, err := cli.GetHosts()
	if err != nil {
		return err
	}
	var out []alertsByHost
	var alertsActive []model.Alert
	var alertsAcknowledge []model.Alert
	var alertsClosed []model.Alert
	var active = "active"
	var acknowledge = "acknowledged"
	var closed = "closed"

	for _, host := range resp {
		alerts, err := cli.GetAlertsByHost(host.UUID, []string{"active", "acknowledged", "closed"})
		for _, alert := range alerts {
			if alert.Status == active {
				alertsActive = append(alertsActive, alert)
			}
			if alert.Status == acknowledge {
				alertsAcknowledge = append(alertsAcknowledge, alert)
			}
			if alert.Status == closed {
				alertsClosed = append(alertsClosed, alert)
			}
		}
		newHost := alertsByHost{
			HostUUID:               host.UUID,
			HostName:               host.Name,
			AlertsActive:           alertsActive,
			AlertsActiveCount:      len(alertsActive),
			AlertsAcknowledge:      alertsAcknowledge,
			AlertsAcknowledgeCount: len(alertsAcknowledge),
			AlertsClosed:           alertsClosed,
			AlertsClosedCount:      len(alertsClosed),
			TotalAlertsCount:       len(alerts),
			Error:                  err,
		}
		out = append(out, newHost)

	}

	return out
}

func (inst *RQL) AddAlert(hostIDName string, body any) any {
	b, err := alertBody(body)
	if err != nil {
		return err
	}
	resp, err := cli.AddAlert(hostIDName, b)
	if err != nil {
		return err
	}
	return resp
}

func (inst *RQL) UpdateAlertStatus(hostIDName, uuid, status string) any {
	resp, err := cli.UpdateAlertStatus(hostIDName, uuid, status)
	if err != nil {
		return err
	}
	return resp
}
