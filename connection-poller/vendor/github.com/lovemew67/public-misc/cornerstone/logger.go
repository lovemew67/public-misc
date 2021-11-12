package cornerstone

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	buildInFieldNumber = 5
	timestampFormat    = time.RFC3339Nano
)

const (
	logKeyAppName    = "app_name"
	logKeyAppVersion = "app_version"
	logKeyTimestamp  = "timestamp"
	logKeyMessage    = "message"
	logKeyLevel      = "level"
)

var (
	defaultOutput = os.Stdout
	defaultLevel  = logrus.DebugLevel
)

var (
	logrusLogger *logrus.Logger
)

// logger

func Trace(ctx Context, v ...interface{}) {
	if logrusLogger.Level >= logrus.TraceLevel {
		getLogrusEntry(ctx).Trace(v...)
	}
}

func Debug(ctx Context, v ...interface{}) {
	if logrusLogger.Level >= logrus.DebugLevel {
		getLogrusEntry(ctx).Debug(v...)
	}
}

func Info(ctx Context, v ...interface{}) {
	if logrusLogger.Level >= logrus.InfoLevel {
		getLogrusEntry(ctx).Info(v...)
	}
}

func Error(ctx Context, v ...interface{}) {
	if logrusLogger.Level >= logrus.ErrorLevel {
		getLogrusEntry(ctx).Error(v...)
	}
}

func Warn(ctx Context, v ...interface{}) {
	if logrusLogger.Level >= logrus.WarnLevel {
		getLogrusEntry(ctx).Warn(v...)
	}
}

func Panic(ctx Context, v ...interface{}) {
	if logrusLogger.Level >= logrus.PanicLevel {
		getLogrusEntry(ctx).Panic(v...)
	}
}

func Tracef(ctx Context, f string, v ...interface{}) {
	if logrusLogger.Level >= logrus.TraceLevel {
		getLogrusEntry(ctx).Tracef(f, v...)
	}
}

func Debugf(ctx Context, f string, v ...interface{}) {
	if logrusLogger.Level >= logrus.DebugLevel {
		getLogrusEntry(ctx).Debugf(f, v...)
	}
}

func Infof(ctx Context, f string, v ...interface{}) {
	if logrusLogger.Level >= logrus.InfoLevel {
		getLogrusEntry(ctx).Infof(f, v...)
	}
}

func Errorf(ctx Context, f string, v ...interface{}) {
	if logrusLogger.Level >= logrus.ErrorLevel {
		getLogrusEntry(ctx).Errorf(f, v...)
	}
}

func Warnf(ctx Context, f string, v ...interface{}) {
	if logrusLogger.Level >= logrus.WarnLevel {
		getLogrusEntry(ctx).Warnf(f, v...)
	}
}

func Panicf(ctx Context, f string, v ...interface{}) {
	if logrusLogger.Level >= logrus.PanicLevel {
		getLogrusEntry(ctx).Panicf(f, v...)
	}
}

func getLogrusEntry(ctx Context) (result *logrus.Entry) {
	ctxMap := ctx.GetAllMap()
	result = logrusLogger.WithFields(logrus.Fields(ctxMap))
	return
}

// formatter

type cornerstoneJSONFormatter struct{}

func (ftr *cornerstoneJSONFormatter) Format(entry *logrus.Entry) (result []byte, err error) {
	data := make(logrus.Fields, len(entry.Data)+buildInFieldNumber)
	for k, v := range entry.Data {
		data[k] = v
	}

	data[logKeyAppName] = appName
	data[logKeyAppVersion] = appVersion
	data[logKeyTimestamp] = entry.Time.UTC().Format(timestampFormat)
	data[logKeyMessage] = entry.Message
	data[logKeyLevel] = entry.Level.String()

	serialized, err := json.Marshal(data)
	if err != nil {
		err = fmt.Errorf("failed to marshal json format, err: %+v", err)
		return
	}

	result = append(serialized, '\n')
	return
}
