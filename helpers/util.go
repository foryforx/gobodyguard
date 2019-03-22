package helpers

import (
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

// NewUUID will generate a new UUID and return back the value as string
func NewUUID() string {
	UUID := uuid.NewV4().String()
	log.Debugln("New UUID generated", UUID)
	return UUID
}

// HandleError will handle error in all places and log the error accordingly
func HandleError(err error, message string, args ...interface{}) {
	if err != nil {
		log.Errorln(err)
		errors.Wrapf(err, message, args...)
	}
}
