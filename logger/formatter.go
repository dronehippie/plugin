package logger

import (
	"github.com/sirupsen/logrus"
)

type Formatter struct {
	logrus.TextFormatter
}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	res, err := f.TextFormatter.Format(entry)

	if err != nil {
		return []byte(""), err
	}

	return res[14:], nil
}
