package apirules

import (
	"github.com/NubeDev/flow-eng/helpers"
)

// PingResult represents the result of a ping operation
// It contains information about the IP address and whether the ping was successful or not.
type PingResult struct {
	Ip string `json:"ip"`
	Ok bool   `json:"ok"`
}

// Ping performs a ping operation on each IP address in the provided list.
// It returns a list of PingResult objects containing the IP address and the ping result.
func (inst *RQL) Ping(ipList []string) []PingResult {
	var r PingResult
	var out []PingResult
	for _, ip := range ipList {
		ok := helpers.CommandPing(ip)
		r.Ip = ip
		r.Ok = ok
		out = append(out, r)
	}
	return out
}
