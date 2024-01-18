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

func (mw *MW) ErrorHandler(c *gin.Context) {
	c.Next()

	for _, err := range c.Errors {
		if err != nil {

			var code int
			eResponse := &ErrorResponse{
				Method: c.Request.Method,
				Path:   c.Request.URL.Path,
				Error:  err.Error(),
			}

			var t *siocore.AppError
			switch {
			default:
				code = http.StatusInternalServerError
				slog.Error(t.Error(), slog.Int(SlogKeyStatusCode, code))
			case errors.As(err.Err, &t):
				slog.Error(t.Error(), slog.Int(SlogKeyStatusCode, t.Code))
			}

			c.JSON(code, eResponse)
		}
	}
}

func (mw *MW) AuthMiddleware(c *gin.Context) {
	tokenHeader := c.Request.Header.Get("Authorization")

	if tokenHeader != "" {
		headerArray := strings.Split(tokenHeader, " ")
		if len(headerArray) == 2 {
			mw.handleHeader(tokenHeader, c)
		} else {
			_ = c.AbortWithError(401, unauthorizedErr)
		}

	} else {
		_ = c.AbortWithError(401, unauthorizedErr)
	}

	c.Next()
}

func (mw *MW) PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		timer := prometheus.NewTimer(prometheus.ObserverFunc(func(duration float64) {
			method := c.Request.Method
			path := c.FullPath()
			status := c.Writer.Status()
			metrics.HttpRequestsTotal.With(prometheus.Labels{"method": method, "path": path, "status_code": strconv.Itoa(status)}).
				Inc()

			metrics.HttpRequestDuration.With(prometheus.Labels{"method": method, "path": path, "status_code": strconv.Itoa(status)}).
				Observe(duration)

			if status >= 500 {
				metrics.HttpRequestFailures.With(prometheus.Labels{"method": method, "path": path, "status_code": strconv.Itoa(status)}).
					Inc()
			}
		}))
		defer timer.ObserveDuration()

		c.Next()
	}
}

func (mw *MW) handleHeader(tokenHeader string, c *gin.Context) {
	tokenHeader = strings.Split(tokenHeader, " ")[1]
	if ir, intoErr := mw.oauth.Introspect(tokenHeader); intoErr != nil {
		_ = c.AbortWithError(401, siocore.NewInternalServerError(intoErr.Error()))
	} else if !ir.Active {
		_ = c.AbortWithError(401, unauthorizedErr)
	}

	c.Next()
}
