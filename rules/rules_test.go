package rules

import (
	"fmt"
	"github.com/NubeIO/module-core-rql/storage"
	"testing"
)

func TestNewRuleEngine(t *testing.T) {

	script := `
	let a = RQL.in1+10
	RQL.Result = a+111111
`

	rule := &storage.RQLRule{
		UUID:              "",
		Name:              "test",
		LatestRunDate:     "",
		Script:            script,
		Schedule:          "",
		Enable:            true,
		ResultStorageSize: 0,
		Result:            nil,
	}
	arg := map[string]interface{}{"in1": 22.2, "in2": 23, "in3": 23}
	props := PropertiesMap{
		"RQL": arg,
	}

	r := NewRuleEngine()

	err := r.AddRule(rule, props)

	if err != nil {
		fmt.Println("add", err)
		return
	}

	res, err := r.ExecuteByName(rule.Name, true)

	if err != nil {
		fmt.Println("run", err)
		return
	}

	fmt.Println(res.String())

}
