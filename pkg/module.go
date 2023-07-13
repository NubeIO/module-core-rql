package pkg

import (
	"github.com/NubeIO/module-core-rql/apirules"
	"github.com/NubeIO/module-core-rql/rules"
	"github.com/NubeIO/module-core-rql/storage"
	"github.com/NubeIO/rubix-os/module/shared"
	"github.com/patrickmn/go-cache"
	"time"
)

type Module struct {
	dbHelper       shared.DBHelper
	moduleName     string
	grpcMarshaller shared.Marshaller
	config         *Config
	store          *cache.Cache

	Rules   *rules.RuleEngine
	Client  *apirules.Client
	Props   rules.PropertiesMap
	Storage storage.IStorage
}

func (m *Module) Init(dbHelper shared.DBHelper, moduleName string) error {
	grpcMarshaller := shared.GRPCMarshaller{DbHelper: dbHelper}
	m.dbHelper = dbHelper
	m.moduleName = moduleName
	m.grpcMarshaller = &grpcMarshaller
	m.store = cache.New(5*time.Minute, 10*time.Minute)
	return nil
}

func (m *Module) GetInfo() (*shared.Info, error) {
	return &shared.Info{
		Name:       name,
		Author:     "Nube iO",
		Website:    "https://nube-io.com",
		License:    "N/A",
		HasNetwork: true,
	}, nil
}
