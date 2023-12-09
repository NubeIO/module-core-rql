package pkg

import (
	"github.com/NubeIO/lib-module-go/http"
	"github.com/NubeIO/lib-module-go/router"
	"github.com/NubeIO/lib-module-go/shared"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/nargs"
)

var route *router.Router

func InitRouter() {
	route = router.NewRouter()

	route.Handle(http.GET, "/api/rules", GetRules)
	route.Handle(http.GET, "/api/rules/:uuid", GetRule)
	route.Handle(http.GET, "/api/vars", GetVars)
	route.Handle(http.GET, "/api/vars/:uuid", GetVar)

	route.Handle(http.POST, "/api/rules", AddRule)
	route.Handle(http.POST, "/api/rules/:uuid/run", RunRule)
	route.Handle(http.POST, "/api/rules/dry", RunDry)
	route.Handle(http.POST, "/api/vars", AddVar)

	route.Handle(http.PATCH, "/api/rules/:uuid", UpdateRule)
	route.Handle(http.PATCH, "/api/vars/:uuid", UpdateVar)

	route.Handle(http.DELETE, "/api/rules/:uuid", DeleteRule)
	route.Handle(http.DELETE, "/api/vars/:uuid", DeleteVar)
}

func (m *Module) CallModule(method http.Method, api string, args nargs.Args, body []byte) ([]byte, error) {
	err := m.check()
	if err != nil {
		return nil, err
	}
	module := shared.Module(m)
	return route.CallHandler(&module, method, api, args, body)
}

func GetRules(m *shared.Module, params map[string]string, args nargs.Args, body []byte) ([]byte, error) {
	return (*m).(*Module).SelectAllRules()
}

func GetRule(m *shared.Module, params map[string]string, args nargs.Args, body []byte) ([]byte, error) {
	return (*m).(*Module).ReuseRuleRun(body, params["uuid"])
}

func GetVars(m *shared.Module, params map[string]string, args nargs.Args, body []byte) ([]byte, error) {
	return (*m).(*Module).SelectAllVariables()
}

func GetVar(m *shared.Module, params map[string]string, args nargs.Args, body []byte) ([]byte, error) {
	return (*m).(*Module).SelectVariable(params["uuid"])
}

func AddRule(m *shared.Module, params map[string]string, args nargs.Args, body []byte) ([]byte, error) {
	return (*m).(*Module).AddRule(body)
}

func RunRule(m *shared.Module, params map[string]string, args nargs.Args, body []byte) ([]byte, error) {
	return (*m).(*Module).ReuseRuleRun(body, params["uuid"])
}

func RunDry(m *shared.Module, params map[string]string, args nargs.Args, body []byte) ([]byte, error) {
	return (*m).(*Module).Dry(body)
}

func AddVar(m *shared.Module, params map[string]string, args nargs.Args, body []byte) ([]byte, error) {
	return (*m).(*Module).AddVariable(body)
}

func UpdateRule(m *shared.Module, params map[string]string, args nargs.Args, body []byte) ([]byte, error) {
	return (*m).(*Module).UpdateRule(params["uuid"], body)
}

func UpdateVar(m *shared.Module, params map[string]string, args nargs.Args, body []byte) ([]byte, error) {
	return (*m).(*Module).UpdateVariable(body, params["uuid"])
}

func DeleteRule(m *shared.Module, params map[string]string, args nargs.Args, body []byte) ([]byte, error) {
	return (*m).(*Module).DeleteRule(params["uuid"])
}

func DeleteVar(m *shared.Module, params map[string]string, args nargs.Args, body []byte) ([]byte, error) {
	return (*m).(*Module).DeleteVariable(params["uuid"])
}
