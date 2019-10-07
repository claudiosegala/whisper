package api

import (
	"github.com/labbsr0x/whisper/web/config"
	"net/http"
)

type ScopeAPI interface {
	LoginGETHandler(route string) http.Handler
	LoginPUTHandler() http.Handler
}

// DefaultScopeAPI holds the default implementation of the User API interface
type DefaultScopeAPI struct {
	*config.WebBuilder
}

// InitFromWebBuilder initializes a default login api instance
func (dapi *DefaultScopeAPI) InitFromWebBuilder(webBuilder *config.WebBuilder) *DefaultScopeAPI {
	dapi.WebBuilder = webBuilder
	return dapi
}

// LoginPOSTHandler post form handler for logging in users
func (dapi *DefaultScopeAPI) ScopePUTHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

// LoginGETHandler prompts the browser to the login UI or redirects it to hydra
func (dapi *DefaultLoginAPI) ScopeGETHandler(route string) http.Handler {
	return http.StripPrefix(route, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	}))
}
