package mw

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/redis/go-redis/v9"
	"github.com/slausonio/sioauth/authz"
	"github.com/slausonio/siocore"
	"github.com/slausonio/siocore/metrics"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	SlogKeyStatusCode = "statusCode"
)

var unauthorizedErr = siocore.NewUnauthorizedError("unauthorized")

type ErrorResponse struct {
	Method string `json:"method"`
	Path   string `json:"path"`
	Error  string `json:"error"`
}

type OauthAuthz interface {
	Get() (*authz.TokenResp, error)
	Introspect(token string) (*authz.IntrospectResp, error)
}

type MW struct {
	oauth OauthAuthz
}

func NewMW(redis *redis.Client, appEnv siocore.Env) *MW {
	return &MW{
		oauth: authz.NewAuthz(redis, appEnv),
	}
}

func (mw *MW) ErrorHandler(ctx *gin.Context) {
	ctx.Next()

	for _, err := range ctx.Errors {
		if err != nil {
			var code int
			eResponse := &ErrorResponse{
				Method: ctx.Request.Method,
				Path:   ctx.Request.URL.Path,
				Error:  err.Error(),
			}

			var appErr *siocore.AppError

			switch {
			default:
				code = http.StatusInternalServerError
				slog.Error(appErr.Error(), slog.Int(SlogKeyStatusCode, code))
			case errors.As(err.Err, &appErr):
				slog.Error(appErr.Error(), slog.Int(SlogKeyStatusCode, appErr.Code))
			}

			ctx.JSON(code, eResponse)
		}
	}
}

func (mw *MW) AuthMiddleware(ctx *gin.Context) {
	tokenHeader := ctx.Request.Header.Get("Authorization")

	if tokenHeader != "" {
		headerArray := strings.Split(tokenHeader, " ")
		if len(headerArray) == 2 {
			mw.handleHeader(tokenHeader, ctx)
		} else {
			_ = ctx.AbortWithError(401, unauthorizedErr)
		}
	} else {
		_ = ctx.AbortWithError(401, unauthorizedErr)
	}

	ctx.Next()
}

func (mw *MW) PrometheusMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		timer := prometheus.NewTimer(prometheus.ObserverFunc(func(duration float64) {
			method := ctx.Request.Method
			path := ctx.FullPath()
			status := ctx.Writer.Status()
			metrics.HttpRequestsTotal.With(
				prometheus.Labels{
					"method":      method,
					"path":        path,
					"status_code": strconv.Itoa(status),
				},
			).
				Inc()

			metrics.HttpRequestDuration.With(
				prometheus.Labels{
					"method":      method,
					"path":        path,
					"status_code": strconv.Itoa(status),
				},
			).
				Observe(duration)

			if status >= 500 {
				metrics.HttpRequestFailures.With(
					prometheus.Labels{
						"method":      method,
						"path":        path,
						"status_code": strconv.Itoa(status),
					},
				).
					Inc()
			}
		}))
		defer timer.ObserveDuration()

		ctx.Next()
	}
}

func (mw *MW) handleHeader(tokenHeader string, ctx *gin.Context) {
	tokenHeader = strings.Split(tokenHeader, " ")[1]
	if ir, intoErr := mw.oauth.Introspect(tokenHeader); intoErr != nil {
		_ = ctx.AbortWithError(401, siocore.NewInternalServerError(intoErr.Error()))
	} else if !ir.Active {
		_ = ctx.AbortWithError(401, unauthorizedErr)
	}

	ctx.Next()
}
