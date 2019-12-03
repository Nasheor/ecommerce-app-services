package main 

import (
	"time"
	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next LoginService
}

func (mw loggingMiddleware) validateCredentials(uname string, pword string, admin bool) (result bool) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "validateCredentials",
			"username", uname,
			"admin", admin,
			"result", result,
			"took", time.Since(begin),
		)
	}(time.Now())
	result = mw.next.validateCredentials(uname, pword, admin)
	return
}