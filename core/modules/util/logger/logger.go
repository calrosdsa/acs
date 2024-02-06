package logger

import (
	_r "acs/domain/repository"
	"fmt"
	"log"
	"time"
	"os"
	"github.com/sirupsen/logrus"
	"strconv"
	"github.com/spf13/viper"
)

type logger struct{}

func New() _r.Logger {
	return &logger{}
}

func (l logger) LogError(method string, file string, err error) {
	log.Println(err)
	now := time.Now()
	t := fmt.Sprintf("%slog/%s-%s-%s", viper.GetString("path"), strconv.Itoa(now.Year()), now.Month().String(), strconv.Itoa(now.Day()))
	f, errL := os.OpenFile(t, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if errL != nil {
		log.Println("FAIL-", errL)
	}
	logrus.SetOutput(f)
	defer func() {
		log.Println("closing file")
		f.Close()
	}()
	ctx := logrus.WithFields(logrus.Fields{
		"method": method,
		"file":   file,
	})
	ctx.Error(err)
}
