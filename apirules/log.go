package apirules

import log "github.com/sirupsen/logrus"

func (inst *RQL) Log(body any) {
	log.Infof("%#v", body)
}
