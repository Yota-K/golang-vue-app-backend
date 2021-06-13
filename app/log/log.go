package log

import (
	"time"
	"github.com/sirupsen/logrus"
)

type Formatter struct {
	logrus.Formatter
}

// ログのフォーマットを指定
func (f Formatter) Format(e *logrus.Entry) ([]byte, error) {
	const Location = "Asia/Tokyo"
	loc, er := time.LoadLocation(Location)
	if er != nil {
		loc = time.FixedZone(Location, 9*60*60)
	}
	now := time.Now().In(loc)
	e.Time = now
	return f.Formatter.Format(e)
}

func init() {
	logrus.SetFormatter(Formatter{&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}})
}

func InfoLog(message, path string) {
	logrus.WithFields(logrus.Fields{
		"path": path,
	}).Info(message)
}
