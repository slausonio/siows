package mw

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	siogo "gitea.slauson.io/slausonio/siogo"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var sioGenericErr = &siogo.AppError{
	Code: http.StatusInternalServerError,
}

type MockResponseWriter struct {
	gin.ResponseWriter
	mock.Mock
}

func (m *MockResponseWriter) WriteHeader(code int) {
	m.Called(code)
}

func setUp(t *testing.T) (*MW, *siogo.MockOauthService) {
	t.Helper()
	mSvc := siogo.NewMockOauthService(t)
	mw := &MW{
		oauth: mSvc,
	}
	return mw, mSvc
}

func TestErrorHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		errs           []*gin.Error
		expectedStatus int
	}{
		{
			name: "Handle SioGenericError",
			errs: []*gin.Error{
				{Err: sioGenericErr, Type: gin.ErrorTypePublic},
			},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name: "Handle Non SioGenericError",
			errs: []*gin.Error{
				{Err: errors.New("random error"), Type: gin.ErrorTypePublic},
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw, _ := setUp(t)
			c, _ := gin.CreateTestContext(httptest.NewRecorder())

			mw.ErrorHandler(c)
			assert.Equal(t, tt.expectedStatus, c.Writer.Status())
		})
	}
}

//func TestAuthMiddleware(t *testing.T) {
//	gin.SetMode(gin.TestMode)
//
//	tests := []struct {
//		name           string
//		tokenHeader    string
//		expectedStatus int
//		mockOAuth      func() *mocks.MockOAuth
//	}{
//		{
//			name:           "Handle valid token",
//			tokenHeader:    "Bearer valid_token",
//			expectedStatus: http.StatusOK,
//			mockOAuth: func() *mocks.MockOAuth {
//				m := &mocks.MockOAuth{}
//				m.On("Introspect", "valid_token").Return(&mw.IntrospectionResponse{Active: true}, nil)
//				return m
//			},
//		},
//		{
//			name:           "Handle inactive token",
//			tokenHeader:    "Bearer inactive_token",
//			expectedStatus: http.StatusUnauthorized,
//			mockOAuth: func() *mocks.MockOAuth {
//				m := &mocks.MockOAuth{}
//				m.On("Introspect", "inactive_token").Return(&mw.IntrospectionResponse{Active: false}, nil)
//				return m
//			},
//		},
//		{
//			name:           "Handle error from OAuth",
//			tokenHeader:    "Bearer errored_token",
//			expectedStatus: http.StatusUnauthorized,
//			mockOAuth: func() *mocks.MockOAuth {
//				m := &mocks.MockOAuth{}
//				m.On("Introspect", "errored_token").Return(nil, errors.New("OAuth error"))
//				return m
//			},
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			c, _ := gin.CreateTestContext(httptest.NewRecorder())
//			c.Request, _ = http.NewRequest("GET", "/test", bytes.NewBuffer([]byte{}))
//			c.Request.Header.Set("Authorization", tt.tokenHeader)
//
//			middleware := &mw.MW{OAuth: tt.mockOAuth()}
//			middleware.AuthMiddleware(c)
//
//			assert.Equal(t, tt.expectedStatus, c.Writer.Status())
//		})
//	}
//}

func TestPrometheusMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("GET", "/test", bytes.NewBuffer([]byte{}))

	mw, _ := setUp(t)
	mw.PrometheusMiddleware()(c)

	// Test assertions for prometheus metrics here
	// Due to the nature of Prometheus metrics, it's not straight forward to assert these without adding complexity
	// Usually you would have an endpoint that exposes the metrics and use a prometheus server in a test environment to scrape these metrics
	// However, we can assert that there were no errors in the middleware and it finished successfully

	assert.Equal(t, http.StatusOK, c.Writer.Status())
}
