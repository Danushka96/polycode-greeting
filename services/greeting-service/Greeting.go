package greeting_service

import (
	"github.com/cloudimpl/next-coder-sdk/polycode"
	"portal/register/model"
)

// @description return a greeting message for given name
func Greeting(ctx polycode.ServiceContext, input model.HelloRequest) (model.HelloResponse, error) {
	return model.HelloResponse{
		Message: "Hello" + input.Name,
	}, nil
}

func WSGreeting(ctx polycode.WorkflowContext, input model.HelloRequest) (model.HelloResponse, error) {
	response := model.HelloResponse{Message: "Hello " + input.Name}
	err := ctx.EmitRealtimeEvent("channel2", response)
	return response, err
}
