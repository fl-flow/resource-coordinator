package logger

import (
  "github.com/sirupsen/logrus"
)


func init() {
  logrus.SetLevel(logrus.InfoLevel)
  logrus.SetFormatter(&logrus.JSONFormatter{})
}
