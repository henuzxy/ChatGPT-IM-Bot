package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log *logrus.Entry

func init() {
	log = logrus.NewEntry(logrus.New())
	log.Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	file, err := os.OpenFile("logs/server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log.Logger.Out = file
}
func Info(str string) {
	log.Info(str)
}
func Infof(formatstr string, args ...interface{}) {
	log.Infof(formatstr, args)
}
func Debug(str string) {
	log.Debug(str)
}
func Error(str string) {
	log.Error(str)
}
func Errorf(formatstr string, args ...interface{}) {
	log.Errorf(formatstr, args)
}
