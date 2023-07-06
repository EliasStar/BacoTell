package main

import (
	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
)

const exampleModalCustomID = pluginID + "_modal"
const exampleModalComponentCustomID = "message"

type ExampleModal struct{}

var _ common.Modal = ExampleModal{}

// CustomID implements bacotell_common.Modal.
func (m ExampleModal) CustomID() (string, error) {
	return exampleModalCustomID, nil
}

// Submit implements bacotell_common.Modal.
func (m ExampleModal) Submit(proxy common.SubmitProxy) error {
	value, err := proxy.InputValue(exampleModalComponentCustomID)
	reply := "User wrote: " + value
	if err != nil {
		logger.Error("no content", "err", err)
		reply = "User wrote nothing!"
	}

	return proxy.Respond(common.Response{Content: reply}, false)
}
