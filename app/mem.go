package app

import (
	"sync"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// // Access represents the mapping between principal and list of resources
// // Here key is principal UUID and value is Resource slice UUID
// type Access map[string][]string

// // OperationAccess represents the mapping between Operation and Access
// // Here key is UUID and value is Access
// type OperationAccess map[string]Access

// MemLoad represents data in memory for permission
type MemLoad struct {
	OperationAccess map[string]map[string][]string
}

var onceMemLoad sync.Once
var memInst *MemLoad

// GetMemData will get us the Data for permission in memory
func GetMemData() *MemLoad {
	onceMemLoad.Do(func() {
		log.Debugln("Loading loadAccess")
		d, err := loadAccess()
		if err != nil {
			log.Errorln(err)
			panic(err)
		}
		memInst = &MemLoad{OperationAccess: d}
	})
	return memInst
}

// LoadAccess will load all the Access from DB to memory during startup
func loadAccess() (map[string](map[string][]string), error) {
	var access = make(map[string](map[string][]string))
	db := GetDBInstance()
	log.Debugln(db.DB)
	authRepo := NewAuthRepository(db.DB)
	access, err := authRepo.LoadAccess()
	if err != nil {
		return access, errors.Wrapf(err, "LoadAccess failed")
	}
	return access, nil
}
