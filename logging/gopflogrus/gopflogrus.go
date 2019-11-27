// Package gopflogrus provides structured logging with logrus.
package gopflogrus

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/kanhaiya15/gopf/utils"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	// Logger is a configured logrus.Logger.
	Logger *logrus.Logger
)

//LogFormat Log formatting struct
type LogFormat struct {
	TimestampFormat string
}

// Format Customt formatting
func (f *LogFormat) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer

	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	appName, err := utils.GetConfValue("APP_NAME")
	if err != nil {
		panic(err.Error())
	}
	b.WriteString("producerType==\"service\" ")
	b.WriteString("producerName==\"" + appName + "\" ")
	b.WriteString("time==\"")
	b.WriteString(entry.Time.UTC().Format(time.RFC3339) + "\"")

	b.WriteString(" ")
	b.WriteString("level==")
	fmt.Fprint(b, "\""+strings.ToUpper(entry.Level.String())+"\"")

	b.WriteString(" ")
	b.WriteString("message==")
	if entry.Message != "" {
		b.WriteString(fmt.Sprintf("\"%s\"", entry.Message))
	}
	b.WriteString(" ")

	//mdc
	for key, value := range entry.Data {
		b.WriteString(" " + key)
		b.WriteString("==\"")
		fmt.Fprint(b, value)
		b.WriteString("\"")
	}

	//end mdc

	b.WriteByte('\n')
	return b.Bytes(), nil
}

// NewLogger creates and configures a new logrus Logger.
func NewLogger() *logrus.Entry {
	dir := "/opt/logs"
	appName, err := utils.GetConfValue("APP_NAME")
	if err != nil {
		panic(err.Error())
	}
	appName += ".log"
	path := dir + "/" + appName

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err.Error())
		}
	}
	lumberjackLogger := &lumberjack.Logger{
		Filename:   path,
		MaxSize:    1,
		MaxBackups: 20,
		MaxAge:     28,
		LocalTime:  true,
	}
	logrus.SetOutput(lumberjackLogger)
	Logger = logrus.New()
	Logger.Formatter = &logrus.JSONFormatter{}
	Logger.Level = logrus.DebugLevel
	Logger.Out = lumberjackLogger

	mw := io.MultiWriter(os.Stdout, lumberjackLogger)
	Logger.Out = mw

	return Logger.WithFields(logrus.Fields{"producerType": "service", "producerName": "TMA"})

}
