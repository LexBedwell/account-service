package logging

import (
	"os"
	"time"
	"github.com/lexbedwell/account-service/internal/usecase/service"
	gokitlog "github.com/go-kit/kit/log"
)

type LoggingMiddleware struct {
	Logger 	gokitlog.Logger
	Next 	service.AccountServiceInterface
}

func (mw *LoggingMiddleware) GetPongFromPing() (output string) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "GetPongFromPing",
			"output", output,
			"took", time.Since(begin),
		)
	}(time.Now())

	output = mw.Next.GetPongFromPing()
	return
}

func (mw *LoggingMiddleware) GetInfoFromId(id string) (output string, err error) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "GetInfoFromId",
			"input", id,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.GetInfoFromId(id)
	return
}

func (mw *LoggingMiddleware) PostUser(email string) (outputId string, outputEmail string, err error) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "PostUser",
			"input", email,
			"outputId", outputId,
			"outputEmail", outputEmail,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	outputId, outputEmail, err = mw.Next.PostUser(email)
	return
}

func NewLoggingMiddleware(svc service.AccountServiceInterface) *LoggingMiddleware {
	logger := gokitlog.NewLogfmtLogger(os.Stderr)
	return &LoggingMiddleware{logger, svc}
}
