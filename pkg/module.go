package pkg

import (
	"github.com/NubeIO/lib-module-go/shared"
	"github.com/NubeIO/module-core-rql/apirules"
	"github.com/NubeIO/module-core-rql/rules"
	"github.com/NubeIO/module-core-rql/storage"
	"github.com/patrickmn/go-cache"
	"time"
)

type Module struct {
	dbHelper        shared.DBHelper
	moduleName      string
	grpcMarshaller  shared.Marshaller
	config          *Config
	store           *cache.Cache
	Rules           *rules.RuleEngine
	Client          *apirules.RQL
	Props           rules.PropertiesMap
	Storage         storage.IStorage
	ErrorOnDB       bool
	moduleDirectory string
	pluginIsEnabled bool
}

func (m *Module) Init(dbHelper shared.DBHelper, moduleName string) error {
	grpcMarshaller := shared.GRPCMarshaller{DbHelper: dbHelper}
	m.dbHelper = dbHelper
	m.moduleName = moduleName
	m.grpcMarshaller = &grpcMarshaller
	m.store = cache.New(5*time.Minute, 10*time.Minute)
	dir, err := m.grpcMarshaller.CreateModuleDir(moduleName)
	if err != nil {
		return err
	}
	m.moduleDirectory = *dir
	return nil
}

func (m *Module) GetInfo() (*shared.Info, error) {
	return &shared.Info{
		Name:       name,
		Author:     "Nube iO",
		Website:    "https://nube-io.com",
		License:    "N/A",
		HasNetwork: false,
	}, nil
}
