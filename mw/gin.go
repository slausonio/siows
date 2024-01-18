package mw

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/redis/go-redis/v9"
	"github.com/slausonio/sioauth/authz"
	"github.com/slausonio/siocore/metrics"

	"gitea.slauson.io/slausonio/go-types/generic"
	siogo "gitea.slauson.io/slausonio/siogo"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var unauthorizedErr = siogo.NewUnauthorizedError("unauthorized")

type OauthAuthz interface {
	Get() (*authz.TokenResp, error)
	Introspect(token string) (*authz.IntrospectResp, error)
}

type MW struct {
	oauth OauthAuthz
}

func NewMW(redis *redis.Client, appEnv map[string]string) *MW {
	return &MW{
		oauth: siogo.NewOauthService(redis, appEnv),
	}
}

func (mw *MW) ErrorHandler(c *gin.Context) {
	c.Next()

	for _, err := range c.Errors {
		if err != nil {

			var code int
			eResponse := &generic.ErrorResponse{
				Method: c.Request.Method,
				Path:   c.Request.URL.Path,
				Error:  err.Error(),
			}

			var t *siogo.AppError
			switch {
			default:
				code = http.StatusInternalServerError
				logUnknownError(t, code)
			case errors.As(err.Err, &t):
				code = t.Code
				logAppError(t)
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
		// Start timer to track request duration
		timer := prometheus.NewTimer(prometheus.ObserverFunc(func(duration float64) {
			method := c.Request.Method
			path := c.FullPath()
			status := c.Writer.Status()
			metrics.HttpRequestsTotal.With(prometheus.Labels{"method": method, "path": path, "status_code": strconv.Itoa(status)}).
				Inc()

			// Record request duration
			metrics.HttpRequestDuration.With(prometheus.Labels{"method": method, "path": path, "status_code": strconv.Itoa(status)}).
				Observe(duration)

			// Record request failures (e.g., status code >= 500)
			if status >= 500 {
				metrics.HttpRequestFailures.With(prometheus.Labels{"method": method, "path": path, "status_code": strconv.Itoa(status)}).
					Inc()
			}
		}))
		defer timer.ObserveDuration()

		// Continue with the next handler in the chain
		c.Next()
	}
}

func (mw *MW) handleHeader(tokenHeader string, c *gin.Context) {
	tokenHeader = strings.Split(tokenHeader, " ")[1]
	if ir, intoErr := mw.oauth.Introspect(tokenHeader); intoErr != nil {
		_ = c.AbortWithError(401, siogo.NewInternalServerError(intoErr.Error()))
	} else if !ir.Active {
		_ = c.AbortWithError(401, unauthorizedErr)
	}

	c.Next()
}

// Used to push formatted error messages to the logrus
func logUnknownError(err error, code int) {
	stackTrace := siogo.GetRuntimeStack()
	slog.Error(err, stackTrace, code)
}

// Used to push formatted error messages to the logrus if AppError
func logAppError(err *siogo.AppError) {
	stackTrace := siogo.GetRuntimeStack()
	slog.Error(err, stackTrace, err.Code)
}
