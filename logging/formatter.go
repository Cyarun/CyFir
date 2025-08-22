package logging

import (
	"github.com/sirupsen/logrus"
	"github.com/Cyarun/CyFir/utils"
)

type JSONFormatter struct {
	*logrus.JSONFormatter
}

func (self *JSONFormatter) Format(e *logrus.Entry) ([]byte, error) {
	e.Time = utils.GetTime().Now().UTC()

	result, err := self.JSONFormatter.Format(e)
	return result, err
}
