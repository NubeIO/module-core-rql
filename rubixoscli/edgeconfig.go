package rubixoscli

import (
	"errors"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/dto"
	"github.com/NubeIO/rubix-os/nresty"
	"github.com/NubeIO/rubix-ui/backend/constants"
	"gopkg.in/yaml.v3"
	"reflect"
)

type BacnetClient struct { // this is for bacnet-master
	Debug    bool     `json:"debug" yaml:"debug"`
	Enable   bool     `json:"enable" yaml:"enable"`
	Commands []string `json:"commands" yaml:"commands"`
	Tokens   []string `json:"tokens" yaml:"tokens"`
}

type Mqtt struct {
	BrokerIp          string `json:"broker_ip"  yaml:"broker_ip"`
	BrokerPort        int    `json:"broker_port"  yaml:"broker_port"`
	Debug             bool   `json:"debug" yaml:"debug"`
	Enable            bool   `json:"enable" yaml:"enable"`
	WriteViaSubscribe bool   `json:"write_via_subscribe" yaml:"write_via_subscribe"`
	RetryEnable       bool   `json:"retry_enable" yaml:"retry_enable"`
	RetryLimit        int    `json:"retry_limit" yaml:"retry_limit"`
	RetryInterval     int    `json:"retry_interval" yaml:"retry_interval"`
}

type ConfigBACnetServer struct {
	ServerName   string       `json:"server_name" yaml:"server_name"`
	DeviceId     int          `json:"device_id" yaml:"device_id"`
	Port         int          `json:"port" yaml:"port"`
	Iface        string       `json:"iface" yaml:"iface"`
	BiMax        int          `json:"bi_max" yaml:"bi_max"`
	BoMax        int          `json:"bo_max" yaml:"bo_max"`
	BvMax        int          `json:"bv_max" yaml:"bv_max"`
	AiMax        int          `json:"ai_max" yaml:"ai_max"`
	AoMax        int          `json:"ao_max" yaml:"ao_max"`
	AvMax        int          `json:"av_max" yaml:"av_max"`
	Objects      []string     `json:"objects" yaml:"objects"`
	Properties   []string     `json:"properties" yaml:"properties"`
	Mqtt         Mqtt         `json:"mqtt" yaml:"mqtt"`
	BacnetClient BacnetClient `json:"bacnet_client" yaml:"bacnet_client"`
}

// EdgeWriteConfig replace the config file of a nube app
func (inst *Client) EdgeWriteConfig(hostUUID, appName string) (*Message, error) {
	pushConfig := false
	var writeConfig dto.HostConfig
	if appName == constants.BacnetServerDriver {
		pushConfig = true
		resp, connectionErr, requestErr := inst.EdgeReadConfig(hostUUID, appName, constants.ConfigYml)
		var config ConfigBACnetServer
		if connectionErr != nil {
			return nil, connectionErr
		}
		if requestErr != nil {
			config = ConfigBACnetServer{}
		}
		if resp != nil {
			err := yaml.Unmarshal(resp.Data, &config)
			if err != nil {
				return nil, err
			}
		}
		writeConfig = dto.HostConfig{
			AppName:    constants.BacnetServerDriver,
			Body:       inst.convertConfigToDynamicMap(inst.defaultWrapperBACnetConfig(config)),
			ConfigName: constants.ConfigYml,
		}
	}
	if pushConfig {
		url := fmt.Sprintf("/api/host/config")
		resp, err := nresty.FormatRestyResponse(inst.Rest.R().
			SetHeader("X-Host", hostUUID).
			SetResult(&Message{}).
			SetBody(writeConfig).
			Post(url))
		if err != nil {
			return nil, err
		}
		return resp.Result().(*Message), nil
	}
	return nil, nil
}

func (inst *Client) EdgeReadConfig(hostUUID, appName, configName string) (*dto.HostConfigResponse, error, error) {
	url := fmt.Sprintf("/api/host/config?app_name=%s&config_name=%s", appName, configName)
	resp, connectionError, requestErr := nresty.FormatRestyV2Response(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&dto.HostConfigResponse{}).
		Get(url))
	if connectionError != nil || requestErr != nil {
		return nil, connectionError, requestErr
	}
	return resp.Result().(*dto.HostConfigResponse), nil, nil
}

func (inst *Client) BACnetWriteConfig(hostUUID, appName string, config ConfigBACnetServer) (*Message, error) {
	var writeConfig dto.HostConfig
	if appName == constants.BacnetServerDriver {
		writeConfig = dto.HostConfig{
			AppName:    constants.BacnetServerDriver,
			Body:       inst.convertConfigToDynamicMap(inst.defaultWrapperBACnetConfig(config)),
			ConfigName: constants.ConfigYml,
		}
	} else {
		return nil, errors.New(fmt.Sprintf("app name must be bacnet: %s", appName))
	}
	url := fmt.Sprintf("/api/host/config")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&Message{}).
		SetBody(writeConfig).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*Message), nil
}

func (inst *Client) defaultWrapperBACnetConfig(config ConfigBACnetServer) ConfigBACnetServer {
	if config.ServerName == "" {
		config.ServerName = "Nube IO"
	}
	if config.DeviceId == 0 {
		config.DeviceId = 2508
	}
	if config.Iface == "" {
		config.Iface = "eth0"
	}
	if config.Port == 0 {
		config.Port = 47808
	}

	config.Objects = []string{"ai", "av", "ao", "bi", "bo", "bv"}
	config.Properties = []string{"name", "pv", "pri"}

	// bacnet-master
	config.BacnetClient.Commands = []string{"whois", "read_value", "write_value", "pics"}
	config.BacnetClient.Tokens = []string{"txn_source", "txn_number"}
	config.BacnetClient.Enable = true

	if config.Mqtt.BrokerIp == "" {
		config.Mqtt.BrokerIp = "127.0.0.1"
	}
	if config.Mqtt.BrokerPort == 0 {
		config.Mqtt.BrokerPort = 1883
	}
	if config.Mqtt.RetryLimit == 0 {
		config.Mqtt.RetryLimit = 5
	}
	if config.Mqtt.RetryInterval == 0 {
		config.Mqtt.RetryInterval = 10
	}
	config.Mqtt.Debug = false
	config.Mqtt.Enable = true // will make the bacnet-server work
	config.Mqtt.RetryEnable = true
	config.Mqtt.WriteViaSubscribe = true // if not enabled the point point values will not update over MQTT

	return config
}

func (inst *Client) convertConfigToDynamicMap(config interface{}) dto.DynamicMap {
	dynamicMap := make(dto.DynamicMap)

	// Use reflection to iterate through the struct fields and populate the DynamicMap
	v := reflect.ValueOf(config)
	t := reflect.TypeOf(config)
	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Tag.Get("json")
		if fieldName == "" {
			continue
		}
		fieldValue := v.Field(i).Interface()
		dynamicMap[fieldName] = fieldValue
	}

	return dynamicMap
}
