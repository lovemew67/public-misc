package cornerstone

import (
	"github.com/sirupsen/logrus"
)

func init() {
	// logger
	logrusLogger = logrus.New()
	logrusLogger.SetOutput(defaultOutput)
	logrusLogger.SetLevel(defaultLevel)
	logrusLogger.SetFormatter(&cornerstoneJSONFormatter{})
}
