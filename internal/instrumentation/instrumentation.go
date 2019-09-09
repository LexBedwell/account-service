package instrumentation

import (
	"fmt"
	"time"
	"github.com/go-kit/kit/metrics"
	"github.com/lexbedwell/account-service/internal/usecase/service"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

type InstrumentingMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	Next           service.AccountServiceInterface
}

func (mw *InstrumentingMiddleware) GetPongFromPing() (output string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetPongFromPing", "error", "false"}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output = mw.Next.GetPongFromPing()
	return
}

func (mw *InstrumentingMiddleware) GetInfoFromId(id string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetInfoFromId", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.GetInfoFromId(id)
	return
}

func (mw *InstrumentingMiddleware) PostUser(email string) (outputId string, outputEmail string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "PostUser", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	outputId, outputEmail, err = mw.Next.PostUser(email)
	return
}

func NewInstrumentingMiddleware(svc service.AccountServiceInterface) *InstrumentingMiddleware {

	fieldKeys := []string{"method", "error"}

	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "account_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
    }, fieldKeys)
    
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "account_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
		}, fieldKeys)
		
		return &InstrumentingMiddleware{requestCount, requestLatency, svc}

}