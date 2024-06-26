package endpoints_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-mail-sender/internal/endpoints"
	"go-mail-sender/internal/internal_errors"

	"github.com/stretchr/testify/assert"
)

func Test_HanderError_WhenEndpointReturnsInternalError(t *testing.T) {
	assert := assert.New(t)

	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 0, internal_errors.ErrInternal
	}

	handlerFunc := endpoints.HandlerError(endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusInternalServerError, res.Code)
	assert.Contains(res.Body.String(), internal_errors.ErrInternal.Error())
}

func Test_HanderError_WhenEndpointReturnsDomainError(t *testing.T) {
	assert := assert.New(t)

	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 0, errors.New("domain error")
	}

	handlerFunc := endpoints.HandlerError(endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusBadRequest, res.Code)
	assert.Contains(res.Body.String(), "domain error")
}

func Test_HanderError_WhenEndpointReturnsObjAndStatus(t *testing.T) {
	assert := assert.New(t)

	type BodyForTest struct {
		ID int
	}

	objExpected := BodyForTest{ID: 2}

	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return objExpected, 201, nil
	}

	handlerFunc := endpoints.HandlerError(endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	objReturned := BodyForTest{}
	json.Unmarshal(res.Body.Bytes(), &objReturned)

	assert.Equal(http.StatusCreated, res.Code)

	assert.Equal(objExpected, objReturned)
}
