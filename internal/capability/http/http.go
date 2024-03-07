package http

import (
	"reflect"
	"soarca/logger"
	"soarca/models/cacao"
	"soarca/models/execution"
	"soarca/utils/http"
)

// Receive HTTP API command data from decomposer/executer
// Validate HTTP API call
// Run HTTP API call
// Return response

type HttpCapability struct {
}

type Empty struct{}

var component = reflect.TypeOf(Empty{}).PkgPath()
var log *logger.Log

func init() {
	log = logger.Logger(component, logger.Info, "", logger.Json)
}

func (httpCapability *HttpCapability) Execute(
	metadata execution.Metadata,
	command cacao.Command,
	authentication cacao.AuthenticationInformation,
	target cacao.AgentTarget,
	variables cacao.VariableMap) (cacao.VariableMap, error) {

	soarca_http_request := new(http.HttpRequest)
	soarca_http_options := http.HttpOptions{
		Target:  &target,
		Command: &command,
		Auth:    &authentication,
	}

	responseBytes, err := soarca_http_request.Request(soarca_http_options)
	if err != nil {
		log.Error(err)
		return cacao.VariableMap{}, err
	}
	respString := string(responseBytes)

	return cacao.VariableMap{
		"__soarca_http_result__": {Name: "result", Value: respString}}, nil

}
