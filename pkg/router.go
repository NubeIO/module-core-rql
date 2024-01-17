package pkg

import (
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/lib-module-go/nmodule"
	"github.com/NubeIO/lib-module-go/router"
	"net/http"
)

var route *router.Router

func InitRouter() {
	route = router.NewRouter()

	route.Handle(nhttp.GET, "/api/rules", GetRules)
	route.Handle(nhttp.GET, "/api/rules/:uuid", GetRule)
	route.Handle(nhttp.GET, "/api/vars", GetVars)
	route.Handle(nhttp.GET, "/api/vars/:uuid", GetVar)

	route.Handle(nhttp.POST, "/api/rules", AddRule)
	route.Handle(nhttp.POST, "/api/rules/:uuid/run", RunRule)
	route.Handle(nhttp.POST, "/api/rules/dry", RunDry)
	route.Handle(nhttp.POST, "/api/vars", AddVar)

	route.Handle(nhttp.PATCH, "/api/rules/:uuid", UpdateRule)
	route.Handle(nhttp.PATCH, "/api/vars/:uuid", UpdateVar)

	route.Handle(nhttp.DELETE, "/api/rules/:uuid", DeleteRule)
	route.Handle(nhttp.DELETE, "/api/vars/:uuid", DeleteVar)
}

func (m *Module) CallModule(method nhttp.Method, urlString string, headers http.Header, body []byte) ([]byte, error) {
	err := m.check()
	if err != nil {
		return nil, err
	}
	module := nmodule.Module(m)
	return route.CallHandler(&module, method, urlString, headers, body)
}

func GetRules(m *nmodule.Module, r *router.Request) ([]byte, error) {
	return (*m).(*Module).SelectAllRules()
}

func GetRule(m *nmodule.Module, r *router.Request) ([]byte, error) {
	return (*m).(*Module).ReuseRuleRun(r.Body, r.PathParams["uuid"])
}

func GetVars(m *nmodule.Module, r *router.Request) ([]byte, error) {
	return (*m).(*Module).SelectAllVariables()
}

func GetVar(m *nmodule.Module, r *router.Request) ([]byte, error) {
	return (*m).(*Module).SelectVariable(r.PathParams["uuid"])
}

func AddRule(m *nmodule.Module, r *router.Request) ([]byte, error) {
	return (*m).(*Module).AddRule(r.Body)
}

func RunRule(m *nmodule.Module, r *router.Request) ([]byte, error) {
	return (*m).(*Module).ReuseRuleRun(r.Body, r.PathParams["uuid"])
}

func RunDry(m *nmodule.Module, r *router.Request) ([]byte, error) {
	return (*m).(*Module).Dry(r.Body)
}

func AddVar(m *nmodule.Module, r *router.Request) ([]byte, error) {
	return (*m).(*Module).AddVariable(r.Body)
}

func UpdateRule(m *nmodule.Module, r *router.Request) ([]byte, error) {
	return (*m).(*Module).UpdateRule(r.PathParams["uuid"], r.Body)
}

func UpdateVar(m *nmodule.Module, r *router.Request) ([]byte, error) {
	return (*m).(*Module).UpdateVariable(r.Body, r.PathParams["uuid"])
}

func DeleteRule(m *nmodule.Module, r *router.Request) ([]byte, error) {
	return (*m).(*Module).DeleteRule(r.PathParams["uuid"])
}

func DeleteVar(m *nmodule.Module, r *router.Request) ([]byte, error) {
	return (*m).(*Module).DeleteVariable(r.PathParams["uuid"])
}
