package hooks

import "net/http"

/*
 * This file is only ever generated once on the first generation and then is free to be modified.
 * Any hooks you wish to add should be registered in the initHooks function. Feel free to define
 * your hooks in this file or in separate files in the hooks package.
 *
 * Hooks are registered per SDK instance, and are valid for the lifetime of the SDK instance.
 */

var _globalProjectId string

// SetProjectId sets the project ID for all SDK requests.
func SetProjectId(projectID string) {
	_globalProjectId = projectID
}

type projectIdHook struct{}

func (h *projectIdHook) BeforeRequest(hookCtx BeforeRequestContext, req *http.Request) (*http.Request, error) {
	if _globalProjectId != "" {
		req.Header.Set("X-Project-Id", _globalProjectId)
	}
	return req, nil
}

func initHooks(h *Hooks) {
	hook := &projectIdHook{}
	h.registerBeforeRequestHook(hook)
}
