package mock_capability

import (
	"soarca/models/cacao"
	"soarca/models/execution"

	"github.com/stretchr/testify/mock"
)

type Mock_Capability struct {
	mock.Mock
}

func (capability *Mock_Capability) Execute(metadata execution.Metadata,
	command cacao.Command,
	authentication cacao.AuthenticationInformation,
	target cacao.AgentTarget,
	variables cacao.VariableMap) (cacao.VariableMap, error) {
	args := capability.Called(metadata, command, authentication, target, variables)
	return args.Get(0).(cacao.VariableMap), args.Error(1)
}

func (capability *Mock_Capability) GetType() string {
	args := capability.Called()
	return args.Get(0).(string)
}
