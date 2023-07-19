package pkg

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/bugs"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	log "github.com/sirupsen/logrus"
	"time"
)

var name = "module-core-rql"

const Permission = 0755

func (inst *Module) networkUpdateSuccess(uuid string) error {
	var network model.Network
	network.InFault = false
	network.MessageLevel = model.MessageLevel.Info
	network.MessageCode = model.CommonFaultCode.Ok
	network.Message = model.CommonFaultMessage.NetworkMessage
	network.LastOk = time.Now().UTC()
	err := inst.grpcMarshaller.UpdateNetworkErrors(uuid, &network)
	if err != nil {
		log.Error(bugs.DebugPrint(name, inst.networkUpdateSuccess, err))
	}
	return err
}

func (inst *Module) networkUpdateErr(uuid, port string, e error) error {
	var network model.Network
	network.InFault = true
	network.MessageLevel = model.MessageLevel.Fail
	network.MessageCode = model.CommonFaultCode.NetworkError
	network.Message = fmt.Sprintf(" port: %s message: %s", port, e.Error())
	network.LastFail = time.Now().UTC()
	err := inst.grpcMarshaller.UpdateNetworkErrors(uuid, &network)
	if err != nil {
		log.Error(bugs.DebugPrint(name, inst.networkUpdateErr, err))
	}
	return err
}

func (inst *Module) handleSerialPayload(data string) {}
