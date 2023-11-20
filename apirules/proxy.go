package apirules

/*
	let apiGet = RQL.Get("rc", "points?with_priority=true");
	let points = RQL.ToPoints(apiGet.Body());

	let resp = {
	  status: apiGet.Status(),
	  points: points,
	  count: points.length,
	};

	RQL.Result = resp;

*/

/*
let apiGet = RQL.Get("rc", "system/device");
let str = JSON.parse(apiGet.String());
RQL.Result = str.global_uuid;
*/

func (inst *RQL) Get(hostIDName, path string) any {
	resp, err := cli.ProxyGET(hostIDName, path)
	if err != nil {
		return err
	}
	return resp
}

func (inst *RQL) Post(hostIDName, path string, body any) any {
	resp, err := cli.ProxyPOST(hostIDName, path, body)
	if err != nil {
		return err
	}
	return resp
}

func (inst *RQL) Patch(hostIDName, path string, body any) any {
	resp, err := cli.ProxyPATCH(hostIDName, path, body)
	if err != nil {
		return err
	}
	return resp
}

func (inst *RQL) Put(hostIDName, path string, body any) any {
	resp, err := cli.ProxyPUT(hostIDName, path, body)
	if err != nil {
		return err
	}
	return resp
}

func (inst *RQL) Delete(hostIDName, path string) any {
	resp, err := cli.ProxyDELETE(hostIDName, path)
	if err != nil {
		return err
	}
	return resp
}
