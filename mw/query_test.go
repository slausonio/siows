package mw

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestIntQueryMiddleware_Success(t *testing.T) {
	req, err := http.NewRequest("GET", "/?param=123", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(recorder)
	c.Request = req

	IntQueryMiddleware()(c)

	if !c.IsAborted() {
		intParamsVal, exists := c.Get("intParams")
		if !exists {
			t.Fatal("Expected intParams to be set in the context")
		}

		intParams, ok := intParamsVal.(map[string]int)
		if !ok {
			t.Fatal("Expected intParams to be of type map[string]int")
		}

		expected := 123
		if val, ok := intParams["param"]; !ok || val != expected {
			t.Errorf("Expected intParams['param'] to be %d, got %d", expected, val)
		}
	} else {
		t.Fatal("Request was aborted when it should not have been")
	}
}

func TestIntQueryMiddleware_Failure(t *testing.T) {
	req, err := http.NewRequest("GET", "/?invalid=abc", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(recorder)
	c.Request = req

	IntQueryMiddleware()(c)

	if c.IsAborted() {
		// Checking status code. Assuming serr.NewSioBadRequestError(serr.INVALID_ID) returns 400.
		if c.Writer.Status() != http.StatusBadRequest {
			t.Errorf("Expected status %d, got %d", http.StatusBadRequest, c.Writer.Status())
		}
	} else {
		t.Fatal("Expected request to be aborted due to invalid parameter")
	}
}
