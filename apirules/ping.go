package apirules

import (
	"github.com/NubeDev/flow-eng/helpers"
)

type PingResult struct {
	Ip string `json:"ip"`
	Ok bool   `json:"ok"`
}

type PingResponse struct {
	Result any
	Error  string
}

// Ping ping an list of IP address eg: ["192.168.15.1", "192.168.15.2"]
// will return []PingResult
func (inst *RQL) Ping(ipList []string) *PingResponse {
	var r PingResult
	var out []PingResult
	for _, ip := range ipList {
		ok := helpers.CommandPing(ip)
		r.Ip = ip
		r.Ok = ok
		out = append(out, r)
	}
	return &PingResponse{
		Result: out,
		Error:  "",
	}
}
