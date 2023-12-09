package rubixoscli

import (
	"errors"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/internaltoken"
	"github.com/NubeIO/rubix-os/installer"
	"github.com/NubeIO/rubix-os/nresty"
	"sync"

	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

var (
	mutex   = &sync.RWMutex{}
	clients = map[string]*Client{}
)

// TODO: WARNING: DON'T OVERRIDE FROM COPYING

type Client struct {
	Rest      *resty.Client
	Installer *installer.Installer
	Ip        string `json:"ip"`
	Port      int    `json:"port"`
	HTTPS     bool   `json:"https"`
}

type ResponseBody struct {
	Response ResponseCommon `json:"response"`
	Status   string         `json:"status"`
	Count    string         `json:"count"`
}

type ResponseCommon struct {
	UUID string `json:"uuid"`
}

func buildUrl(overrideUrl ...string) string {
	if len(overrideUrl) > 0 {
		if overrideUrl[0] != "" {
			return overrideUrl[0]
		}
	}
	return ""
}

func New(cli *Client, installer *installer.Installer) *Client {
	mutex.Lock()
	defer mutex.Unlock()
	if cli == nil {
		log.Fatal("rubix-os client cli can not be empty")
		return nil
	}
	baseURL := getBaseUrl(cli)
	if client, found := clients[baseURL]; found {
		return client
	}
	cli.Rest = resty.New()
	cli.Installer = installer
	cli.Rest.SetBaseURL(baseURL)
	cli.SetTokenHeader()
	clients[baseURL] = cli
	return cli
}

func getBaseUrl(cli *Client) string {
	if cli.Ip == "" {
		cli.Ip = "0.0.0.0"
	}
	if cli.Port == 0 {
		cli.Port = 1659
	}
	var baseURL string
	if cli.HTTPS {
		baseURL = fmt.Sprintf("https://%s:%d/ros", cli.Ip, cli.Port)
	} else {
		baseURL = fmt.Sprintf("http://%s:%d/ros", cli.Ip, cli.Port)
	}
	return baseURL
}

func (inst *Client) SetTokenHeader() *Client {
	token := internaltoken.GetInternalToken(true)
	inst.Rest.Header.Set("Authorization", token)
	return inst
}

type Path struct {
	Path string
}

var Paths = struct {
	Hosts        Path
	Ping         Path
	Groups       Path
	Locations    Path
	Edge         Path
	Apps         Path
	Tasks        Path
	Transactions Path
	System       Path
	Networking   Path
	Wires        Path
	Alerts       Path
}{
	Hosts:        Path{Path: "/api/hosts"},
	Ping:         Path{Path: "/api/system/ping"},
	Groups:       Path{Path: "/api/groups"},
	Locations:    Path{Path: "/api/locations"},
	Edge:         Path{Path: "/api/edgeapi"},
	Apps:         Path{Path: "/api/edgeapi/apps"},
	Tasks:        Path{Path: "/api/tasks"},
	Transactions: Path{Path: "/api/transactions"},
	System:       Path{Path: "/api/system"},
	Networking:   Path{Path: "/api/networking"},
	Wires:        Path{Path: "/api/wires"},
	Alerts:       Path{Path: "/api/alerts"},
}

type Response struct {
	StatusCode int         `json:"code"`
	Message    interface{} `json:"message"`
	resty      *resty.Response
}

func (response Response) buildResponse(resp *resty.Response, err error) *Response {
	response.StatusCode = resp.StatusCode()
	response.resty = resp
	if resp.IsError() {
		response.Message = resp.Error()
	}
	if resp.StatusCode() == 0 {
		response.Message = "server is unreachable"
		response.StatusCode = 503
	}
	return &response
}

func (inst *Client) PingRubixOs() (bool, error) {
	path := fmt.Sprintf("/api/system/time")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		Get(path))
	if err != nil {
		return false, err
	}
	if resp.StatusCode() > 300 {
		return false, errors.New("failed to ping connection")
	} else {
		return true, nil
	}
}

func (inst *Client) PingEdge(hostIDName string) (bool, error) {
	time, err := inst.EdgeSystemTime(hostIDName)
	if err != nil {
		return false, err
	}
	if time == nil {
		return false, errors.New("failed to ping connection")
	} else {
		return true, nil
	}
}
