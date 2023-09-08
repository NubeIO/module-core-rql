package apirules

import (
	"github.com/NubeIO/module-core-rql/storage"
	"github.com/go-resty/resty/v2"
)

func (inst *RQL) GetScripts() any {
	client := resty.New()
	url := "http://0.0.0.0:1666/api/rules"
	resp, err := client.R().
		SetResult(&[]storage.RQLRule{}).
		Get(url)
	if err != nil {
		return err
	}
	var out []storage.RQLRule
	out = *resp.Result().(*[]storage.RQLRule)
	return out
}
