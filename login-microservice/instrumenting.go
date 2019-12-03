package main 

import (
	"time"


	"github.com/go-kit/kit/metrics"
)

type instrumentingMiddleware struct {
	requestCount metrics.Counter
	requestLatency metrics.Histogram
	countResult metrics.Histogram
	next LoginService
}

func (mw instrumentingMiddleware) validateCredentials(uname string, pword string, admin bool) (result bool){
	defer func(begin time.Time) {
		lsv := []string{"method", "validateCredentials"}
		mw.requestCount.With(lsv...).Add(1)
		mw.requestLatency.With(lsv...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	result = mw.next.validateCredentials(uname, pword, admin)
	return 
}
