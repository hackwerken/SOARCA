package fin

import (
	"reflect"
	"soarca/internal/fin/protocol"
	"soarca/logger"
	"soarca/models/cacao"
	"soarca/models/execution"
	finModel "soarca/models/fin"
)

type FinCapability struct {
	finProtocol protocol.IFinProtocol
}

var component = reflect.TypeOf(FinCapability{}).PkgPath()
var log *logger.Log

func init() {
	log = logger.Logger(component, logger.Info, "", logger.Json)
}

func New(finProtocol protocol.IFinProtocol) *FinCapability {
	return &FinCapability{finProtocol: finProtocol}
}

func (finCapability *FinCapability) Execute(
	metadata execution.Metadata,
	command cacao.Command,
	authentication cacao.AuthenticationInformation,
	target cacao.AgentTarget,
	variables map[string]cacao.Variable) (map[string]cacao.Variable, error) {

	finCommand := finModel.NewCommand()
	finCommand.CommandSubstructure.Command = command.Command
	finCommand.CommandSubstructure.Authentication = authentication
	finCommand.CommandSubstructure.Variables = variables
	finCommand.CommandSubstructure.Context.ExecutionId = metadata.ExecutionId.String()
	finCommand.CommandSubstructure.Context.PlaybookId = metadata.PlaybookId
	finCommand.CommandSubstructure.Context.StepId = metadata.StepId

	log.Trace("created command ", finCommand)
	return finCapability.finProtocol.SendCommand(finCommand)
}
