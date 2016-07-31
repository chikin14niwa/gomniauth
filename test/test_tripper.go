package test

import (
	"github.com/chikin14niwa/gomniauth/common"
	"github.com/chikin14niwa/testify/mock"
	"net/http"
)

type TestTripper struct {
	mock.Mock
}

func (t *TestTripper) Credentials() *common.Credentials {
	return t.Called().Get(0).(*common.Credentials)
}

func (t *TestTripper) Provider() common.Provider {
	return t.Called().Get(0).(common.Provider)
}

func (t *TestTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	args := t.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}
