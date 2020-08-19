package authorization

import (
	"github.com/go-kit/kit/log"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

type loggingService struct {
	logger log.Logger
	Service
}

func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}
