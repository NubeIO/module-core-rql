package apirules

import (
	"github.com/go-resty/resty/v2"
)

/*
let apiGet = RQL.Get("rc", "points");
let points = RQL.ToPoints(apiGet);

let resp = {
  code: apiGet.Status(),
  points: points,
};

RQL.Result = resp;
*/

func (inst *RQL) Get(hostIDName, path string) *resty.Response {
	resp, _ := cli.ProxyGET(hostIDName, path)
	return resp
}

func (inst *RQL) Post(hostIDName, path string, body any) *resty.Response {
	resp, _ := cli.ProxyPOST(hostIDName, path, body)
	return resp
}

func (inst *RQL) Patch(hostIDName, path string, body any) *resty.Response {
	resp, _ := cli.ProxyPATCH(hostIDName, path, body)
	return resp
}

func (inst *RQL) Put(hostIDName, path string, body any) *resty.Response {
	resp, _ := cli.ProxyPUT(hostIDName, path, body)
	return resp
}

func (inst *RQL) Delete(hostIDName, path string) *resty.Response {
	resp, _ := cli.ProxyDELETE(hostIDName, path)
	return resp
}
